package discovery

import (
	"errors"
	"log"
	"net/http"
	"strings"
	"testing"
	"time"

	"golang.org/x/net/context"

	pb "github.com/brotherlogic/discovery/proto"
)

func TestGetExternalIP(t *testing.T) {
	s := InitTestServer()
	externalIP := s.getExternalIP(prodHTTPGetter{})
	if strings.HasSuffix(externalIP, "10.0.1") {
		t.Errorf("External IP is not external enough: %v", externalIP)
	}
}

type testPassChecker struct{}

func (healthChecker testPassChecker) Check(val int, entry *pb.RegistryEntry) int {
	return 0
}

type testFailChecker struct{}

func (healthChecker testFailChecker) Check(val int, entry *pb.RegistryEntry) int {
	return 100
}

type testFailGetter struct{}

func (httpGetter testFailGetter) Get(url string) (*http.Response, error) {
	return nil, errors.New("Built To Fail")
}
func TestGetExternalIPFail(t *testing.T) {
	s := InitTestServer()
	externalIP := s.getExternalIP(testFailGetter{})
	if externalIP != "" {
		t.Errorf("External IP is not blank: %v", externalIP)
	}
}

func TestStartAsSlave(t *testing.T) {
	s := InitTestServer()
	entry1 := &pb.RegistryEntry{Ip: "10.0.1.17", Identifier: "server1", Name: "Job1"}
	s.RegisterService(context.Background(), entry1)

	// Shouldn't be able to find
	entry, err := s.Discover(context.Background(), &pb.RegistryEntry{Name: "Job1"})
	if err == nil {
		t.Fatalf("Haven't failed to discover: %v", entry)
	}

	entry1.Master = true
	s.RegisterService(context.Background(), entry1)

	entry, err = s.Discover(context.Background(), &pb.RegistryEntry{Name: "Job1"})
	if err != nil {
		t.Fatalf("Failed to discover: %v", err)
	}

	if entry.Identifier != "server1" {
		t.Errorf("Weird Error %v", entry)
	}
}

func TestReturnMaster(t *testing.T) {
	s := InitTestServer()
	entry1 := &pb.RegistryEntry{Ip: "10.0.1.17", Identifier: "server1", Name: "Job1", Master: true}
	entry2 := &pb.RegistryEntry{Ip: "10.0.1.18", Identifier: "server2", Name: "Job1", Master: true}

	s.RegisterService(context.Background(), entry1)
	s.RegisterService(context.Background(), entry2)

	entry, err := s.Discover(context.Background(), &pb.RegistryEntry{Name: "Job1"})
	if err != nil {
		t.Fatalf("Error on discover: %v", err)
	}

	if entry.Identifier != "server2" && entry.Identifier != "server1" {
		t.Errorf("Failed to return actual master: %v", entry)
	}
}

func TestRegisterMasterOverride(t *testing.T) {
	s := InitTestServer()
	entry1 := &pb.RegistryEntry{Ip: "10.0.1.17", Identifier: "server1", Name: "Job1", Master: true}
	entry2 := &pb.RegistryEntry{Ip: "10.0.1.18", Identifier: "server2", Name: "Job1", Master: true}

	s.RegisterService(context.Background(), entry1)
	s.RegisterService(context.Background(), entry2)

	r, err := s.ListAllServices(context.Background(), &pb.Empty{})
	if err != nil {
		t.Errorf("Error receiving service list: %v", err)
	}

	if len(r.Services) != 2 {
		t.Fatalf("Wrong number of services received: %v", len(r.Services))
	}

	index1 := 0
	if r.Services[0].Ip != "10.0.1.17" {
		index1 = 2
	}

	if !r.Services[index1].Master {
		t.Errorf("Entry1 is no longers master: %v", r.Services)
	}
}

func TestGetAll(t *testing.T) {
	s := InitTestServer()
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
	s := InitTestServer()
	s.setExternalIP(prodHTTPGetter{})
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
	s := InitTestServer()
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
	s := InitTestServer()
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
	s := InitTestServer()
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

	log.Printf("Trying this one: %v", r2)

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
	s := InitTestServer()
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
	s := InitTestServer()
	entry := &pb.RegistryEntry{Ip: "10.0.4.5", Name: "Testing"}
	r, err := s.RegisterService(context.Background(), entry)
	if err != nil {
		t.Errorf("Error registering service: %v", err)
	}

	if r.Name != entry.Name {
		t.Errorf("Problem with name resolution %v vs %v", r.Name, entry.Name)
	}

	if r.RegisterTime == 0 {
		t.Errorf("Bad time on register: %v", r)
	}

	log.Printf("REGISTERED: %v", r.RegisterTime)
}

func TestCleanWithNoEntries(t *testing.T) {
	s := InitTestServer()
	s.hc = testFailChecker{}

	entry1 := &pb.RegistryEntry{Ip: "10.0.4.5", Name: "Testing1"}
	entry2 := &pb.RegistryEntry{Ip: "10.0.4.5", Name: "Testing2"}

	s.RegisterService(context.Background(), entry1)
	s.RegisterService(context.Background(), entry2)

	s.cleanEntries(time.Now().Add(time.Second * 4))

	r2, err := s.ListAllServices(context.Background(), &pb.Empty{})
	if err != nil {
		t.Fatalf("Failed to list: %v", err)
	}

	if len(r2.Services) > 0 {
		t.Errorf("Services were not actually removed: %v", r2)
	}
}

func TestRegisterWithReregisterService(t *testing.T) {
	s := InitTestServer()
	s.hc = testFailChecker{}
	entry := &pb.RegistryEntry{Ip: "10.0.4.5", Name: "Testing"}
	r, err := s.RegisterService(context.Background(), entry)
	if err != nil {
		t.Errorf("Error registering service: %v", err)
	}

	if r.Name != entry.Name {
		t.Errorf("Problem with name resolution %v vs %v", r.Name, entry.Name)
	}

	if r.Port <= 0 {
		t.Errorf("Register has not received a port number: %v", r)
	}

	s.cleanEntries(time.Now())

	r2, err := s.RegisterService(context.Background(), r)

	if err != nil {
		t.Fatalf("Failed to register second time: %v", err)
	}

	if r2.Name != entry.Name {
		t.Errorf("Re-registry failed")
	}

	if r2.Port != r.Port {
		t.Errorf("Got a different port the second time: %v vs %v", r, r2)
	}
}

func TestSearchWithIdentifier(t *testing.T) {
	s := InitTestServer()
	_, err := s.RegisterService(context.Background(), &pb.RegistryEntry{Ip: "10.0.4.5", Port: 50051, Name: "Testing", Identifier: "serverone"})
	if err != nil {
		t.Errorf("Error registering service: %v", err)
	}
	_, err = s.RegisterService(context.Background(), &pb.RegistryEntry{Ip: "10.0.4.6", Port: 50051, Name: "Testing", Identifier: "servertwo"})
	if err != nil {
		t.Errorf("Error registering service: %v", err)
	}
	_, err = s.RegisterService(context.Background(), &pb.RegistryEntry{Ip: "10.0.4.7", Port: 50051, Name: "Testing", Identifier: "serverthree"})
	if err != nil {
		t.Errorf("Error registering service: %v", err)
	}

	entry := &pb.RegistryEntry{Name: "Testing", Identifier: "servertwo"}
	r, err := s.Discover(context.Background(), entry)
	if err != nil {
		t.Errorf("Cannot discover %v", err)
	}
	if r.Ip != "10.0.4.6" {
		t.Errorf("Wrong server discovered: %v", r)
	}
}

func TestFailedDiscover(t *testing.T) {
	s := InitTestServer()

	entry := &pb.RegistryEntry{Name: "Testing"}
	_, err := s.Discover(context.Background(), entry)
	if err == nil {
		t.Errorf("Disoovering non existing service did not fail: %v", err)
	}
}

func TestDiscover(t *testing.T) {
	s := InitTestServer()
	entryAdd := &pb.RegistryEntry{Ip: "10.0.4.5", Port: 50051, Name: "Testing", Master: true}
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

func InitTestServer() Server {
	s := InitServer()
	s.hc = testPassChecker{}
	return s
}
