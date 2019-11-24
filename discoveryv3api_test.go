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

func TestRegisterV2MasterFail(t *testing.T) {
	s := InitTestServer()

	resp, err := s.RegisterV2(context.Background(), &pb.RegisterRequest{Service: &pb.RegistryEntry{Master: true, Name: "test_job", Identifier: "test_server", Version: pb.RegistryEntry_V2}})

	if err != nil || resp.GetService().Master {
		t.Errorf("Unable to register %v,%v", resp, err)
	}
}

func TestRegisterV2WithAcquireFail(t *testing.T) {
	s := InitTestServer()

	resp, err := s.RegisterV2(context.Background(), &pb.RegisterRequest{Service: &pb.RegistryEntry{Name: "test_job", Identifier: "test_server", Version: pb.RegistryEntry_V2}})

	if err != nil {
		t.Errorf("Unable to register %v", err)
	}

	if resp.Service.Port == 0 {
		t.Errorf("Port number not assigned")
	}

	s.failAcquire = true

	resp2, err := s.MasterElect(context.Background(), &pb.MasterRequest{Service: &pb.RegistryEntry{Name: "test_job", Identifier: "test_server", Version: pb.RegistryEntry_V2}, MasterElect: true})
	if err == nil {
		t.Errorf("Register with lock fail succeeded: %v", resp2)
	}
}

func TestMasterv3(t *testing.T) {
	s := InitTestServer()

	resp, err := s.RegisterV2(context.Background(), &pb.RegisterRequest{Service: &pb.RegistryEntry{Name: "test_job", Identifier: "test_server", Version: pb.RegistryEntry_V2}})

	if err != nil {
		t.Errorf("Unable to register %v", err)
	}

	if resp.Service.Port == 0 {
		t.Errorf("Port number not assigned")
	}

	_, err = s.MasterElect(context.Background(), &pb.MasterRequest{Service: &pb.RegistryEntry{Name: "test_job", Identifier: "test_server"}, MasterElect: true})

	if err != nil {
		t.Errorf("Reg failed: %v", err)
	}
}

func TestMasterv3Fanout(t *testing.T) {
	s := InitTestServer()

	_, err := s.RegisterV2(context.Background(), &pb.RegisterRequest{Service: &pb.RegistryEntry{Name: "test_job", Identifier: "test_server", Version: pb.RegistryEntry_V2, Master: true}, Fanout: true})
	if err != nil {
		t.Errorf("Unable to fanout register %v", err)
	}
}

func TestMasterv3LockFail(t *testing.T) {
	s := InitTestServer()

	resp, err := s.RegisterV2(context.Background(), &pb.RegisterRequest{Service: &pb.RegistryEntry{Name: "test_job", Identifier: "test_server", Version: pb.RegistryEntry_V2}})

	if err != nil {
		t.Errorf("Unable to register %v", err)
	}

	if resp.Service.Port == 0 {
		t.Errorf("Port number not assigned")
	}

	_, err = s.MasterElect(context.Background(), &pb.MasterRequest{Service: &pb.RegistryEntry{Name: "test_job", Identifier: "test_server"}, MasterElect: true, Fanout: true, LockKey: int64(5)})

	if err == nil {
		t.Errorf("Should have failed")
	}

	s.locks["test_job"] = int64(5)
	_, err = s.MasterElect(context.Background(), &pb.MasterRequest{Service: &pb.RegistryEntry{Name: "test_job", Identifier: "test_server"}, MasterElect: true, Fanout: true, LockKey: int64(5)})

	if err != nil {
		t.Errorf("Should not have failed: %v", err)
	}

}
