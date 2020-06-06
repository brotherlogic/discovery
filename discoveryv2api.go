package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"

	pb "github.com/brotherlogic/discovery/proto"
)

func (s *Server) MasterElect(ctx context.Context, req *pb.MasterRequest) (*pb.MasterResponse, error) {
	curr, _ := s.getJob(req.GetService())
	if curr == nil {
		return nil, status.Errorf(codes.FailedPrecondition, "Job %v is not registered", req.GetService().GetName())
	}

	if req.GetFanout() {
		if val, ok := s.locks[req.GetService().GetName()]; !ok || val != req.GetLockKey() {
			return nil, fmt.Errorf("Lock was not acquired here")
		}
		s.addMaster(curr)
		curr.Master = true
		return &pb.MasterResponse{Service: curr}, nil
	}

	m, t := s.getCMaster(req.GetService())
	if m != nil && time.Now().Sub(t) < time.Minute {
		return nil, status.Errorf(codes.FailedPrecondition, "Cannot become master until %v -> current master is %v (%v is making the request)", t.Add(time.Minute), m, req.GetService())
	}

	s.elector.unelect(ctx, m)

	key := time.Now().UnixNano()
	err := s.acquireMasterLock(ctx, curr.GetName(), key)
	if err != nil {
		return nil, err
	}

	curr.Master = true
	s.addMaster(curr)
	req.Fanout = true
	req.LockKey = key
	s.fanoutMaster(ctx, req)
	return &pb.MasterResponse{Service: curr}, nil
}

// Register a server
func (s *Server) RegisterV2(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	s.countV2Register++

	// Fail register until we're ready to serve
	if s.friendTime <= 0 && !req.GetFanout() {
		return nil, status.Errorf(codes.FailedPrecondition, "Discover is not yet ready to perform registration")
	}

	// Collapse a master registration
	if req.GetService().GetMaster() && !req.GetFanout() {
		req.GetService().Master = false
	}

	curr, _ := s.getJob(req.GetService())

	s.version.Store(req.GetService().GetName(), int32(1))

	// Fast path on a re-register
	if curr != nil {
		if curr.Master && !req.GetService().GetMaster() {
			s.removeMaster(curr)
			curr.Master = false
			if !req.Fanout {
				req.Fanout = true
				s.fanoutRegister(ctx, req)
			}
		}
		return &pb.RegisterResponse{Service: curr}, nil
	}

	//Place this in the port map
	s.addToPortMap(req.GetService())

	// This is a new registration - update the port map
	req.GetService().LastSeenTime = time.Now().UnixNano()

	if !req.Fanout {
		req.Fanout = true
		s.fanoutRegister(ctx, req)
	}

	if req.Fanout && req.GetService().GetMaster() {
		s.addMaster(req.GetService())
	}

	return &pb.RegisterResponse{Service: req.GetService()}, nil
}

// Get an entry from the registry
func (s *Server) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	s.getLoad++
	defer func() {
		s.getLoad--
	}()

	jobName := "unknown"
	if len(req.GetJob()) > 0 {
		jobName = req.GetJob()
	}
	s.mapLock.Lock()
	s.getMapB[jobName]++
	s.mapLock.Unlock()

	defer func() {
		s.mapLock.Lock()
		s.getMapB[jobName]--
		s.mapLock.Unlock()
	}()

	if s.getLoad > 50 {
		s.RaiseIssue(ctx, "Overload", fmt.Sprintf("Discover on %v is recording %v get calls: %v", s.Registry, s.getLoad, s.getMapB), false)
	}

	if s.getLoad > 100 {
		fmt.Printf("Severe Overlad\n")
		os.Exit(1)
	}

	if len(req.GetFriend()) > 0 {
		found := false
		for _, friend := range s.friends {
			if friend == req.GetFriend() {
				found = true
				break
			}
		}

		if !found {
			s.friends = append(s.friends, req.GetFriend())
			elems := strings.Split(req.GetFriend(), ":")
			if len(elems) > 1 {
				blah, _ := strconv.Atoi(elems[1])
				s.countMap[blah] = fmt.Sprintf("%v %v", time.Now(), "FROM_API")
			}
			Friends.Set(float64(len(s.friends)))
		}
	}

	if val, ok := s.version.Load(req.GetJob()); ok {
		if val.(int32) == 0 {
			resp, err := s.Discover(ctx, &pb.DiscoverRequest{Caller: "v2", Request: &pb.RegistryEntry{Name: req.GetJob(), Identifier: req.GetServer()}})
			return &pb.GetResponse{Services: []*pb.RegistryEntry{resp.GetService()}}, err
		}
	}

	if len(req.Job) != 0 {
		cval, ok := s.getMap.LoadOrStore(req.GetJob(), 1)
		if ok {
			s.getMap.Store(req.GetJob(), cval.(int)+1)
		}
		jobs := []*pb.RegistryEntry{}
		for _, job := range s.portMap {
			if job != nil {
				if (len(req.GetServer()) == 0 || job.Identifier == req.Server) && job.Name == req.Job {
					jobs = append(jobs, job)
				}
			}
		}
		if len(jobs) > 0 {
			return &pb.GetResponse{Services: jobs}, nil
		}

		return nil, status.Errorf(codes.NotFound, "%v not found on %v", req.Job, req.Server)
	}

	resp := &pb.GetResponse{Services: []*pb.RegistryEntry{}}

	for _, job := range s.portMap {
		if job != nil {
			resp.Services = append(resp.Services, job)
		}
	}
	return resp, nil
}

//Unregister a service from the listing
func (s *Server) Unregister(ctx context.Context, req *pb.UnregisterRequest) (*pb.UnregisterResponse, error) {
	if req.GetService() == nil {
		p, _ := peer.FromContext(ctx)
		return nil, status.Errorf(codes.InvalidArgument, "Attempting to unregister empty service: %v: %+v", req, p)
	}

	s.removeFromPortMap(req.GetService())

	master, _ := s.getCMaster(req.GetService())
	if master != nil && master.GetIdentifier() == req.GetService().GetIdentifier() {
		s.removeMaster(req.GetService())
	}

	if !req.Fanout {
		req.Fanout = true
		s.fanoutUnregister(ctx, req)
	}

	return &pb.UnregisterResponse{}, nil
}

//Lock in prep for master elect
func (s *Server) Lock(ctx context.Context, req *pb.LockRequest) (*pb.LockResponse, error) {
	if val, ok := s.locks[req.GetJob()]; ok {
		if time.Now().Sub(time.Unix(0, val)) < time.Second*4 && req.GetLockKey() > val {
			if v2, ok2 := s.lockNames[req.GetJob()]; ok2 {
				if v2 != req.GetRequestor() {
					return nil, status.Errorf(codes.FailedPrecondition, "Unable to acquire master on %v for %v lock (held by %v not %v): %v or %v acq %v", s.Registry, req.GetJob(), v2, req.GetRequestor(), time.Now().Sub(time.Unix(0, val)), req.GetLockKey()-val, time.Unix(0, val))
				}
			}
		}
	}
	s.locks[req.GetJob()] = req.GetLockKey()
	s.lockNames[req.GetJob()] = req.GetRequestor()
	return &pb.LockResponse{}, nil
}
