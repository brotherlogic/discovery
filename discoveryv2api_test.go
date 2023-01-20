package main

import (
	"context"
	"testing"
	"time"

	pb "github.com/brotherlogic/discovery/proto"
)

type testElector struct {
	lastElect string
}

func (p *testElector) elect(ctx context.Context, entry *pb.RegistryEntry) error {
	p.lastElect = entry.Name
	return nil
}

func (p *testElector) unelect(ctx context.Context, entry *pb.RegistryEntry) error {
	return nil
}

func TestPlainRegisterFail(t *testing.T) {
	s := InitTestServer()
	s.friendTime = 0

	resp, err := s.RegisterV2(context.Background(), &pb.RegisterRequest{Service: &pb.RegistryEntry{Name: "test_job", Identifier: "test_server"}})

	if err == nil {
		t.Errorf("Register did not fail: %v", resp)
	}
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
		t.Fatalf("Unable to get %v", err)
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

	_, err = s.RegisterV2(context.Background(), &pb.RegisterRequest{Service: &pb.RegistryEntry{Name: "test_job", Identifier: "test_server"}})

	if err != nil {
		t.Errorf("Error on re-register %v", err)
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

func TestPlainUnregisterRun(t *testing.T) {
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

	_, err = s.Unregister(context.Background(), &pb.UnregisterRequest{Reason: "Testing", Service: &pb.RegistryEntry{Name: "test_job", Identifier: "test_server"}})
	if err != nil {
		t.Fatalf("Error unregistering: %v", err)
	}

	respg, err = s.Get(context.Background(), &pb.GetRequest{Job: "test_job", Server: "test_server"})
	if err == nil {
		t.Fatalf("Get succeeded: %v", respg)
	}

}

func TestPlainUnregisterRunWithOnlyIdentifier(t *testing.T) {
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

	_, err = s.Unregister(context.Background(), &pb.UnregisterRequest{Reason: "Testing", Service: &pb.RegistryEntry{Identifier: "test_server"}})
	if err != nil {
		t.Fatalf("Error unregistering")
	}

	respg, err = s.Get(context.Background(), &pb.GetRequest{Job: "test_job", Server: "test_server"})
	if err == nil {
		t.Errorf("Get succeeded: %v", respg)
	}

}

func TestPlainUnregisterRunWithBadCall(t *testing.T) {
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

	_, err = s.Unregister(context.Background(), &pb.UnregisterRequest{Reason: "Testing", Service: &pb.RegistryEntry{Name: "test_job2", Identifier: "test_server"}})
	if err != nil {
		t.Fatalf("Error unregistering")
	}

	_, err = s.Get(context.Background(), &pb.GetRequest{Job: "test_job", Server: "test_server"})
	if err != nil {
		t.Fatalf("Should have succeeded: %v", err)
	}
}

func TestLockLocks(t *testing.T) {
	s := InitTestServer()

	_, err := s.Lock(context.Background(), &pb.LockRequest{Job: "hello", LockKey: time.Now().UnixNano()})

	if err != nil {
		t.Fatalf("Unable to lock")
	}

	_, err = s.Lock(context.Background(), &pb.LockRequest{Job: "hello", LockKey: time.Now().UnixNano(), Requestor: "blah"})
	if err == nil {
		t.Errorf("Lock did fail")
	}
}

func TestEmptyUnRegisterFail(t *testing.T) {
	s := InitTestServer()
	s.friendTime = 0

	resp, err := s.Unregister(context.Background(), &pb.UnregisterRequest{})

	if err == nil {
		t.Errorf("Register did not fail: %v", resp)
	}
}
