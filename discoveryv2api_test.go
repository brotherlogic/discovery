package main

import (
	"context"
	"testing"

	pb "github.com/brotherlogic/discovery/proto"
)

type testElector struct {
	lastElect string
}

func (p *testElector) elect(ctx context.Context, entry *pb.RegistryEntry) error {
	p.lastElect = entry.Name
	return nil
}

func TestPlainRegisterRun(t *testing.T) {
	s := InitTestServer()

	resp, err := s.RegisterV2(context.Background(), &pb.RegisterRequest{Service: &pb.RegistryEntry{Name: "test_job", Identifier: "test_server"}})

	if err != nil {
		t.Errorf("Unable to register %v", err)
	}

	if resp.Service.Port == 0 {
		t.Errorf("Port number not assigned")
	}

	respg, err := s.Get(context.Background(), &pb.GetRequest{Job: "test_job", Server: "test_server"})
	if err != nil {
		t.Errorf("Unable to get %v", err)
	}

	if len(respg.Services) != 1 {
		t.Errorf("Service not returned")
	}
}

func TestPlainRegisterRunWithBadGet(t *testing.T) {
	s := InitTestServer()

	resp, err := s.RegisterV2(context.Background(), &pb.RegisterRequest{Service: &pb.RegistryEntry{Name: "test_job", Identifier: "test_server"}})

	if err != nil {
		t.Errorf("Unable to register %v", err)
	}

	if resp.Service.Port == 0 {
		t.Errorf("Port number not assigned")
	}

	respg, err := s.Get(context.Background(), &pb.GetRequest{Job: "test_job", Server: "test_serversss"})
	if err == nil {
		t.Errorf("Get did not fail: %v", respg)
	}
}

func TestDoubleRegisterV2(t *testing.T) {
	s := InitTestServer()

	resp, err := s.RegisterV2(context.Background(), &pb.RegisterRequest{Service: &pb.RegistryEntry{Name: "test_job", Identifier: "test_server"}})

	if err != nil {
		t.Errorf("Unable to register %v", err)
	}

	if resp.Service.Port == 0 {
		t.Errorf("Port number not assigned")
	}

	resp, err = s.RegisterV2(context.Background(), &pb.RegisterRequest{Service: &pb.RegistryEntry{Name: "test_job", Identifier: "test_server"}})

	if err == nil {
		t.Errorf("No error on re-register %v", resp)
	}

}

func TestGetAllV2(t *testing.T) {
	s := InitTestServer()

	_, err := s.RegisterV2(context.Background(), &pb.RegisterRequest{Service: &pb.RegistryEntry{Name: "test_job", Identifier: "test_server"}})

	if err != nil {
		t.Errorf("Unable to register %v", err)
	}

	respg, err := s.Get(context.Background(), &pb.GetRequest{})
	if err != nil {
		t.Errorf("Unable to get %v", err)
	}

	if len(respg.Services) != 1 {
		t.Errorf("Service not returned")
	}
}

func TestGetMaster(t *testing.T) {
	s := InitTestServer()

	resp, err := s.RegisterV2(context.Background(), &pb.RegisterRequest{Service: &pb.RegistryEntry{Name: "test_job", Identifier: "test_server"}})

	if err != nil {
		t.Errorf("Unable to register %v", err)
	}

	// Push this into the master map
	s.masterv2["test_job"] = resp.Service

	respg, err := s.Get(context.Background(), &pb.GetRequest{Job: "test_job"})
	if err != nil {
		t.Errorf("Unable to get %v", err)
	}

	if len(respg.Services) != 1 {
		t.Errorf("Service has been returned")
	}
}

func TestPassMasterElect(t *testing.T) {
	s := InitTestServer()
	te := &testElector{}
	s.elector = te

	_, err := s.RegisterV2(context.Background(), &pb.RegisterRequest{Service: &pb.RegistryEntry{Name: "test_job", Identifier: "test_server"}})
	if err != nil {
		t.Errorf("Error registering serveR: %v", err)
	}

	// Force a master elect
	s.Get(context.Background(), &pb.GetRequest{Job: "test_job"})

	if te.lastElect != "test_job" {
		t.Errorf("No Election took place")
	}
}

func TestFailMasterElect(t *testing.T) {
	s := InitTestServer()
	te := &testElector{}
	s.elector = te

	_, err := s.RegisterV2(context.Background(), &pb.RegisterRequest{Service: &pb.RegistryEntry{Name: "test_job", Identifier: "test_server"}})
	if err != nil {
		t.Errorf("Error registering serveR: %v", err)
	}

	// Force a master elect
	s.Get(context.Background(), &pb.GetRequest{Job: "test_jobssss"})

	if te.lastElect != "" {
		t.Errorf("An election took place: %v", te.lastElect)
	}

}
