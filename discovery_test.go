package main

import (
	"errors"
	"net/http"
	"strings"
	"testing"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/brotherlogic/discovery/proto"
)

func TestIgnoreIdentifierInDiscover(t *testing.T) {
	s := InitServer()
	s.RegisterService(context.Background(), &pb.RegistryEntry{Name: "test", Identifier: "Blah"})
	f, err := s.Discover(context.Background(), &pb.RegistryEntry{Name: "test"})
	if err != nil {
		t.Errorf("Failure: %v", err)
	} else {
		if f == nil || f.Identifier != "Blah" {
			t.Errorf("Failure to match: %v", f)
		}
	}
}

func TestGetExternalIP(t *testing.T) {
	s := InitServer()
	externalIP := s.getExternalIP(prodHTTPGetter{})
	if strings.HasSuffix(externalIP, "10.0.1") {
		t.Errorf("External IP is not external enough: %v", externalIP)
	}
}

type testFailGetter struct{}

func (httpGetter testFailGetter) Get(url string) (*http.Response, error) {
	return nil, errors.New("Built To Fail")
}
func TestGetExternalIPFail(t *testing.T) {
	s := InitServer()
	externalIP := s.getExternalIP(testFailGetter{})
	if externalIP != "" {
		t.Errorf("External IP is not blank: %v", externalIP)
	}
}

func TestSaveState(t *testing.T) {
	s := InitServer()
	s.checkFile = "test-check"
	entry1 := &pb.RegistryEntry{Ip: "10.0.1.17", Name: "Blah1"}

	s.RegisterService(context.Background(), entry1)

	s2 := InitServer()
	s2.loadCheckFile("test-check")

	r, err := s2.ListAllServices(context.Background(), &pb.Empty{})
	if err != nil {
		t.Errorf("Error receiving service list: %v", err)
	}

	if len(r.Services) != 1 {
		t.Errorf("Wrong number of services received %v", len(r.Services))
	}
}

func TestGetAll(t *testing.T) {
	s := InitServer()
	entry1 := &pb.RegistryEntry{Ip: "10.0.1.17", Name: "Blah1"}
	entry2 := &pb.RegistryEntry{Ip: "10.0.1.17", Name: "Blah2"}

	s.RegisterService(context.Background(), entry1)
	s.RegisterService(context.Background(), entry2)

	r, err := s.ListAllServices(context.Background(), &pb.Empty{})
	if err != nil {
		t.Errorf("Error receiving service list: %v", err)
	}

	if len(r.Services) != 2 {
		t.Errorf("Wrong number of services received: %v", len(r.Services))
	}
}

func TestRegisterForExternalPort(t *testing.T) {
	s := InitServer()
	entry := &pb.RegistryEntry{Ip: "10.0.1.17", Name: "Testing", ExternalPort: true}
	r, err := s.RegisterService(context.Background(), entry)
	if err != nil {
		t.Errorf("Error registering service: %v", err)
	}

	if r.Port <= 0 {
		t.Errorf("Request for external port failed: %v", r)
	}

	if r.Ip == "10.0.1.17" || r.Ip == "" {
		t.Errorf("Request for external port has not returned an external IP: %v", r.Ip)
	}
}

func TestRefreshIP(t *testing.T) {
	s := InitServer()
	entry := &pb.RegistryEntry{Ip: "10.0.1.17", Name: "Testing", Identifier: "Magic"}

	_, err := s.RegisterService(context.Background(), entry)
	if err != nil {
		t.Errorf("Error registering service: %v", err)
	}

	entry2 := &pb.RegistryEntry{Ip: "10.0.1.20", Name: "Testing", Identifier: "Magic"}
	r2, err := s.RegisterService(context.Background(), entry2)
	if err != nil {
		t.Errorf("Error registering service: %v", err)
	}

	if r2.Ip != "10.0.1.20" {
		t.Errorf("Request has not refreshed IP: %v", r2)
	}
}

func TestRegisterMACAddressRefresh(t *testing.T) {
	s := InitServer()
	entry := &pb.RegistryEntry{Ip: "10.0.1.17", Name: "Testing", Identifier: "Magic"}

	r, err := s.RegisterService(context.Background(), entry)
	if err != nil {
		t.Errorf("Error registering service: %v", err)
	}

	entry2 := &pb.RegistryEntry{Ip: "10.0.1.17", Name: "Testing", Identifier: "Magic"}
	r2, err := s.RegisterService(context.Background(), entry2)
	if err != nil {
		t.Errorf("Error registering service: %v", err)
	}

	if r2.Port != r.Port {
		t.Errorf("Same identifier has led to different ports: %v vs %v", r, r2)
	}

	entry3 := &pb.RegistryEntry{Ip: "10.0.1.17", Name: "Testing", Identifier: "MagicJohnson"}
	r3, err := s.RegisterService(context.Background(), entry3)
	if err != nil {
		t.Errorf("Error registering service: %v", err)
	}

	if r3.Port == r2.Port {
		t.Errorf("Different identified but same port: %v vs %v", r3, r2)
	}

}

func TestRegisterMACAddressRefreshWithExternalPort(t *testing.T) {
	s := InitServer()
	entry := &pb.RegistryEntry{Ip: "10.0.1.17", Name: "Testing", Identifier: "Magic", ExternalPort: true}

	r, err := s.RegisterService(context.Background(), entry)
	if err != nil {
		t.Errorf("Error registering service: %v", err)
	}

	entry2 := &pb.RegistryEntry{Ip: "10.0.1.17", Name: "Testing", Identifier: "Magic", ExternalPort: true}
	r2, err := s.RegisterService(context.Background(), entry2)
	if err != nil {
		t.Errorf("Error registering service: %v", err)
	}

	if r2.Port != r.Port {
		t.Errorf("Same identifier has led to different ports: %v vs %v", r, r2)
	}

	entry3 := &pb.RegistryEntry{Ip: "10.0.1.17", Name: "Testing", Identifier: "MagicJohnson", ExternalPort: true}
	r3, err := s.RegisterService(context.Background(), entry3)
	if err != nil {
		t.Errorf("Error registering service: %v", err)
	}

	if r3.Port == r2.Port {
		t.Errorf("Different identified but same port: %v vs %v", r3, r2)
	}
}

func TestRegisterForExternalPortTooManyTimes(t *testing.T) {
	s := InitServer()
	var entries = [...]*pb.RegistryEntry{
		&pb.RegistryEntry{Ip: "10.0.1.17", Name: "Testing1", ExternalPort: true},
		&pb.RegistryEntry{Ip: "10.0.1.17", Name: "Testing2", ExternalPort: true},
		&pb.RegistryEntry{Ip: "10.0.1.17", Name: "Testing3", ExternalPort: true},
	}

	var err error
	for _, entry := range entries {
		_, err = s.RegisterService(context.Background(), entry)
	}

	if err == nil {
		t.Errorf("Over registering does not lead to failure")
	}
}

func TestRegisterService(t *testing.T) {
	s := InitServer()
	entry := &pb.RegistryEntry{Ip: "10.0.4.5", Port: 50051, Name: "Testing"}
	r, err := s.RegisterService(context.Background(), entry)
	if err != nil {
		t.Errorf("Error registering service: %v", err)
	}

	if r.Name != entry.Name {
		t.Errorf("Problem with name resolution %v vs %v", r.Name, entry.Name)
	}
}

func TestFailedDiscover(t *testing.T) {
	s := InitServer()

	entry := &pb.RegistryEntry{Name: "Testing"}
	_, err := s.Discover(context.Background(), entry)
	if err == nil {
		t.Errorf("Disoovering non existing service did not fail: %v", err)
	}
}

func TestDiscover(t *testing.T) {
	s := InitServer()
	entryAdd := &pb.RegistryEntry{Ip: "10.0.4.5", Port: 50051, Name: "Testing"}
	s.RegisterService(context.Background(), entryAdd)
	entry := &pb.RegistryEntry{Name: "Testing"}
	r, err := s.Discover(context.Background(), entry)
	if err != nil {
		t.Errorf("Error registering service: %v", err)
	}

	if r.Ip != entryAdd.Ip {
		t.Errorf("Discovery process failed %v vs %v", r.Ip, entryAdd.Ip)
	}
}

func TestRunServer(t *testing.T) {
	go func() {
		Serve()
	}()

	go func() {
		conn, err := grpc.Dial("localhost:50055", grpc.WithInsecure())
		if err != nil {
			t.Errorf("Error connecting to port")
		}

		defer conn.Close()
		client := pb.NewDiscoveryServiceClient(conn)

		entry := pb.RegistryEntry{}

		_, err = client.RegisterService(context.Background(), &entry)
		if err != nil {
			t.Errorf("Error registering service: %v", err)
		}

		_, err = client.Discover(context.Background(), &entry)
		if err != nil {
			t.Errorf("Error performing discovery: %v", err)
		}

	}()

	time.Sleep(10 * time.Second)
}
