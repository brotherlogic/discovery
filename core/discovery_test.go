package discovery

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"testing"
	"time"

	"golang.org/x/net/context"

	pb "github.com/brotherlogic/discovery/proto"
)

func TestGetState(t *testing.T) {
	s := InitTestServer()
	s.RegisterService(context.Background(), &pb.RegisterRequest{Service: &pb.RegistryEntry{Name: "Blah"}})
	state, err := s.State(context.Background(), &pb.StateRequest{})
	if err != nil {
		t.Fatalf("Error getting state: %v", err)
	}
	if state.Frequency != 1 {
		t.Errorf("HMmm: %v", state.Frequency)
	}
}

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

func TestServerDiscover(t *testing.T) {
	s := InitTestServer()
	entry := &pb.RegistryEntry{Ip: "10.0.1.17", Identifier: "Server1", Name: "Job1"}
	_, err := s.RegisterService(context.Background(), &pb.RegisterRequest{Service: entry})

	if err != nil {
		t.Fatalf("Error registering service: %v", err)
	}

	val, err := s.Discover(context.Background(), &pb.DiscoverRequest{Request: &pb.RegistryEntry{Name: "Job1"}})
	if err == nil {
		t.Errorf("Master discover has succeeded: %v", err)
	}

	val, err = s.Discover(context.Background(), &pb.DiscoverRequest{Request: &pb.RegistryEntry{Name: "Job1", Identifier: "Server1"}})
	if err != nil {
		t.Errorf("Server discover has failed: %v", err)
	}

	if val.Service.Ip != "10.0.1.17" {
		t.Errorf("Response has come back wrong: %v", val)
	}
}

func TestRegisterPort(t *testing.T) {
	s := InitTestServer()
	entry := &pb.RegistryEntry{Ip: "10.0.1.17", Identifier: "Server1", Name: "Job1", Port: 12345}
	_, err := s.RegisterService(context.Background(), &pb.RegisterRequest{Service: entry})

	if err != nil {
		t.Fatalf("Error registering service: %v", err)
	}

	val, err := s.Discover(context.Background(), &pb.DiscoverRequest{Request: &pb.RegistryEntry{Name: "Job1", Identifier: "Server1"}})
	if err != nil {
		t.Errorf("Server discover has failed: %v", err)
	}

	if val.Service.Ip != "10.0.1.17" || val.Service.Port != 12345 {
		t.Errorf("Response has come back wrong: %v", val)
	}

}

func TestDoubleRegister(t *testing.T) {
	s := InitTestServer()
	entry := &pb.RegistryEntry{Ip: "10.0.1.17", Identifier: "Server1", Name: "Job1"}
	res, err := s.RegisterService(context.Background(), &pb.RegisterRequest{Service: entry})
	if err != nil {
		t.Fatalf("Error registering service: %v", err)
	}

	time.Sleep(time.Second * 2)

	entry2 := &pb.RegistryEntry{Ip: "10.0.1.17", Identifier: "Server1", Name: "Job1"}
	res2, err := s.RegisterService(context.Background(), &pb.RegisterRequest{Service: entry2})
	if err != nil {
		t.Fatalf("Error registering service: %v", err)
	}

	if res.GetService().GetRegisterTime() == res2.GetService().GetRegisterTime() {
		t.Errorf("Two things are the same: %v and %v", res, res2)
	}
}

func TestFailAsMasterRegister(t *testing.T) {
	s := InitTestServer()
	entry := &pb.RegistryEntry{Ip: "10.0.1.17", Identifier: "server1", Name: "Job1", Master: true}
	res, err := s.RegisterService(context.Background(), &pb.RegisterRequest{Service: entry})

	if err == nil {
		t.Errorf("Master register has not failed: %v", res)
	}
}

func TestStartAsSlave(t *testing.T) {
	s := InitTestServer()
	entry1 := &pb.RegistryEntry{Ip: "10.0.1.17", Identifier: "server1", Name: "Job1"}
	s.RegisterService(context.Background(), &pb.RegisterRequest{Service: entry1})

	// Shouldn't be able to find
	entry, err := s.Discover(context.Background(), &pb.DiscoverRequest{Request: &pb.RegistryEntry{Name: "Job1"}})
	if err == nil {
		t.Fatalf("Haven't failed to discover: %v", entry)
	}

	entry1.Master = true
	s.RegisterService(context.Background(), &pb.RegisterRequest{Service: entry1})

	entry, err = s.Discover(context.Background(), &pb.DiscoverRequest{Request: &pb.RegistryEntry{Name: "Job1"}})
	if err != nil {
		t.Fatalf("Failed to discover: %v", err)
	}

	if entry.GetService().Identifier != "server1" {
		t.Errorf("Weird Error %v", entry)
	}
}

func TestGetAll(t *testing.T) {
	s := InitTestServer()
	entry1 := &pb.RegistryEntry{Ip: "10.0.1.17", Name: "Blah1"}
	entry2 := &pb.RegistryEntry{Ip: "10.0.1.17", Name: "Blah2"}

	s.RegisterService(context.Background(), &pb.RegisterRequest{Service: entry1})
	s.RegisterService(context.Background(), &pb.RegisterRequest{Service: entry2})

	r, err := s.ListAllServices(context.Background(), &pb.ListRequest{})
	if err != nil {
		t.Errorf("Error receiving service list: %v", err)
	}

	if len(r.GetServices().Services) != 2 {
		t.Errorf("Wrong number of services received: %v", len(r.GetServices().Services))
	}
}

func TestRegisterForExternalPort(t *testing.T) {
	s := InitTestServer()
	s.setExternalIP(prodHTTPGetter{})
	entry := &pb.RegistryEntry{Ip: "10.0.1.17", Name: "Testing", ExternalPort: true}
	r, err := s.RegisterService(context.Background(), &pb.RegisterRequest{Service: entry})
	if err != nil {
		t.Errorf("Error registering service: %v", err)
	}

	if r.GetService().Port <= 0 {
		t.Errorf("Request for external port failed: %v", r)
	}

	if r.GetService().Ip == "10.0.1.17" || r.GetService().Ip == "" {
		t.Errorf("Request for external port has not returned an external IP: %v", r.GetService().Ip)
	}
}

func TestRefreshIP(t *testing.T) {
	s := InitTestServer()
	entry := &pb.RegistryEntry{Ip: "10.0.1.17", Name: "Testing", Identifier: "Magic"}

	_, err := s.RegisterService(context.Background(), &pb.RegisterRequest{Service: entry})
	if err != nil {
		t.Errorf("Error registering service: %v", err)
	}

	entry2 := &pb.RegistryEntry{Ip: "10.0.1.20", Name: "Testing", Identifier: "Magic"}
	r2, err := s.RegisterService(context.Background(), &pb.RegisterRequest{Service: entry2})
	if err != nil {
		t.Errorf("Error registering service: %v", err)
	}

	if r2.GetService().Ip != "10.0.1.20" {
		t.Errorf("Request has not refreshed IP: %v", r2)
	}
}

func TestRegisterMACAddressRefresh(t *testing.T) {
	s := InitTestServer()
	entry := &pb.RegistryEntry{Ip: "10.0.1.17", Name: "Testing", Identifier: "Magic"}

	r, err := s.RegisterService(context.Background(), &pb.RegisterRequest{Service: entry})
	if err != nil {
		t.Errorf("Error registering service: %v", err)
	}

	r2, err := s.RegisterService(context.Background(), &pb.RegisterRequest{Service: r.GetService()})
	if err != nil {
		t.Errorf("Error registering service: %v", err)
	}

	if r2.GetService().Port != r.GetService().Port {
		t.Errorf("Same identifier has led to different ports: %v vs %v", r, r2)
	}

	entry3 := &pb.RegistryEntry{Ip: "10.0.1.17", Name: "Testing", Identifier: "MagicJohnson"}
	r3, err := s.RegisterService(context.Background(), &pb.RegisterRequest{Service: entry3})
	if err != nil {
		t.Errorf("Error registering service: %v", err)
	}

	if r3.GetService().Port == r2.GetService().Port {
		t.Errorf("Different identified but same port: %v vs %v", r3, r2)
	}

}

func TestRegisterMACAddressRefreshWithExternalPort(t *testing.T) {
	s := InitTestServer()
	entry := &pb.RegistryEntry{Ip: "10.0.1.17", Name: "Testing", Identifier: "Magic", ExternalPort: true}

	r, err := s.RegisterService(context.Background(), &pb.RegisterRequest{Service: entry})
	if err != nil {
		t.Errorf("Error registering service: %v", err)
	}

	r2, err := s.RegisterService(context.Background(), &pb.RegisterRequest{Service: r.GetService()})
	if err != nil {
		t.Errorf("Error registering service: %v", err)
	}

	if r2.GetService().Port != r.GetService().Port {
		t.Errorf("Same identifier has led to different ports: %v vs %v", r, r2)
	}

	entry3 := &pb.RegistryEntry{Ip: "10.0.1.17", Name: "Testing", Identifier: "MagicJohnson", ExternalPort: true}
	r3, err := s.RegisterService(context.Background(), &pb.RegisterRequest{Service: entry3})
	if err != nil {
		t.Errorf("Error registering service: %v", err)
	}

	if r3.GetService().Port == r2.GetService().Port {
		t.Errorf("Different identified but same port: %v vs %v", r3, r2)
	}
}

func TestRegisterForExternalPortTooManyTimes(t *testing.T) {
	s := InitTestServer()
	var entries = [...]*pb.RegistryEntry{
		&pb.RegistryEntry{Ip: "10.0.1.17", Name: "Testing1", ExternalPort: true, Port: 1234},
		&pb.RegistryEntry{Ip: "10.0.1.17", Name: "Testing2", ExternalPort: true, Port: 1234},
		&pb.RegistryEntry{Ip: "10.0.1.17", Name: "Testing3", ExternalPort: true, Port: 1234},
	}

	var err error
	for _, entry := range entries {
		_, err = s.RegisterService(context.Background(), &pb.RegisterRequest{Service: entry})
	}

	if err == nil {
		t.Errorf("Over registering does not lead to failure")
	}
}

func TestRegisterService(t *testing.T) {
	s := InitTestServer()
	entry := &pb.RegistryEntry{Ip: "10.0.4.5", Name: "Testing"}
	r, err := s.RegisterService(context.Background(), &pb.RegisterRequest{Service: entry})
	if err != nil {
		t.Errorf("Error registering service: %v", err)
	}

	if r.GetService().Name != entry.Name {
		t.Errorf("Problem with name resolution %v vs %v", r.GetService().Name, entry.Name)
	}

	if r.GetService().RegisterTime == 0 {
		t.Errorf("Bad time on register: %v", r)
	}
}

func TestCleanWithNoEntries(t *testing.T) {
	s := InitTestServer()

	entry1 := &pb.RegistryEntry{Ip: "10.0.4.5", Name: "Testing1"}
	entry2 := &pb.RegistryEntry{Ip: "10.0.4.5", Name: "Testing2"}

	s.RegisterService(context.Background(), &pb.RegisterRequest{Service: entry1})
	s.RegisterService(context.Background(), &pb.RegisterRequest{Service: entry2})

	s.cleanEntries(time.Now().Add(time.Second * 6))

	r2, err := s.ListAllServices(context.Background(), &pb.ListRequest{})
	if err != nil {
		t.Fatalf("Failed to list: %v", err)
	}

	if len(r2.GetServices().Services) > 0 {
		t.Errorf("Services were not actually removed: %v", r2)
	}
}

func TestRegisterWithReregisterService(t *testing.T) {
	s := InitTestServer()
	entry := &pb.RegistryEntry{Ip: "10.0.4.5", Name: "Testing"}
	r, err := s.RegisterService(context.Background(), &pb.RegisterRequest{Service: entry})
	if err != nil {
		t.Errorf("Error registering service: %v", err)
	}

	if r.GetService().Name != entry.Name {
		t.Errorf("Problem with name resolution %v vs %v", r.GetService().Name, entry.Name)
	}

	if r.GetService().Port <= 0 {
		t.Errorf("Register has not received a port number: %v", r)
	}

	s.cleanEntries(time.Now())

	r2, err := s.RegisterService(context.Background(), &pb.RegisterRequest{Service: r.GetService()})

	if err != nil {
		t.Fatalf("Failed to register second time: %v", err)
	}

	if r2.GetService().Name != entry.Name {
		t.Errorf("Re-registry failed")
	}

	if r2.GetService().Port != r.GetService().Port {
		t.Errorf("Got a different port the second time: %v vs %v", r, r2)
	}
}

func TestSearchWithIdentifier(t *testing.T) {
	s := InitTestServer()
	_, err := s.RegisterService(context.Background(), &pb.RegisterRequest{Service: &pb.RegistryEntry{Ip: "10.0.4.5", Port: 50051, Name: "Testing", Identifier: "serverone"}})
	if err != nil {
		t.Errorf("Error registering service: %v", err)
	}
	_, err = s.RegisterService(context.Background(), &pb.RegisterRequest{Service: &pb.RegistryEntry{Ip: "10.0.4.6", Port: 50051, Name: "Testing", Identifier: "servertwo"}})
	if err != nil {
		t.Errorf("Error registering service: %v", err)
	}
	_, err = s.RegisterService(context.Background(), &pb.RegisterRequest{Service: &pb.RegistryEntry{Ip: "10.0.4.7", Port: 50051, Name: "Testing", Identifier: "serverthree"}})
	if err != nil {
		t.Errorf("Error registering service: %v", err)
	}

	entry := &pb.RegistryEntry{Name: "Testing", Identifier: "servertwo"}
	r, err := s.Discover(context.Background(), &pb.DiscoverRequest{Request: entry})
	if err != nil {
		t.Errorf("Cannot discover %v", err)
	}
	if r.GetService().Ip != "10.0.4.6" {
		t.Errorf("Wrong server discovered: %v", r)
	}
}

func TestFailedDiscover(t *testing.T) {
	s := InitTestServer()

	entry := &pb.RegistryEntry{Name: "Testing"}
	_, err := s.Discover(context.Background(), &pb.DiscoverRequest{Request: entry})
	if err == nil {
		t.Errorf("Disoovering non existing service did not fail: %v", err)
	}
}

func InitTestServer() Server {
	s := InitServer()
	return s
}

func TestBecomeMaster(t *testing.T) {
	s := InitTestServer()

	entry1 := &pb.RegistryEntry{Ip: "10.0.4.5", Identifier: "ShouldBeMaster", Name: "Testing"}
	entry2 := &pb.RegistryEntry{Ip: "10.0.4.6", Identifier: "ShouldBeSlave", Name: "Testing"}

	s.RegisterService(context.Background(), &pb.RegisterRequest{Service: entry1})
	s.RegisterService(context.Background(), &pb.RegisterRequest{Service: entry2})

	v, err := s.Discover(context.Background(), &pb.DiscoverRequest{Request: &pb.RegistryEntry{Name: "Testing"}})
	if err == nil {
		t.Fatalf("Successful discover with non master: %v", v)
	}

	entry1.Master = true
	_, err = s.RegisterService(context.Background(), &pb.RegisterRequest{Service: entry1})
	if err != nil {
		t.Fatalf("Unable to re-register as master: %v", err)
	}

	v, err = s.Discover(context.Background(), &pb.DiscoverRequest{Request: &pb.RegistryEntry{Name: "Testing"}})
	if err != nil || v.GetService().GetIp() != "10.0.4.5" {
		t.Fatalf("Master is incorrect: %v", v)
	}

	entry1.Master = false
	entry1Back, err := s.RegisterService(context.Background(), &pb.RegisterRequest{Service: entry1})
	if err != nil {
		t.Fatalf("Unable to re-register as slave: %v", err)
	}

	if entry1Back.GetService().MasterTime > 0 {
		t.Fatalf("Master time not reset: %v", entry1Back)
	}

	v, err = s.Discover(context.Background(), &pb.DiscoverRequest{Request: &pb.RegistryEntry{Name: "Testing"}})
	if err == nil {
		t.Fatalf("Master is being returned: %v", v)
	}

	entry2.Master = true
	entry2Back, err := s.RegisterService(context.Background(), &pb.RegisterRequest{Service: entry2})
	if err != nil {
		t.Errorf("Unable to promote to master: %v", err)
	}

	if entry2Back.GetService().GetMasterTime() == 0 {
		t.Errorf("Master time not set on new master: %v", entry2Back)
	}

}

func TestFailHeartbeat(t *testing.T) {
	s := InitTestServer()

	entry1 := &pb.RegistryEntry{Ip: "10.0.4.5", Identifier: "ShouldBeMaster", Name: "Testing"}
	entry2 := &pb.RegistryEntry{Ip: "10.0.4.6", Identifier: "ShouldBeSlave", Name: "Testing"}

	s.RegisterService(context.Background(), &pb.RegisterRequest{Service: entry1})
	s.RegisterService(context.Background(), &pb.RegisterRequest{Service: entry2})

	v, err := s.Discover(context.Background(), &pb.DiscoverRequest{Request: &pb.RegistryEntry{Name: "Testing"}})
	if err == nil {
		t.Fatalf("Successful discover with non master: %v", v)
	}

	entry1.Master = true
	_, err = s.RegisterService(context.Background(), &pb.RegisterRequest{Service: entry1})
	if err != nil {
		t.Fatalf("Unable to re-register as master: %v", err)
	}

	v, err = s.Discover(context.Background(), &pb.DiscoverRequest{Request: &pb.RegistryEntry{Name: "Testing"}})
	if err != nil || v.GetService().GetIp() != "10.0.4.5" {
		t.Fatalf("Master is incorrect: %v", v)
	}

	entry2.Master = true
	v2, err := s.RegisterService(context.Background(), &pb.RegisterRequest{Service: entry2})
	if err == nil {
		t.Errorf("Succesful promote to master: %v", v2)
	}
}

func TestCleanWithMaster(t *testing.T) {
	s := InitTestServer()
	entry := &pb.RegistryEntry{Ip: "10.0.4.5", Identifier: "ShouldBeMaster", Name: "Testing"}
	_, err := s.RegisterService(context.Background(), &pb.RegisterRequest{Service: entry})
	if err != nil {
		t.Fatalf("Error doing initial reg: %v", err)
	}
	entry.Master = true
	_, err = s.RegisterService(context.Background(), &pb.RegisterRequest{Service: entry})
	if err != nil {
		t.Fatalf("Error registering as master")
	}

	s.cleanEntries(time.Now().Add(time.Hour))

	v, err := s.RegisterService(context.Background(), &pb.RegisterRequest{Service: entry})
	if err == nil {
		t.Errorf("Reregister as master after clean has not failed: %v", v)
	}
}

func TestFailHeartbeatExternal(t *testing.T) {
	s := InitTestServer()

	entry1 := &pb.RegistryEntry{Ip: "10.0.4.5", Identifier: "ShouldBeMaster", Name: "Testing", ExternalPort: true}
	entry2 := &pb.RegistryEntry{Ip: "10.0.4.6", Identifier: "ShouldBeSlave", Name: "Testing", ExternalPort: true}

	s.RegisterService(context.Background(), &pb.RegisterRequest{Service: entry1})
	s.RegisterService(context.Background(), &pb.RegisterRequest{Service: entry2})

	v, err := s.Discover(context.Background(), &pb.DiscoverRequest{Request: &pb.RegistryEntry{Name: "Testing"}})
	if err == nil {
		t.Fatalf("Successful discover with non master: %v", v)
	}

	entry1.Master = true
	_, err = s.RegisterService(context.Background(), &pb.RegisterRequest{Service: entry1})
	if err != nil {
		t.Fatalf("Unable to re-register as master: %v", err)
	}

	v, err = s.Discover(context.Background(), &pb.DiscoverRequest{Request: &pb.RegistryEntry{Name: "Testing"}})
	if err != nil || v.GetService().GetIdentifier() != "ShouldBeMaster" {
		t.Fatalf("Master is incorrect: %v", v)
	}

	entry2.Master = true
	v2, err := s.RegisterService(context.Background(), &pb.RegisterRequest{Service: entry2})
	if err == nil {
		t.Errorf("Succesful promote to master: %v", v2)
	}
}

func TestFailExternalGet(t *testing.T) {
	s := InitTestServer()

	//Register 100 regular servers
	for i := 0; i < 100; i++ {
		_, err := s.RegisterService(context.Background(), &pb.RegisterRequest{Service: &pb.RegistryEntry{Ip: "10.0.1.1", Name: fmt.Sprintf("Service-%v", i)}})
		if err != nil {
			t.Errorf("Bad register: %v", err)
		}
	}

	r, err := s.RegisterService(context.Background(), &pb.RegisterRequest{Service: &pb.RegistryEntry{Ip: "10.0.1.1", Name: "thing", ExternalPort: true}})

	if err != nil {
		t.Errorf("Bad register: %v", err)
	}

	if r.Service.Port > 50053 {
		t.Errorf("Oh dead: %v", r)
	}
}
