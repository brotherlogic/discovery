package main

import (
	"context"
	"fmt"
	"time"

	pb "github.com/brotherlogic/discovery/proto"
)

// Register a server
func (s *Server) RegisterV2(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	s.countV2Register++

	curr, _ := s.getJob(req.GetService())

	s.version.Store(req.GetService().GetName(), int32(1))

	// Fail a re-register
	if curr != nil {

		// Perform master election if needed
		if req.GetMasterElect() {
			m, t := s.getCMaster(req.GetService())
			if m != nil && time.Now().Sub(t) < time.Minute {
				return nil, fmt.Errorf("Cannot become master until %v", t.Add(time.Minute))
			}

			curr.Master = true
			s.addMaster(curr)
			return &pb.RegisterResponse{Service: curr}, nil
		}

		return nil, fmt.Errorf("Already registered")
	}

	//Place this in the port map
	s.addToPortMap(req.GetService())

	// This is a new registration - update the port map
	req.GetService().LastSeenTime = time.Now().UnixNano()
	return &pb.RegisterResponse{Service: req.GetService()}, nil
}

// Get an entry from the registry
func (s *Server) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	if val, ok := s.version.Load(req.GetJob()); ok {
		if val.(int32) == 0 {
			resp, err := s.Discover(ctx, &pb.DiscoverRequest{Caller: "v2", Request: &pb.RegistryEntry{Name: req.GetJob(), Identifier: req.GetServer()}})
			return &pb.GetResponse{Services: []*pb.RegistryEntry{resp.GetService()}}, err
		}
	}

	if len(req.Job) != 0 {
		for _, job := range s.portMap {
			if job != nil {
				if job.Identifier == req.Server && job.Name == req.Job {
					return &pb.GetResponse{Services: []*pb.RegistryEntry{job}}, nil
				}
			}
		}

		return nil, fmt.Errorf("%v not found on %v", req.Job, req.Server)
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

	return &pb.UnregisterResponse{}, nil
}
