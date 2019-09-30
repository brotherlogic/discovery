package main

import (
	"context"
	"errors"
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

	pr, _ := peer.FromContext(ctx)
	s.discoverPeer = fmt.Sprintf("%+v", pr)

	//Reject requests without caller
	if req.Caller == "" {
		pr, _ := peer.FromContext(ctx)
		s.peerFail = fmt.Sprintf("%+v", pr)
		return nil, fmt.Errorf("Must specify caller")
	}

	if val, ok := s.version.Load(req.GetRequest().GetName()); ok {
		if val.(int32) == 1 {
			resp, err := s.Get(ctx, &pb.GetRequest{Job: req.GetRequest().GetName(), Server: req.GetRequest().GetIdentifier()})
			return &pb.DiscoverResponse{Service: resp.GetServices()[0]}, err
		}
	}
	s.countDiscover++
	in := req.GetRequest()

	s.callerCountM.Lock()
	s.reqCountM.Lock()
	s.callerCount[req.Caller]++
	s.reqCount[in.GetName()]++
	s.reqCountM.Unlock()
	s.callerCountM.Unlock()

	// Check if we've been asked for something specific
	if in.GetIdentifier() != "" && in.GetName() != "" {
		in.Port = s.hashPortNumber(in.GetIdentifier(), in.GetName(), SEP)
		val := s.getCurr(in)
		if val != nil {
			return &pb.DiscoverResponse{Service: val}, nil
		}

		return nil, fmt.Errorf("Unable to locate %v on server %v", in.GetName(), in.GetIdentifier())
	}

	// Return the master if it exists
	val := s.getCMaster(in)
	if val != nil && val.LastSeenTime+val.TimeToClean*1000000 > time.Now().UnixNano() && !val.WeakMaster {
		return &pb.DiscoverResponse{Service: val}, nil
	}

	return &pb.DiscoverResponse{}, errors.New("Cannot find master for " + in.GetName() + " on server " + in.GetIdentifier())
}

//State gets the state of the server
func (s *Server) State(ctx context.Context, in *pb.StateRequest) (*pb.StateResponse, error) {
	return &pb.StateResponse{}, nil
}
