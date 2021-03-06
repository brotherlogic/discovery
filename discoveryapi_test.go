package main

import (
	"context"
	"testing"
	"time"

	pb "github.com/brotherlogic/discovery/proto"
)

func TestRedirect(t *testing.T) {
	s := InitTestServer()
	entry := &pb.RegistryEntry{Ip: "10.0.1.17", Identifier: "Server1", Name: "Job1", TimeToClean: 100}
	_, err := s.RegisterService(context.Background(), &pb.RegisterRequest{Service: entry})

	if err != nil {
		t.Fatalf("Error registering service: %v", err)
	}

	val, err := s.Get(context.Background(), &pb.GetRequest{Job: "Job1", Server: "Server1"})
	if err != nil || len(val.GetServices()) != 1 {
		t.Errorf("Master discover has succeeded: %v, %v -> %v", err, len(val.GetServices()), val)
	}
}

func TestClean(t *testing.T) {
	s := InitTestServer()

	entry := &pb.RegistryEntry{Ip: "10.0.1.17", Identifier: "Server1", Name: "Job1", TimeToClean: 100, Master: true}
	_, err := s.RegisterService(context.Background(), &pb.RegisterRequest{Service: entry})

	if err != nil {
		t.Fatalf("Error registering service: %v", err)
	}

	time.Sleep(time.Second)

	val, err := s.Discover(context.Background(), &pb.DiscoverRequest{Caller: "yep", Request: &pb.RegistryEntry{Name: "Job1"}})
	if err == nil {
		t.Errorf("Discover succeeded: %v", val)
	}
}
