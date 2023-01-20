package main

import (
	"context"
	"fmt"
	"time"

	pb "github.com/brotherlogic/discovery/proto"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
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

	pr, _ := peer.FromContext(ctx)
	s.registerPeer = fmt.Sprintf("%+v", pr)

	//legacyPeer.With(prometheus.Labels{"method": "Register", "peer": fmt.Sprintf("%v", strings.Split(pr.Addr.String(), ":")[0])}).Inc()

	if req.GetService().GetName() != "PictureFrame" && req.GetService().GetName() != "GraphPlotter" && req.GetService().GetName() != "RecordSelector" {
		s.RaiseIssue("Bad Register", fmt.Sprintf("%v is a v1 register", req))
	}
	s.countRegister++

	//Reject the request with no time to clean
	if req.GetService().GetTimeToClean() == 0 {
		return &pb.RegisterResponse{}, fmt.Errorf("You must specify a clean time")
	}

	s.version.Store(req.GetService().GetName(), pb.RegistryEntry_Version_value[req.GetService().GetVersion().String()])
	curr := s.getJob(req.GetService())

	//Place this in the port map
	if curr == nil {
		curr = req.GetService()
		s.addToPortMap(curr)
	}

	// This is a new registration - update the port map
	curr.LastSeenTime = time.Now().UnixNano()
	return &pb.RegisterResponse{Service: curr}, nil
}

var (
	legacyPeer = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "discovery_peer",
		Help: "Push Size",
	}, []string{"method", "peer"})
)

// Discover supports the Discover rpc end point
func (s *Server) Discover(ctx context.Context, req *pb.DiscoverRequest) (*pb.DiscoverResponse, error) {
	pr, _ := peer.FromContext(ctx)
	s.discoverPeer = fmt.Sprintf("%+v", pr)
	//legacyPeer.With(prometheus.Labels{"method": "Discover", "peer": fmt.Sprintf("%v", strings.Split(pr.Addr.String(), ":")[0])}).Inc()

	//Reject requests without caller
	if req.Caller == "" {
		pr, _ := peer.FromContext(ctx)
		s.peerFail = fmt.Sprintf("%+v", pr)
		return nil, fmt.Errorf("Must specify caller")
	}

	if val, ok := s.version.Load(req.GetRequest().GetName()); ok {
		if val.(int32) == 1 {
			resp, err := s.Get(ctx, &pb.GetRequest{Job: req.GetRequest().GetName(), Server: req.GetRequest().GetIdentifier()})
			if len(resp.GetServices()) > 0 {
				return &pb.DiscoverResponse{Service: resp.GetServices()[0]}, err
			}
			return &pb.DiscoverResponse{}, err
		}
	}
	s.countDiscover++
	in := req.GetRequest()

	// Check if we've been asked for something specific
	if in.GetIdentifier() != "" && in.GetName() != "" {
		in.Port = s.hashPortNumber(in.GetIdentifier(), in.GetName(), SEP)
		val := s.getCurr(in)
		if val != nil {
			return &pb.DiscoverResponse{Service: val}, nil
		}

		return nil, fmt.Errorf("Unable to locate %v on server %v", in.GetName(), in.GetIdentifier())
	}

	return &pb.DiscoverResponse{}, status.Error(codes.Unavailable, fmt.Sprintf("Cannot find master for "+in.GetName()+" on server "+in.GetIdentifier()))
}

//State gets the state of the server
func (s *Server) State(ctx context.Context, in *pb.StateRequest) (*pb.StateResponse, error) {
	return &pb.StateResponse{}, nil
}
