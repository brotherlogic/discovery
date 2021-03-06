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
		t.Errorf("Unable to get %v", err)
	}

	if len(respg.Services) != 1 {
		t.Errorf("Service not returned")
	}
}

func TestMasterElect(t *testing.T) {
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

	_, err = s.MasterElect(context.Background(), &pb.MasterRequest{Service: &pb.RegistryEntry{Name: "test_job", Identifier: "test_server"}, MasterElect: true})
	if err != nil {
		t.Errorf("Unable to register")
	}

	respg, err = s.Get(context.Background(), &pb.GetRequest{Job: "test_job", Server: "test_server"})
	if err != nil {
		t.Errorf("Unable to get %v", err)
	}

	if len(respg.Services) != 1 || !respg.Services[0].Master {
		t.Errorf("Service not returned or master is wrong: %v", respg.Services)
	}

	resp, err = s.RegisterV2(context.Background(), &pb.RegisterRequest{Service: &pb.RegistryEntry{Name: "test_job", Identifier: "test_server"}})
	if err != nil {
		t.Errorf("Error on reg")
	}

	if resp.GetService().Master {
		t.Fatalf("Rereg held master")
	}

	resp, err = s.RegisterV2(context.Background(), &pb.RegisterRequest{Service: &pb.RegistryEntry{Name: "test_job", Identifier: "test_server2"}})
	if err != nil {
		t.Errorf("Error on reg")
	}

	if resp.GetService().Master {
		t.Fatalf("Rereg held master")
	}

	resp2, err := s.MasterElect(context.Background(), &pb.MasterRequest{Service: &pb.RegistryEntry{Name: "test_job", Identifier: "test_server2"}, MasterElect: true})
	if err != nil {
		t.Errorf("Quick master reg did not fail: %v", resp2)
	}

	resp2, err = s.MasterElect(context.Background(), &pb.MasterRequest{Service: &pb.RegistryEntry{Name: "test_job", Identifier: "test_server"}, MasterElect: true})
	if err == nil {
		t.Errorf("Quick master reg did not fail: %v", resp2)
	}

}

func TestRedirectV2(t *testing.T) {
	s := InitTestServer()

	resp, err := s.RegisterV2(context.Background(), &pb.RegisterRequest{Service: &pb.RegistryEntry{Name: "test_job", Identifier: "test_server"}})

	if err != nil {
		t.Errorf("Unable to register %v", err)
	}

	if resp.Service.Port == 0 {
		t.Errorf("Port number not assigned")
	}

	respg, err := s.Discover(context.Background(), &pb.DiscoverRequest{Caller: "test", Request: &pb.RegistryEntry{Name: "test_job", Identifier: "test_server"}})
	if err != nil {
		t.Errorf("Unable to get %v", err)
	}

	if respg.Service.Name != "test_job" {
		t.Errorf("Service not returned")
	}
}

func TestRedirectV2Master(t *testing.T) {
	s := InitTestServer()

	resp, err := s.RegisterV2(context.Background(), &pb.RegisterRequest{Service: &pb.RegistryEntry{Name: "test_job", Identifier: "test_server"}})

	if err != nil {
		t.Errorf("Unable to register %v", err)
	}

	if resp.Service.Port == 0 {
		t.Errorf("Port number not assigned")
	}

	_, err = s.MasterElect(context.Background(), &pb.MasterRequest{Service: resp.GetService()})
	if err != nil {
		t.Fatalf("Error becoming master: %v", err)
	}

	respg, err := s.Discover(context.Background(), &pb.DiscoverRequest{Caller: "test", Request: &pb.RegistryEntry{Name: "test_job", Identifier: "test_server"}})
	if err != nil {
		t.Errorf("Unable to get %v", err)
	}

	if respg.Service.Name != "test_job" {
		t.Errorf("Service not returned")
	}
}

func TestRedirectV2Fail(t *testing.T) {
	s := InitTestServer()

	resp, err := s.RegisterV2(context.Background(), &pb.RegisterRequest{Service: &pb.RegistryEntry{Name: "test_job", Identifier: "test_server"}})

	if err != nil {
		t.Errorf("Unable to register %v", err)
	}

	if resp.Service.Port == 0 {
		t.Errorf("Port number not assigned")
	}
	_, err = s.Unregister(context.Background(), &pb.UnregisterRequest{Service: &pb.RegistryEntry{Name: "test_job", Identifier: "test_server"}})
	if err != nil {
		t.Errorf("Unable to unregsiter: %v", err)
	}

	_, err = s.Discover(context.Background(), &pb.DiscoverRequest{Caller: "test", Request: &pb.RegistryEntry{Name: "test_job", Identifier: "test_server"}})
	if err == nil {
		t.Errorf("Get did not fail")
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

	_, err = s.Unregister(context.Background(), &pb.UnregisterRequest{Service: &pb.RegistryEntry{Name: "test_job", Identifier: "test_server"}})
	if err != nil {
		t.Fatalf("Error unregistering")
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

	_, err = s.Unregister(context.Background(), &pb.UnregisterRequest{Service: &pb.RegistryEntry{Identifier: "test_server"}})
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

	_, err = s.Unregister(context.Background(), &pb.UnregisterRequest{Service: &pb.RegistryEntry{Name: "test_job2", Identifier: "test_server"}})
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

func TestMasterElectNoRegister(t *testing.T) {
	s := InitTestServer()
	_, err := s.MasterElect(context.Background(), &pb.MasterRequest{Service: &pb.RegistryEntry{Name: "blah"}})
	if err == nil {
		t.Errorf("Should have failed")
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
