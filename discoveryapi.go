package main

import (
	"context"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"

	pb "github.com/brotherlogic/discovery/proto"
	pbg "github.com/brotherlogic/goserver/proto"
)

type prodHealthChecker struct{}

func (healthChecker prodHealthChecker) Check(entry *pb.RegistryEntry) bool {
	conn, err := grpc.Dial(entry.Ip+":"+strconv.Itoa(int(entry.Port)), grpc.WithInsecure())
	if err != nil {
		log.Printf("Can't event dial %v -> %v", entry, err)
	}
	defer conn.Close()

	registry := pbg.NewGoserverServiceClient(conn)
	_, err = registry.IsAlive(context.Background(), &pbg.Alive{})
	if err != nil {
		log.Printf("Error reading health of %v -> %v", entry, err)
		return false
	}
	return true
}

// Serve main server function
func Serve() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Unable to get tcp port %v -> %v", port, err)
	}
	s := grpc.NewServer()
	server := InitServer()
	server.loadCheckFile("checkfile")
	pb.RegisterDiscoveryServiceServer(s, &server)
	err = s.Serve(lis)
	log.Printf("Failed to serve: %v", err)
}

func main() {
	Serve()
}
