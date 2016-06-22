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
