package main

import (
	"context"
	"fmt"
	"time"

	pb "github.com/brotherlogic/discovery/proto"
	"google.golang.org/grpc/peer"
)

// ListAllServices returns a list of all the services
func (s *Server) ListAllServices(ctx context.Context, in *pb.ListRequest) (*pb.ListResponse, error) {
	s.countList++
	entries := []*pb.RegistryEntry{}
	for _, val := range s.portMap {
		if val != nil {
			entries = append(entries, val)
		}
	}
	return &pb.ListResponse{Services: &pb.ServiceList{Services: entries}}, nil
}

// RegisterService supports the RegisterService rpc end point
func (s *Server) RegisterService(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	if req.GetService().GetName() != "PictureFrame" && req.GetService().GetName() != "GraphPlotter" && req.GetService().GetName() != "RecordSelector" {
		s.RaiseIssue(ctx, "Bad Register", fmt.Sprintf("%v is a v1 register", req), false)
	}
	s.countRegister++

	pr, _ := peer.FromContext(ctx)
	s.registerPeer = fmt.Sprintf("%+v", pr)

	//Reject the request with no time to clean
	if req.GetService().GetTimeToClean() == 0 {
		return &pb.RegisterResponse{}, fmt.Errorf("You must specify a clean time")
	}

	s.version.Store(req.GetService().GetName(), pb.RegistryEntry_Version_value[req.GetService().GetVersion().String()])
	curr, master := s.getJob(req.GetService())

	//Reject if this is a master request
	if req.GetService().GetMaster() {
		if master != nil && master.GetIdentifier() != req.GetService().GetIdentifier() && !master.WeakMaster {
			req.Service.Master = false
			return nil, fmt.Errorf("Unable to register as master - already exists on %v", master.GetIdentifier())
		}

		if curr == nil {
			curr = req.GetService()
			s.addToPortMap(curr)
		}

		// Do the master promote
		if master == nil || master.WeakMaster {
			curr.Master = true
			curr.WeakMaster = false
			s.addMaster(curr)
			master = curr
		}

	} else if master != nil && master.GetIdentifier() == req.GetService().GetIdentifier() {
		s.removeMaster(req.GetService())
	}

	if !req.GetService().GetMaster() {
		req.GetService().MasterTime = 0
	}

	//Place this in the port map
	if curr == nil {
		curr = req.GetService()
		s.addToPortMap(curr)
	}

	// Apply the weak lease
	if master == nil {
		curr.WeakMaster = true
	} else {
		curr.WeakMaster = false
	}

	// This is a new registration - update the port map
	curr.LastSeenTime = time.Now().UnixNano()
	return &pb.RegisterResponse{Service: curr}, nil
}

// Discover supports the Discover rpc end point
func (s *Server) Discover(ctx context.Context, req *pb.DiscoverRequest) (*pb.DiscoverResponse, error) {
	resp, err := s.Get(ctx, &pb.GetRequest{Job: req.GetRequest().GetName(), Server: req.GetRequest().GetIdentifier()})
	if len(resp.GetServices()) > 0 {
		return &pb.DiscoverResponse{Service: resp.GetServices()[0]}, err
	}
	return &pb.DiscoverResponse{}, err
}

//State gets the state of the server
func (s *Server) State(ctx context.Context, in *pb.StateRequest) (*pb.StateResponse, error) {
	return &pb.StateResponse{}, nil
}
