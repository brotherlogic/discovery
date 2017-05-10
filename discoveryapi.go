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
	conn, _ := grpc.Dial(entry.Ip+":"+strconv.Itoa(int(entry.Port)), grpc.WithInsecure())
	defer conn.Close()

	registry := pbg.NewGoserverServiceClient(conn)
	_, err := registry.IsAlive(context.Background(), &pbg.Alive{})
	if err != nil {
		log.Printf("Error reading health of %v", entry)
		return false
	}
	return true
}

// Serve main server function
func Serve() {
	lis, _ := net.Listen("tcp", port)
	s := grpc.NewServer()
	server := InitServer()
	server.loadCheckFile("checkfile")
	pb.RegisterDiscoveryServiceServer(s, &server)
	err := s.Serve(lis)
	log.Printf("Failed to serve: %v", err)
}

func main() {
	Serve()
}
