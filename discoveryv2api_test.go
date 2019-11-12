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

	resp, err = s.RegisterV2(context.Background(), &pb.RegisterRequest{Service: &pb.RegistryEntry{Name: "test_job", Identifier: "test_server", Master: true}})
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

	resp, err = s.RegisterV2(context.Background(), &pb.RegisterRequest{Service: &pb.RegistryEntry{Name: "test_job", Identifier: "test_server2"}})
	if err != nil {
		t.Errorf("Error on reg")
	}

	resp, err = s.RegisterV2(context.Background(), &pb.RegisterRequest{Service: &pb.RegistryEntry{Name: "test_job", Identifier: "test_server2", Master: true}})
	if err == nil {
		t.Errorf("Quick master reg did not fail: %v", resp)
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
