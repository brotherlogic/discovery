package main

import (
	"errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"

	pb "github.com/brotherlogic/discovery/proto"
)

const (
	port = ":50055"
)

var externalPorts = map[string][]int32{"10.0.1.17": []int32{50052, 50053}}

// Server the central server object
type Server struct {
	entries []pb.RegistryEntry
}

// InitServer builds a server item ready for use
func InitServer() Server {
	s := Server{}
	s.entries = make([]pb.RegistryEntry, 0)
	return s
}

// RegisterService supports the RegisterService rpc end point
func (s *Server) RegisterService(ctx context.Context, in *pb.RegistryEntry) (*pb.RegistryEntry, error) {

	// Server is requesting an external port
	if in.ExternalPort {
		availablePorts := externalPorts[in.Ip]
		taken := false
		for _, port := range availablePorts {
			for _, service := range s.entries {
				if service.Ip == in.Ip && service.Port == port {
					taken = true
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
	}

	s.entries = append(s.entries, *in)
	return in, nil
}

// Discover supports the Discover rpc end point
func (s *Server) Discover(ctx context.Context, in *pb.RegistryEntry) (*pb.RegistryEntry, error) {
	for _, entry := range s.entries {
		if entry.Name == in.Name {
			return &entry, nil
		}
	}

	return &pb.RegistryEntry{}, errors.New("Cannot find service")
}

// Serve main server function
func Serve() {
	lis, _ := net.Listen("tcp", port)
	s := grpc.NewServer()
	server := InitServer()
	pb.RegisterDiscoveryServiceServer(s, &server)
	s.Serve(lis)
}
