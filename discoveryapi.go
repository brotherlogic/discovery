package main

import (
	"context"
	"flag"
	"io/ioutil"
	"log"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc"

	pb "github.com/brotherlogic/discovery/proto"
	pbg "github.com/brotherlogic/goserver/proto"
)

type prodHealthChecker struct{}

func (healthChecker prodHealthChecker) Check(entry *pb.RegistryEntry) bool {
	log.Printf("Dialing for health: %v", entry)
	conn, err := grpc.Dial(entry.Ip+":"+strconv.Itoa(int(entry.Port)), grpc.WithTimeout(time.Second*5), grpc.WithInsecure())
	if err != nil {
		log.Printf("Can't event dial %v -> %v", entry, err)
		return false
	}
	defer conn.Close()

	registry := pbg.NewGoserverServiceClient(conn)
	log.Printf("Asking if %v is alive", entry)
	_, err = registry.IsAlive(context.Background(), &pbg.Alive{})
	if err != nil {
		log.Printf("Error reading health of %v -> %v", entry, err)
		return false
	}
	log.Printf("Got the answer (nil is good): %v", err)
	return true
}

// Serve main server function
func Serve() {
	var quiet = flag.Bool("quiet", false, "Show all output")
	flag.Parse()
	if *quiet {
		log.SetFlags(0)
		log.SetOutput(ioutil.Discard)
	}
	log.Printf("Logging is on!")

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
