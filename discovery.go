package main

import (
	"errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"

	pb "github.com/brotherlogic/discovery/proto"
)

const (
	port = ":50055"
)

var externalPorts = map[string][]int32{"10.0.1.17": []int32{50052, 50053}}

// Server the central server object
type Server struct {
	entries []*pb.RegistryEntry
}

// InitServer builds a server item ready for useo
func InitServer() Server {
	s := Server{}
	s.entries = make([]*pb.RegistryEntry, 0)
	return s
}

// ListAllServices returns a list of all the services
func (s *Server) ListAllServices(ctx context.Context, in *pb.Empty) (*pb.ServiceList, error) {
	return &pb.ServiceList{Services: s.entries}, nil
}

// RegisterService supports the RegisterService rpc end point
func (s *Server) RegisterService(ctx context.Context, in *pb.RegistryEntry) (*pb.RegistryEntry, error) {
	// Server is requesting an external port
	if in.ExternalPort {
		availablePorts := externalPorts[in.Ip]
		for _, port := range availablePorts {
			taken := false
			for _, service := range s.entries {
				if service.Ip == in.Ip && service.Port == port {
					log.Printf("TAKEN")
					taken = true
				}
				// If we've already registered this service, return immediately
				if service.Identifier == in.Identifier && service.Name == in.Name {
					return service, nil
				}

			}

			if !taken {
				in.Port = port
				break
			}
		}

		//Throw an error if we can't find a port number
		if in.Port <= 0 {
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
					return service, nil
				}
			}
			if !taken {
				in.Port = portNumber
				break
			}
		}
	}

	s.entries = append(s.entries, in)
	return in, nil
}

// Discover supports the Discover rpc end point
func (s *Server) Discover(ctx context.Context, in *pb.RegistryEntry) (*pb.RegistryEntry, error) {
	for _, entry := range s.entries {
		if entry.Name == in.Name {
			return entry, nil
		}
	}

	return &pb.RegistryEntry{}, errors.New("Cannot find service called " + in.Name)
}

// Serve main server function
func Serve() {
	lis, _ := net.Listen("tcp", port)
	s := grpc.NewServer()
	server := InitServer()
	pb.RegisterDiscoveryServiceServer(s, &server)
	s.Serve(lis)
}
