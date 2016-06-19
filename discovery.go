package main

import(
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"	

	pb "github.com/brotherlogic/discovery/proto"
)

const (
      port = ":50052"
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
func (s *Server) RegisterService(ctx context.Context, in *pb.RegistryEntry) (*pb.RegistryEntry, error){
     return &pb.RegistryEntry{}, nil
}

// Discover supports the Discover rpc end point
func (s *Server) Discover(ctx context.Context, in *pb.RegistryEntry) (*pb.RegistryEntry, error){
     return &pb.RegistryEntry{}, nil
}

func main() {
     lis, err := net.Listen("tcp", port)
     if err != nil {
       log.Fatalf("Failed to listen on %v", err)
     }

     s := grpc.NewServer()
     server := InitServer()
     pb.RegisterDiscoveryServiceServer(s, &server)
     s.Serve(lis)
}