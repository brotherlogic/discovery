package discovery

import (
	"context"
	"testing"

	pb "github.com/brotherlogic/discovery/proto"
)

func TestPlainRegisterRun(t *testing.T) {
	s := InitTestServer()

	resp, err := s.Register(context.Background(), &pb.RegisterRequest{Service: &pb.RegistryEntry{Name: "test_job", Identifier: "test_server"}})

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

	resp, err := s.Register(context.Background(), &pb.RegisterRequest{Service: &pb.RegistryEntry{Name: "test_job", Identifier: "test_server"}})

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

	resp, err := s.Register(context.Background(), &pb.RegisterRequest{Service: &pb.RegistryEntry{Name: "test_job", Identifier: "test_server"}})

	if err != nil {
		t.Errorf("Unable to register %v", err)
	}

	if resp.Service.Port == 0 {
		t.Errorf("Port number not assigned")
	}

	resp, err = s.Register(context.Background(), &pb.RegisterRequest{Service: &pb.RegistryEntry{Name: "test_job", Identifier: "test_server"}})

	if err == nil {
		t.Errorf("No error on re-register %v", resp)
	}

}

func TestGetAllV2(t *testing.T) {
	s := InitTestServer()

	_, err := s.Register(context.Background(), &pb.RegisterRequest{Service: &pb.RegistryEntry{Name: "test_job", Identifier: "test_server"}})

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

func TestGetMasterFail(t *testing.T) {
	s := InitTestServer()

	_, err := s.Register(context.Background(), &pb.RegisterRequest{Service: &pb.RegistryEntry{Name: "test_job", Identifier: "test_server"}})

	if err != nil {
		t.Errorf("Unable to register %v", err)
	}

	respg, err := s.Get(context.Background(), &pb.GetRequest{Job: "test_job"})
	if err == nil {
		t.Fatalf("Got %v", respg)
	}
}

func TestGetMaster(t *testing.T) {
	s := InitTestServer()

	resp, err := s.Register(context.Background(), &pb.RegisterRequest{Service: &pb.RegistryEntry{Name: "test_job", Identifier: "test_server"}})

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
