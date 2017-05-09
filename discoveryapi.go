package main

import (
	"context"
	"log"
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

func main() {
	Serve()
}
