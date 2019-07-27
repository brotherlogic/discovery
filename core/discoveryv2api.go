package discovery

import (
	"context"
	"fmt"
	"time"

	pb "github.com/brotherlogic/discovery/proto"
)

// Register a server
func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	s.countV2Register++

	curr, _ := s.getJob(req.GetService())

	// Fail a re-register
	if curr != nil {
		return nil, fmt.Errorf("Already registered")
	}

	//Place this in the port map
	s.addToPortMap(req.GetService())

	// This is a new registration - update the port map
	req.GetService().LastSeenTime = time.Now().UnixNano()
	return &pb.RegisterResponse{Service: req.GetService()}, nil
}

func (s *Server) getMaster(ctx context.Context, job string) (*pb.GetResponse, error) {
	s.masterv2Mutex.Lock()
	defer s.masterv2Mutex.Unlock()

	if val, ok := s.masterv2[job]; ok {
		return &pb.GetResponse{Services: []*pb.RegistryEntry{val}}, nil
	}

	return nil, fmt.Errorf("Master elect not implemented yet")
}

// Get an entry from the registry
func (s *Server) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	if len(req.Server) == 0 && len(req.Job) != 0 {
		return s.getMaster(ctx, req.Job)
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
