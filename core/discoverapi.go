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
	s.portMapMutex.RLock()
	defer s.portMapMutex.RUnlock()
	for _, val := range s.portMap {
		entries = append(entries, val)
	}
	return &pb.ListResponse{Services: &pb.ServiceList{Services: entries}}, nil
}

func (s *Server) updateCounts(in *pb.RegistryEntry) {
	s.countM.Lock()
	s.counts[in.GetName()]++
	s.countM.Unlock()
}

func (s *Server) setPortNumber(in *pb.RegistryEntry) error {
	if in.Port == 0 {
		if in.ExternalPort {
			in.Port = 50053
		} else {
			in.Port = s.hashPortNumber(in.Identifier, in.Name)
		}
	} else if !in.ExternalPort {
		if s.hashPortNumber(in.Identifier, in.Name) != in.Port {
			return fmt.Errorf("Unable to register %v under %v port is %v but it should be %v", in.Name, in.Identifier, in.Port, s.hashPortNumber(in.Identifier, in.Name))
		}
	}

	return nil
}

func (s *Server) setOrKeepMaster(in *pb.RegistryEntry) error {
	seenMaster := false
	s.mm.RLock()
	val, ok := s.masterMap[in.GetName()]
	s.mm.RUnlock()
	if ok {
		seenMaster = true
		// Someone else is master if they have a lease and it has not expired yet
		if val.Identifier != in.Identifier && val.LastSeenTime+val.TimeToClean*1000000 >= time.Now().UnixNano() {
			return fmt.Errorf("Unable to register as master - already exists(%v) -> %v also %v", val.Identifier, in.Identifier, val.LastSeenTime+val.TimeToClean*1000000-time.Now().UnixNano())
		}
	}

	// If this is a new master - set the master time
	if !seenMaster {
		in.MasterTime = time.Now().UnixNano()
	}
	s.mm.Lock()
	s.masterMap[in.GetName()] = in
	s.mm.Unlock()
	return nil
}

func (s *Server) clearMaster(in *pb.RegistryEntry) {
	s.mm.RLock()
	val, ok := s.masterMap[in.GetName()]
	s.mm.RUnlock()
	if ok && val.Identifier == in.Identifier {
		s.mm.Lock()
		delete(s.masterMap, in.GetName())
		s.mm.Unlock()
	}
}

// RegisterService supports the RegisterService rpc end point
func (s *Server) RegisterService(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	s.countRegister++
	in := req.GetService()

	//Reject the request with no time to clean
	if in.GetTimeToClean() == 0 {
		return &pb.RegisterResponse{}, fmt.Errorf("You must specify a clean time")
	}

	s.portMapMutex.RLock()
	val, ok := s.portMap[in.Port]
	s.portMapMutex.RUnlock()
	if ok {
		//If we have a port number - update the last seen time
		if in.GetMaster() {
			err := s.setOrKeepMaster(in)
			if err != nil {
				return nil, err
			}
		} else {
			s.clearMaster(val)
			val.MasterTime = 0
			val.LastSeenTime = time.Now().UnixNano()
			return &pb.RegisterResponse{Service: val}, nil
		}
	}

	//We're now dealing with a new registration

	// Set the port information up front
	if in.ExternalPort {
		in.Ip = s.getExternalIP(prodHTTPGetter{})
	}

	err := s.setPortNumber(in)
	if err != nil {
		return nil, err
	}

	if in.RegisterTime == 0 {
		in.RegisterTime = time.Now().UnixNano()
	}

	s.updateCounts(in)

	//Deal with request to be master
	if in.GetMaster() {
		err := s.setOrKeepMaster(in)
		if err != nil {
			return nil, err
		}
	} else {
		s.clearMaster(in)
		in.MasterTime = 0
	}

	// This is a new registration - update the port map
	in.LastSeenTime = time.Now().UnixNano()
	s.portMapMutex.Lock()
	s.portMap[in.Port] = in
	s.portMapMutex.Unlock()
	return &pb.RegisterResponse{Service: in}, nil
}

// Discover supports the Discover rpc end point
func (s *Server) Discover(ctx context.Context, req *pb.DiscoverRequest) (*pb.DiscoverResponse, error) {
	s.countDiscover++
	in := req.GetRequest()

	// Check if we've been asked for something specific
	if in.GetIdentifier() != "" && in.GetName() != "" {
		port := s.hashPortNumber(in.GetIdentifier(), in.GetName())
		s.portMapMutex.RLock()
		defer s.portMapMutex.RUnlock()
		if val, ok := s.portMap[port]; ok {
			return &pb.DiscoverResponse{Service: val}, nil
		}

		return nil, fmt.Errorf("Unable to locate %v on server %v", in.GetName(), in.GetIdentifier())
	}

	// Return the master if it exists
	s.mm.RLock()
	val, ok := s.masterMap[in.GetName()]
	s.mm.RUnlock()
	if ok && val.LastSeenTime+val.TimeToClean*1000000 > time.Now().UnixNano() {
		return &pb.DiscoverResponse{Service: val}, nil
	}

	return &pb.DiscoverResponse{}, errors.New("Cannot find master for " + in.GetName() + " on server (maybe): " + in.GetIdentifier())
}

//State gets the state of the server
func (s *Server) State(ctx context.Context, in *pb.StateRequest) (*pb.StateResponse, error) {
	s.countM.Lock()
	longest := ""
	longestCount := 0
	for name, number := range s.counts {
		if number > longestCount {
			longestCount = number
			longest = name
		}
	}

	s.countM.Unlock()
	return &pb.StateResponse{MostFrequent: longest, Frequency: int32(longestCount), LongestCall: s.longest, Count: fmt.Sprintf("D %v, R %v, L %v", s.countDiscover, s.countRegister, s.countList)}, nil
}
