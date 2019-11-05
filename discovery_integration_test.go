package main

import (
	"context"
	"testing"

	pb "github.com/brotherlogic/discovery/proto"
)

func TestUnregisterRun(t *testing.T) {
	s := InitTestServer()

	_, err := s.RegisterService(context.Background(), &pb.RegisterRequest{Service: &pb.RegistryEntry{Name: "test", Identifier: "jazz", TimeToClean: 100}})
	if err != nil {
		t.Fatalf("Bad register: %v", err)
	}

	r, err := s.Discover(context.Background(), &pb.DiscoverRequest{Caller: "me", Request: &pb.RegistryEntry{Name: "test"}})
	if err == nil {
		t.Fatalf("We were able to discover with no master: %v", r)
	}

	_, err = s.RegisterService(context.Background(), &pb.RegisterRequest{Service: &pb.RegistryEntry{Name: "test", Identifier: "jazz", TimeToClean: 100, Master: true}})
	if err != nil {
		t.Fatalf("Bad register as master: %v", err)
	}

	r, err = s.Discover(context.Background(), &pb.DiscoverRequest{Caller: "me", Request: &pb.RegistryEntry{Name: "test"}})
	if err != nil {
		t.Fatalf("Discovery issue: %v", err)
	}

	_, err = s.Unregister(context.Background(), &pb.UnregisterRequest{Service: &pb.RegistryEntry{Name: "test", Identifier: "jazz"}})
	if err != nil {
		t.Fatalf("Error on unregister")
	}

	_, err = s.RegisterService(context.Background(), &pb.RegisterRequest{Service: &pb.RegistryEntry{Name: "test", Identifier: "jazz2", TimeToClean: 100}})
	if err != nil {
		t.Fatalf("Bad register on jazz2: %v", err)
	}

	r, err = s.Discover(context.Background(), &pb.DiscoverRequest{Caller: "me", Request: &pb.RegistryEntry{Name: "test"}})
	if err == nil {
		t.Fatalf("We were able to discover on jazz2: %v", r)
	}

	_, err = s.RegisterService(context.Background(), &pb.RegisterRequest{Service: &pb.RegistryEntry{Name: "test", Identifier: "jazz2", TimeToClean: 100, Master: true}})
	if err != nil {
		t.Fatalf("Bad register as master on jazz2: %v", err)
	}

	r, err = s.Discover(context.Background(), &pb.DiscoverRequest{Caller: "me", Request: &pb.RegistryEntry{Name: "test"}})
	if err != nil {
		t.Fatalf("Discovery issue: %v", err)
	}

}
