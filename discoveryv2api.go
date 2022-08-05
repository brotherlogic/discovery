package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"

	pb "github.com/brotherlogic/discovery/proto"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const (
	IP_FILE = "/media/scratch/discovery-list"
)

var (
	register = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "discovery_register",
		Help: "The size of the print queue",
	}, []string{"service", "origin"})
	get = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "discovery_get",
		Help: "The size of the get queue",
	}, []string{"service", "origin"})

	unregister = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "discovery_unregister",
		Help: "The size of the print queue",
	}, []string{"service", "origin"})
)

func (s *Server) addIP(ip string) {
	for _, lip := range s.iplist {
		if lip == ip {
			return
		}
	}

	s.iplist = append(s.iplist, ip)
	s.writeIplist(s.iplist)
}

func (s *Server) writeIplist(lis []string) {
	if _, err := os.Stat(IP_FILE); errors.Is(err, os.ErrNotExist) {
		_, err := os.Create(IP_FILE)
		if err != nil {
			return
		}
	}

	fw, err := os.OpenFile(IP_FILE, os.O_WRONLY, 0777)
	if err != nil {
		return
	}
	defer fw.Close()

	for _, str := range lis {
		fw.WriteString(fmt.Sprintf("%v\n", str))
	}
}

func (s *Server) isFriend(host string) bool {
	for _, f := range s.friends {
		if f == host {
			return true
		}
	}

	return false
}

// Register a server
func (s *Server) RegisterV2(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	defer s.doWrite()
	s.countV2Register++
	register.With(prometheus.Labels{"service": req.GetService().GetName(), "origin": req.GetCaller()}).Inc()

	s.addIP(req.GetService().GetIp())

	// Fail register until we're ready to serve
	if s.friendTime <= 0 && !req.GetFanout() {
		return nil, status.Errorf(codes.FailedPrecondition, "Discover is not yet ready to perform registration (%v and %v)", s.friendTime, req.GetFanout())
	}

	s.checkFriend(ctx, fmt.Sprintf("%v", req.GetService().GetIp()))

	curr := s.getJob(req.GetService())

	s.version.Store(req.GetService().GetName(), int32(1))

	// Fast path on a re-register
	if curr != nil {
		if !req.Fanout {
			req.Fanout = true
			s.fanoutRegister(ctx, req)
		}
		return &pb.RegisterResponse{Service: curr}, nil
	}

	//Place this in the port map
	s.addToPortMap(req.GetService())

	// This is a new registration - update the port map
	req.GetService().LastSeenTime = time.Now().UnixNano()

	// Update the friend time
	for key, value := range s.config.GetFriendState() {
		if key == fmt.Sprintf("%v:50055", req.GetService().GetIp()) {
			value.LastSeen = time.Now().Unix()
		}
	}

	if !req.Fanout {
		req.Fanout = true
		s.fanoutRegister(ctx, req)
	}

	return &pb.RegisterResponse{Service: req.GetService()}, nil
}

// Get an entry from the registry
func (s *Server) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	jobName := "unknown"
	if len(req.GetJob()) > 0 {
		jobName = req.GetJob()
	}

	get.With(prometheus.Labels{"service": jobName, "origin": req.GetFriend()}).Inc()

	s.mapLock.Lock()
	s.getMapB[jobName]++
	s.mapLock.Unlock()

	defer func() {
		s.mapLock.Lock()
		s.getMapB[jobName]--
		s.mapLock.Unlock()
	}()

	if len(req.GetFriend()) > 0 {
		found := false
		for _, friend := range s.friends {
			if friend == req.GetFriend() {
				found = true
				break
			}
		}

		if !found {
			check, _ := s.readFriend(ctx, req.GetFriend())
			if check {
				if !s.isFriend(req.GetFriend()) {
					s.friends = append(s.friends, req.GetFriend())
					Friends.With(prometheus.Labels{"state": fmt.Sprintf("%v", s.state)}).Set(float64(len(s.friends)))
				}
			}
		}
	}

	if val, ok := s.version.Load(req.GetJob()); ok {
		if val.(int32) == 0 {
			resp, err := s.Discover(ctx, &pb.DiscoverRequest{Caller: "v2", Request: &pb.RegistryEntry{Name: req.GetJob(), Identifier: req.GetServer()}})
			return &pb.GetResponse{Services: []*pb.RegistryEntry{resp.GetService()}, State: s.state}, err
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
			return &pb.GetResponse{Services: jobs, State: s.state}, nil
		}

		return nil, status.Errorf(codes.Unavailable, "%v has not found on %v (via %v)", req.Job, req.Server, s.Registry.GetIdentifier())
	}

	resp := &pb.GetResponse{Services: []*pb.RegistryEntry{}, State: s.state}

	for _, job := range s.portMap {
		if job != nil {
			resp.Services = append(resp.Services, job)
		}
	}
	return resp, nil
}

func (s *Server) doWrite() {
	if s.writePrometheus {
		var services []*pb.RegistryEntry

		for _, job := range s.portMap {
			if job != nil {
				services = append(services, job)
			}
		}
		s.writeFile("/etc/prometheus/jobs.json", services)
	}
}

//Unregister a service from the listing
func (s *Server) Unregister(ctx context.Context, req *pb.UnregisterRequest) (*pb.UnregisterResponse, error) {
	defer s.doWrite()

	if req.GetReason() == "" {
		return nil, fmt.Errorf("Unable to unregister without a good reason")
	}

	unregister.With(prometheus.Labels{"service": req.GetService().GetName(), "origin": req.GetCaller()}).Inc()
	if req.GetService() == nil {
		p, _ := peer.FromContext(ctx)
		return nil, status.Errorf(codes.InvalidArgument, "Attempting to unregister empty service: %v: %+v", req, p)
	}

	s.removeFromPortMap(req.GetService())

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

func (s *Server) GetFriends(_ context.Context, req *pb.GetFriendsRequest) (*pb.GetFriendsResponse, error) {
	return &pb.GetFriendsResponse{Friends: s.friends}, nil
}

func (s *Server) GetInternalState(_ context.Context, req *pb.GetStateRequest) (*pb.GetStateResponse, error) {
	return &pb.GetStateResponse{State: s.internalState}, nil
}

func (s *Server) GetConfig(_ context.Context, req *pb.GetConfigRequest) (*pb.GetConfigResponse, error) {
	return &pb.GetConfigResponse{
		Config: s.config}, nil
}
