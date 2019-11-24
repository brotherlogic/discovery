package main

import (
	"context"
	"fmt"
	"time"

	pb "github.com/brotherlogic/discovery/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) MasterElect(ctx context.Context, req *pb.MasterRequest) (*pb.MasterResponse, error) {
	curr, _ := s.getJob(req.GetService())

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
		return nil, fmt.Errorf("Cannot become master until %v", t.Add(time.Minute))
	}

	s.elector.unelect(ctx, m)

	key := time.Now().UnixNano()
	err := s.acquireMasterLock(ctx, curr.GetName(), key)
	if err != nil {
		return nil, fmt.Errorf("Unable to acquire lock to become master")
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
		return nil, fmt.Errorf("Discover is not yet ready to perform registration")
	}

	// Reject a master registration
	if req.GetService().GetMaster() && !req.GetFanout() {
		req.GetService().Master = false
	}

	curr, _ := s.getJob(req.GetService())

	s.version.Store(req.GetService().GetName(), int32(1))

	// Fast path on a re-register
	if curr != nil {
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
		}
	}

	if val, ok := s.version.Load(req.GetJob()); ok {
		if val.(int32) == 0 {
			resp, err := s.Discover(ctx, &pb.DiscoverRequest{Caller: "v2", Request: &pb.RegistryEntry{Name: req.GetJob(), Identifier: req.GetServer()}})
			return &pb.GetResponse{Services: []*pb.RegistryEntry{resp.GetService()}}, err
		}
	}

	if len(req.Job) != 0 {
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
	s.removeFromPortMap(req.GetService())

	master, _ := s.getCMaster(req.GetService())
	if master.GetIdentifier() == req.GetService().GetIdentifier() {
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
		if time.Now().Sub(time.Unix(0, val)) < time.Minute || req.GetLockKey() < val {
			return nil, fmt.Errorf("Unable to acquire master lock: %v or %v", time.Now().Sub(time.Unix(0, val)), req.GetLockKey()-val)
		}
	}
	s.locks[req.GetJob()] = req.GetLockKey()
	return &pb.LockResponse{}, nil
}
