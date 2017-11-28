package main

import (
	"errors"
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
					//Refresh the IP and store the checkfile
					service.Ip = in.Ip
					service.Master = in.Master
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
					//Refresh the IP and store the checkfile
					service.Ip = in.Ip
					service.Master = in.Master
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

	s.recordTime("Register-New", time.Now().Sub(t))
	s.entries = append(s.entries, in)
	return in, nil
}
