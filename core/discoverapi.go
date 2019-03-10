package discovery

import (
	"errors"
	"fmt"
	"time"

	"golang.org/x/net/context"

	pb "github.com/brotherlogic/discovery/proto"
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

func (s *Server) setPortNumber(in *pb.RegistryEntry) error {
	if in.Port == 0 {
		if in.ExternalPort {
			in.Port = 50053
		} else {
			in.Port = s.hashPortNumber(in.Identifier, in.Name)
		}
	}

	return nil
}

func (s *Server) setupPort(in *pb.RegistryEntry) {
	if in.Port == 0 {
		in.RegisterTime = time.Now().UnixNano()
		if in.ExternalPort {
			in.Ip = s.getExternalIP(prodHTTPGetter{})
		}

		s.setPortNumber(in)
	}
}

func (s *Server) getCurr(in *pb.RegistryEntry) *pb.RegistryEntry {
	return s.portMap[in.Port-50052]
}

func (s *Server) getCMaster(in *pb.RegistryEntry) *pb.RegistryEntry {
	s.mm.RLock()
	defer s.mm.RUnlock()
	return s.masterMap[in.GetName()]
}

func (s *Server) getJob(in *pb.RegistryEntry) (*pb.RegistryEntry, *pb.RegistryEntry) {
	// Setup the port info
	s.setupPort(in)

	return s.getCurr(in), s.getCMaster(in)
}

func (s *Server) addMaster(in *pb.RegistryEntry) {
	//Register as master if there is none
	in.MasterTime = time.Now().UnixNano()
	s.mm.Lock()
	s.masterMap[in.GetName()] = in
	s.mm.Unlock()
}

func (s *Server) removeMaster(in *pb.RegistryEntry) {
	// Remove if we're re-registering without master
	s.mm.Lock()
	delete(s.masterMap, in.GetName())
	s.mm.Unlock()

}

func (s *Server) addToPortMap(in *pb.RegistryEntry) {
	s.portMap[in.Port-50052] = in
}

// RegisterService supports the RegisterService rpc end point
func (s *Server) RegisterService(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	s.countRegister++

	//Reject the request with no time to clean
	if req.GetService().GetTimeToClean() == 0 {
		return &pb.RegisterResponse{}, fmt.Errorf("You must specify a clean time")
	}

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

		if master == nil || master.WeakMaster {
			curr.Master = true
			s.addMaster(curr)
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

	// This is a new registration - update the port map
	curr.LastSeenTime = time.Now().UnixNano()
	return &pb.RegisterResponse{Service: curr}, nil
}

// Discover supports the Discover rpc end point
func (s *Server) Discover(ctx context.Context, req *pb.DiscoverRequest) (*pb.DiscoverResponse, error) {
	s.countDiscover++
	in := req.GetRequest()

	// Check if we've been asked for something specific
	if in.GetIdentifier() != "" && in.GetName() != "" {
		in.Port = s.hashPortNumber(in.GetIdentifier(), in.GetName())
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
