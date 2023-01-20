package main

import (
	"context"
	"testing"

	pb "github.com/brotherlogic/discovery/proto"
)

func TestRegisterV2(t *testing.T) {
	s := InitTestServer()

	resp, err := s.RegisterV2(context.Background(), &pb.RegisterRequest{Service: &pb.RegistryEntry{Name: "test_job", Identifier: "test_server", Version: pb.RegistryEntry_V2}})

	if err != nil {
		t.Errorf("Unable to register %v", err)
	}

	if resp.Service.Port == 0 {
		t.Errorf("Port number not assigned")
	}

	s.friends = append(s.friends, "blah")
	respg, err := s.Get(context.Background(), &pb.GetRequest{Job: "test_job", Server: "test_server", Friend: "madeup"})
	if err != nil {
		t.Errorf("Unable to get %v", err)
	}

	if len(respg.Services) != 1 {
		t.Errorf("Service not returned")
	}
}

func TestRegisterV2NoFriend(t *testing.T) {
	s := InitTestServer()

	resp, err := s.RegisterV2(context.Background(), &pb.RegisterRequest{Service: &pb.RegistryEntry{Name: "test_job", Identifier: "test_server", Version: pb.RegistryEntry_V2}})

	if err != nil {
		t.Errorf("Unable to register %v", err)
	}

	if resp.Service.Port == 0 {
		t.Errorf("Port number not assigned")
	}

	s.friends = append(s.friends, "madeup")
	respg, err := s.Get(context.Background(), &pb.GetRequest{Job: "test_job", Server: "test_server", Friend: "madeup"})
	if err != nil {
		t.Errorf("Unable to get %v", err)
	}

	if len(respg.Services) != 1 {
		t.Errorf("Service not returned")
	}
}
