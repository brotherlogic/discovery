package main

import (
	"testing"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/brotherlogic/discovery/proto"
)

func TestRegisterService(t *testing.T) {
	s := InitServer()
	entry := &pb.RegistryEntry{Ip: "10.0.4.5", Port: 50051, Name: "Testing"}
	r, err := s.RegisterService(context.Background(), entry)
	if err != nil {
		t.Errorf("Error registering service: %v", err)
	}

	if r.Name != entry.Name {
		t.Errorf("Problem with name resolution %v vs %v", r.Name, entry.Name)
	}
}

func TestFailedDiscover(t *testing.T) {
	s := InitServer()

	entry := &pb.RegistryEntry{Name: "Testing"}
	_, err := s.Discover(context.Background(), entry)
	if err == nil {
		t.Errorf("Disoovering non existing service did not fail: %v", err)
	}
}

func TestDiscover(t *testing.T) {
	s := InitServer()
	entryAdd := &pb.RegistryEntry{Ip: "10.0.4.5", Port: 50051, Name: "Testing"}
	s.RegisterService(context.Background(), entryAdd)
	entry := &pb.RegistryEntry{Name: "Testing"}
	r, err := s.Discover(context.Background(), entry)
	if err != nil {
		t.Errorf("Error registering service: %v", err)
	}

	if r.Ip != entryAdd.Ip {
		t.Errorf("Discovery process failed %v vs %v", r.Ip, entryAdd.Ip)
	}
}

func TestRunServer(t *testing.T) {
	Serve()

	go func() {
		conn, err := grpc.Dial("localhost:50055", grpc.WithInsecure())
		if err != nil {
			t.Errorf("Error connecting to port")
		}

		defer conn.Close()
		client := pb.NewDiscoveryServiceClient(conn)

		entry := pb.RegistryEntry{}

		_, err = client.RegisterService(context.Background(), &entry)
		if err != nil {
			t.Errorf("Error registering service: %v", err)
		}

		_, err = client.Discover(context.Background(), &entry)
		if err != nil {
			t.Errorf("Error performing discovery: %v", err)
		}

	}()

	time.Sleep(10 * time.Second)
}

func TestMainForCoverage(t *testing.T) {
	main()
}
