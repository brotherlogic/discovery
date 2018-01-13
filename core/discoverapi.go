package discovery

import (
	"errors"
	"fmt"
	"time"

	"golang.org/x/net/context"

	pb "github.com/brotherlogic/discovery/proto"
)

// ListAllServices returns a list of all the services
func (s *Server) ListAllServices(ctx context.Context, in *pb.Empty) (*pb.ServiceList, error) {
	t := time.Now()
	s.recordTime("ListAllServices", time.Now().Sub(t))
	return &pb.ServiceList{Services: s.entries}, nil
}

// RegisterService supports the RegisterService rpc end point
func (s *Server) RegisterService(ctx context.Context, in *pb.RegistryEntry) (*pb.RegistryEntry, error) {
	t := time.Now()
	in.LastSeenTime = t.Unix()

	s.mm.Lock()
	if _, ok := s.masterMap[in.GetName()]; ok && !in.GetMaster() {
		delete(s.masterMap, in.GetName())

	}
	s.mm.Unlock()

	// Adjust the clean time if necessary (default to 3 seconds)
	if in.GetTimeToClean() == 0 {
		in.TimeToClean = 1000 * 3
	}

	// Server is requesting an external port
	if in.ExternalPort {
		availablePorts := externalPorts["main"]
		// Reset the request IP to an external IP
		in.Ip = s.getExternalIP(prodHTTPGetter{})

		for _, port := range availablePorts {
			taken := false
			for _, service := range s.entries {
				if service.Ip == in.Ip && service.Port == port {
					taken = true
				}
				// If we've already registered this service, return immediately
				if service.Identifier == in.Identifier && service.Name == in.Name {

					// Add to master map if this is master
					if in.GetMaster() {
						if val, ok := s.masterMap[in.GetName()]; ok && val.Identifier != in.Identifier {
							return nil, fmt.Errorf("Unable to register as master - already exists(%v) -> %v", val, in)
						}
						s.mm.Lock()
						s.masterMap[in.GetName()] = service
						s.mm.Unlock()
					}

					//Refresh the IP and store the checkfile
					service.Ip = in.Ip
					service.Master = in.Master
					service.LastSeenTime = time.Now().Unix()
					s.recordTime("Register-External-Found", time.Now().Sub(t))
					return service, nil
				}
			}

			if !taken {
				in.Port = port
				in.RegisterTime = time.Now().Unix()
				break
			}
		}

		//Throw an error if we can't find a port number
		if in.Port <= 0 {
			s.recordTime("Register-External-NoAllocate", time.Now().Sub(t))
			return in, errors.New("Unable to allocate external port")
		}
	} else {
		var portNumber int32
		for portNumber = 50055 + 1; portNumber < 60000; portNumber++ {
			taken := false
			for _, service := range s.entries {
				if service.Port == portNumber {
					taken = true
				}

				// If we've already registered this service, return immediately
				if service.Identifier == in.Identifier && service.Name == in.Name {

					// Add to master map if this is master
					if in.GetMaster() {
						s.mm.Lock()
						if val, ok := s.masterMap[in.GetName()]; ok && val.Identifier != in.Identifier {
							return nil, fmt.Errorf("Unable to register as master - already exists(%v) -> %v", val, in)
						}
						s.masterMap[in.GetName()] = service
						s.mm.Unlock()
					}

					//Refresh the IP and store the checkfile
					service.Ip = in.Ip
					service.Master = in.Master
					service.LastSeenTime = time.Now().Unix()
					s.recordTime("Register-Internal-Found", time.Now().Sub(t))
					return service, nil
				}
			}
			if !taken {
				//Only set the port if it's not set
				if in.Port <= 0 {
					in.Port = portNumber
					in.RegisterTime = time.Now().Unix()
				}
				break
			}
		}
	}

	//Reject any master registrations that are new
	if in.GetMaster() {
		return nil, fmt.Errorf("Unable to register as master (%v)", in)
	}

	s.recordTime("Register-New", time.Now().Sub(t))
	s.entries = append(s.entries, in)
	return in, nil
}

// Discover supports the Discover rpc end point
func (s *Server) Discover(ctx context.Context, in *pb.RegistryEntry) (*pb.RegistryEntry, error) {
	t := time.Now()
	var nonmaster *pb.RegistryEntry
	for _, entry := range s.entries {
		if entry.Name == in.Name && (in.Identifier == "" || in.Identifier == entry.Identifier) {
			if entry.Master || in.Identifier != "" {
				s.recordTime("Discover-foundmaster", time.Now().Sub(t))
				return entry, nil
			}
			nonmaster = entry
		}
	}

	//Return the non master if possible
	if nonmaster != nil {
		s.recordTime("Discover-nonmaster", time.Now().Sub(t))
		return nil, errors.New("Cannot find a master for service called " + in.Name + " on server (maybe): " + in.Identifier)
	}

	s.recordTime("Discover-fail", time.Now().Sub(t))
	return &pb.RegistryEntry{}, errors.New("Cannot find service called " + in.Name + " on server (maybe): " + in.Identifier)
}
