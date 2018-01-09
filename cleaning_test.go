package main

import (
	"context"
	"log"
	"net"
	"strconv"
	"testing"
	"time"

	"google.golang.org/grpc"

	pb "github.com/brotherlogic/discovery/proto"
)

// Get - return free open TCP port
func Get() (port string, err error) {
	ln, err := net.Listen("tcp", "[::]:0")
	if err != nil {
		return "", err
	}
	portn := ln.Addr().(*net.TCPAddr).Port
	err = ln.Close()
	return ":" + strconv.Itoa(portn), err
}

func runServer(testPort string) *grpc.Server {
	s := Serve(testPort)

	// Give the binary 10 seconds to become alive
	time.Sleep(time.Second * 2)

	log.Printf("Listening on port %v", testPort)

	return s
}

func register(port string, binary string, timeToLive int64) error {
	log.Printf("Registering on port %v", port)
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		return err
	}

	client := pb.NewDiscoveryServiceClient(conn)
	resp, err := client.RegisterService(context.Background(), &pb.RegistryEntry{Name: binary, Identifier: "testing", TimeToClean: timeToLive})

	log.Printf("Got register response: %v", resp)

	return err
}

func list(port string) (*pb.ServiceList, error) {
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		return nil, err
	}

	client := pb.NewDiscoveryServiceClient(conn)
	return client.ListAllServices(context.Background(), &pb.Empty{})
}

func TestNormalClean(t *testing.T) {
	port, err2 := Get()
	if err2 != nil {
		log.Fatalf("Unable to find port: %v", err2)
	}

	s := runServer(port)

	err := register(port, "test-binary", -1)
	if err != nil {
		t.Fatalf("Unable to register: %v", err)
	}

	//Wait 1 second (no cleaning should happen)
	time.Sleep(time.Second)
	servers, err := list(port)
	if err != nil {
		t.Fatalf("Unable to list: %v", err)
	}
	if len(servers.GetServices()) != 1 {
		t.Fatalf("Error running test, server is not registered: %v", servers)
	}

	//Wait for 4 seconds to allow time for cleaning
	time.Sleep(time.Second * 3)

	servers, err = list(port)
	if err != nil {
		t.Fatalf("Unable to list: %v", err)
	}
	if len(servers.Services) > 0 {
		t.Errorf("Server has not been cleaned after 4 seconds: %v", servers)
	}

	s.Stop()
}

func TestLongClean(t *testing.T) {
	port, err2 := Get()
	if err2 != nil {
		log.Fatalf("Unable to find port: %v", err2)
	}
	s := runServer(port)

	err := register(port, "test-binary", 5000)
	if err != nil {
		t.Fatalf("Unable to register: %v", err)
	}

	//Wait 1 second (no cleaning should happen)
	time.Sleep(time.Second)
	servers, err := list(port)
	if err != nil {
		t.Fatalf("Unable to list: %v", err)
	}
	if len(servers.GetServices()) != 1 {
		t.Fatalf("Error running test, server is not registered: %v", servers)
	}

	//Wait for 4 seconds and still should not be cleaned
	time.Sleep(time.Second * 3)

	servers, err = list(port)
	if err != nil {
		t.Fatalf("Unable to list: %v", err)
	}
	if len(servers.Services) != 1 {
		t.Errorf("Server has been cleaned after 4 seconds: %v", servers)
	}

	//Wait for 4 seconds to allow time for cleaning
	time.Sleep(time.Second * 2)

	servers, err = list(port)
	if err != nil {
		t.Fatalf("Unable to list: %v", err)
	}
	if len(servers.Services) > 0 {
		t.Errorf("Server has not been cleaned after 5 seconds: %v", servers)
	}

	s.Stop()
}
