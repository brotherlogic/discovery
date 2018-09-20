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
	return &pb.ListResponse{Services: &pb.ServiceList{Services: s.entries}}, nil
}

// RegisterService supports the RegisterService rpc end point
func (s *Server) RegisterService(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	in := req.GetService()

	if in.ExternalPort {
		in.Ip = s.getExternalIP(prodHTTPGetter{})
	}

	//Validate us if we're trying to use a port number already used
	for _, en := range s.entries {
		if en.Port == in.Port && en.Name != in.Name && en.Ip == in.Ip {
			return nil, fmt.Errorf("This port is already registered")
		}
	}

	s.countM.Lock()
	if _, ok := s.counts[in.GetName()]; !ok {
		s.counts[in.GetName()] = 0
	}
	s.counts[in.GetName()]++
	s.countM.Unlock()

	in.LastSeenTime = time.Now().Unix()

	s.mm.Lock()
	if val, ok := s.masterMap[in.GetName()]; ok && val.GetIdentifier() == in.GetIdentifier() && !in.GetMaster() {
		delete(s.masterMap, in.GetName())
		if in.MasterTime > 0 {
			in.MasterTime = 0
		}
	}
	s.mm.Unlock()

	// Adjust the clean time if necessary (default to 5 seconds)
	if in.GetTimeToClean() == 0 {
		in.TimeToClean = 1000 * 5
	}

	// Get the new port number
	for portNumber := int32(50055 + 1); in.Port == 0 && portNumber < 60000; portNumber++ {
		taken := false
		for _, service := range s.entries {
			if service.Port == portNumber {
				taken = true
			}
		}
		if !taken {
			in.Port = portNumber
		}

	}

	// If we've already registered this service, return immediately
	for _, service := range s.entries {
		if in.GetRegisterTime() > 0 && service.Identifier == in.Identifier && service.Name == in.Name && service.Port == in.Port {
			// Add to master map if this is master
			if in.GetMaster() {
				s.mm.Lock()
				if val, ok := s.masterMap[in.GetName()]; ok {
					if val.Identifier != in.Identifier && ((time.Now().Unix()-val.GetLastSeenTime())*1000 < val.GetTimeToClean()) {
						s.mm.Unlock()
						return nil, fmt.Errorf("Unable to register as master - already exists(%v) -> %v", val, in)
					}
				} else {
					service.MasterTime = time.Now().Unix()
				}

				s.masterMap[in.GetName()] = service
				s.mm.Unlock()
			}

			//Refresh the IP and store the checkfile
			service.Ip = in.Ip
			service.Master = in.Master
			service.Port = in.Port
			service.LastSeenTime = time.Now().Unix()
			return &pb.RegisterResponse{Service: service}, nil
		}
	}

	in.RegisterTime = time.Now().Unix()

	//Reject any master registrations that are new
	if in.GetMaster() {
		return nil, fmt.Errorf("Unable to register as master (%v)", in)
	}

	s.entries = append(s.entries, in)
	return &pb.RegisterResponse{Service: in}, nil
}

// Discover supports the Discover rpc end point
func (s *Server) Discover(ctx context.Context, req *pb.DiscoverRequest) (*pb.DiscoverResponse, error) {
	in := req.GetRequest()
	var nonmaster *pb.RegistryEntry
	for _, entry := range s.entries {
		if entry.Name == in.GetName() && (in.GetIdentifier() == "" || in.GetIdentifier() == entry.Identifier) {
			if entry.Master || in.Identifier != "" {
				return &pb.DiscoverResponse{Service: entry}, nil
			}
			nonmaster = entry
		}
	}

	//Return the non master if possible
	if nonmaster != nil {
		return nil, errors.New("Cannot find a master for service called " + in.GetName() + " on server (maybe): " + in.GetIdentifier())
	}

	return &pb.DiscoverResponse{}, errors.New("Cannot find service called " + in.GetName() + " on server (maybe): " + in.GetIdentifier())
}

//State gets the state of the server
func (s *Server) State(ctx context.Context, in *pb.StateRequest) (*pb.StateResponse, error) {
	s.countM.Lock()
	resp := fmt.Sprintf("Counts: %v", s.counts)
	s.countM.Unlock()
	return &pb.StateResponse{Counts: resp, Len: int32(len(s.entries))}, nil
}
