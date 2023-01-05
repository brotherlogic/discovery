package main

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"

	pb "github.com/brotherlogic/discovery/proto"
	"github.com/brotherlogic/goserver/utils"
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

func (s *Server) SetZone(ctx context.Context, req *pb.SetZoneRequest) (*pb.SetZoneResponse, error) {
	ioutil.WriteFile("/home/simon/zone", []byte(req.GetZone()), 0644)
	return &pb.SetZoneResponse{}, nil
}

// Register a server
func (s *Server) RegisterV2(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	// We're receiving a registration about ourselves?
	if req.GetService().GetIdentifier() == s.Registry.Identifier {
		if req.GetService().GetZone() != s.zone && req.GetService().GetZone() != "" {
			return nil, status.Errorf(codes.OutOfRange, "Zone mismatch: '%v' -> '%v'", req.GetService().GetZone(), s.Registry.GetZone())
		}
	}

	defer s.doWrite()
	s.countV2Register++
	register.With(prometheus.Labels{"service": req.GetService().GetName(), "origin": req.GetCaller()}).Inc()

	s.addIP(req.GetService().GetIp())

	// Set the zone if this is an origin request
	if req.GetService().GetZone() == "" {
		req.GetService().Zone = s.zone
	}

	// Fail register until we're ready to serve
	if s.friendTime <= 0 && !req.GetFanout() {
		return nil, status.Errorf(codes.FailedPrecondition, "Discover is not yet ready to perform registration")
	}

	s.checkFriend(fmt.Sprintf("%v", req.GetService().GetIp()))

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
	//Immediate fail if we have no context key
	key, err := utils.GetContextKey(ctx)
	if err != nil || key == "" {
		pinfo, ok := peer.FromContext(ctx)
		s.CtxLog(ctx, fmt.Sprintf("PEER: %v -> %+v", ok, pinfo))
		return nil, fmt.Errorf("You need to provide a context key: %v, %v", key, err)
	}

	jobName := "unknown"
	if len(req.GetJob()) > 0 {
		jobName = req.GetJob()
	}

	if jobName == "tracer" {
		return nil, fmt.Errorf("no tracer")
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
			s.friends = append(s.friends, req.GetFriend())
			Friends.With(prometheus.Labels{"state": fmt.Sprintf("%v", s.state)}).Set(float64(len(s.friends)))
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

// Unregister a service from the listing
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

	err := s.removeFromPortMap(ctx, req.GetService())

	if err == nil && !req.Fanout {
		req.Fanout = true
		s.fanoutUnregister(ctx, req)
	}

	return &pb.UnregisterResponse{}, err
}

// Lock in prep for master elect
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
	return &pb.GetConfigResponse{Config: s.config}, nil
}
