package main

import (
	"context"
	"testing"
	"time"

	pb "github.com/brotherlogic/discovery/proto"
)

func TestClean(t *testing.T) {
	s := InitTestServer()

	entry := &pb.RegistryEntry{Ip: "10.0.1.17", Identifier: "Server1", Name: "Job1", TimeToClean: 100}
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
