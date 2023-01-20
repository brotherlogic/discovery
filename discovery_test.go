package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"testing"
	"time"

	"golang.org/x/net/context"

	pb "github.com/brotherlogic/discovery/proto"
)

func InitTestServer() *Server {
	s := InitServer()
	s.elector = &testElector{}
	s.friendTime = time.Minute
	s.Registry = &pb.RegistryEntry{}
	s.SkipIssue = true
	s.SkipLog = true
	return s
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
	entry := &pb.RegistryEntry{Ip: "10.0.1.17", Identifier: "Server1", Name: "Job1", TimeToClean: 100}
	_, err := s.RegisterService(context.Background(), &pb.RegisterRequest{Service: entry})

	if err != nil {
		t.Fatalf("Error registering service: %v", err)
	}

	val, err := s.Discover(context.Background(), &pb.DiscoverRequest{Caller: "test", Request: &pb.RegistryEntry{Name: "Job1", TimeToClean: 100}})
	if err == nil {
		t.Errorf("Master discover has succeeded: %v", err)
	}

	val, err = s.Discover(context.Background(), &pb.DiscoverRequest{Caller: "test", Request: &pb.RegistryEntry{Name: "Job1", Identifier: "Server1", TimeToClean: 100}})
	if err != nil {
		t.Fatalf("Server discover has failed: %v", err)
	}

	if val.Service.Ip != "10.0.1.17" {
		t.Errorf("Response has come back wrong: %v", val)
	}

	val, err = s.Discover(context.Background(), &pb.DiscoverRequest{Caller: "test", Request: &pb.RegistryEntry{Name: "Job2", Identifier: "Server2", TimeToClean: 100}})
	if err == nil {
		t.Fatalf("Server discover has passed: %v", val)
	}

}

func TestState(t *testing.T) {
	s := InitTestServer()
	s.State(context.Background(), &pb.StateRequest{})
}

func TestDoubleRegister(t *testing.T) {
	s := InitTestServer()
	entry := &pb.RegistryEntry{Ip: "10.0.1.17", Identifier: "Server1", Name: "Job1", TimeToClean: 100}
	res, err := s.RegisterService(context.Background(), &pb.RegisterRequest{Service: entry})
	if err != nil {
		t.Fatalf("Error registering service: %v", err)
	}

	time.Sleep(time.Second * 2)

	entry2 := &pb.RegistryEntry{Ip: "10.0.1.17", Identifier: "Server1", Name: "Job1", TimeToClean: 100}
	res2, err := s.RegisterService(context.Background(), &pb.RegisterRequest{Service: entry2})
	if err != nil {
		t.Fatalf("Error registering service: %v", err)
	}

	if res.GetService().GetRegisterTime() != res2.GetService().GetRegisterTime() {
		t.Errorf("Two things are the same: %v and %v", res, res2)
	}
}

func TestGetAll(t *testing.T) {
	s := InitTestServer()
	entry1 := &pb.RegistryEntry{Ip: "10.0.1.17", Name: "Blah1", TimeToClean: 100}
	entry2 := &pb.RegistryEntry{Ip: "10.0.1.17", Name: "Blah2", TimeToClean: 100}

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
	entry := &pb.RegistryEntry{Ip: "10.0.1.17", Name: "Testing", ExternalPort: true, TimeToClean: 100}
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
	entry := &pb.RegistryEntry{Ip: "10.0.1.17", Name: "Testing", Identifier: "Magic", TimeToClean: 100}

	_, err := s.RegisterService(context.Background(), &pb.RegisterRequest{Service: entry})
	if err != nil {
		t.Errorf("Error registering service: %v", err)
	}

	entry2 := &pb.RegistryEntry{Ip: "10.0.1.20", Name: "Testing", Identifier: "Magic2", TimeToClean: 100}
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
	entry := &pb.RegistryEntry{Ip: "10.0.1.17", Name: "Testing", Identifier: "Magic", TimeToClean: 100}

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

	entry3 := &pb.RegistryEntry{Ip: "10.0.1.17", Name: "Testing", Identifier: "MagicJohnson", TimeToClean: 100}
	_, err = s.RegisterService(context.Background(), &pb.RegisterRequest{Service: entry3})
	if err != nil {
		t.Errorf("Error registering service: %v", err)
	}

}

func TestRegisterService(t *testing.T) {
	s := InitTestServer()
	entry := &pb.RegistryEntry{Ip: "10.0.4.5", Name: "Testing", TimeToClean: 100}
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

	entry1 := &pb.RegistryEntry{Ip: "10.0.4.5", Name: "Testing1", TimeToClean: 100}
	entry2 := &pb.RegistryEntry{Ip: "10.0.4.5", Name: "Testing2", TimeToClean: 100}

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
	entry := &pb.RegistryEntry{Ip: "10.0.4.5", Name: "Testing", TimeToClean: 1}
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

	s.cleanEntries(time.Now().Add(time.Second))

	r4, err := s.ListAllServices(context.Background(), &pb.ListRequest{})
	if err != nil || len(r4.GetServices().Services) != 0 {
		t.Fatalf("Service has not been cleaned %v or %v", r4, err)
	}

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
	_, err := s.RegisterService(context.Background(), &pb.RegisterRequest{Service: &pb.RegistryEntry{Ip: "10.0.4.5", Name: "Testing", Identifier: "serverone", TimeToClean: 100}})
	if err != nil {
		t.Errorf("Error registering service: %v", err)
	}
	_, err = s.RegisterService(context.Background(), &pb.RegisterRequest{Service: &pb.RegistryEntry{Ip: "10.0.4.6", Name: "Testing", Identifier: "servertwo", TimeToClean: 100}})
	if err != nil {
		t.Errorf("Error registering service: %v", err)
	}
	_, err = s.RegisterService(context.Background(), &pb.RegisterRequest{Service: &pb.RegistryEntry{Ip: "10.0.4.7", Name: "Testing", Identifier: "serverthree", TimeToClean: 100}})
	if err != nil {
		t.Errorf("Error registering service: %v", err)
	}

	entry := &pb.RegistryEntry{Name: "Testing", Identifier: "servertwo"}
	r, err := s.Discover(context.Background(), &pb.DiscoverRequest{Caller: "test", Request: entry})
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
	_, err := s.Discover(context.Background(), &pb.DiscoverRequest{Caller: "test", Request: entry})
	if err == nil {
		t.Errorf("Disoovering non existing service did not fail: %v", err)
	}
}

func TestFailExternalGet(t *testing.T) {
	s := InitTestServer()

	//Register 100 regular servers
	for i := 0; i < 100; i++ {
		_, err := s.RegisterService(context.Background(), &pb.RegisterRequest{Service: &pb.RegistryEntry{Ip: "10.0.1.1", Name: fmt.Sprintf("Service-%v", i), TimeToClean: 100}})
		if err != nil {
			t.Errorf("Bad register: %v", err)
		}
	}

	r, err := s.RegisterService(context.Background(), &pb.RegisterRequest{Service: &pb.RegistryEntry{Ip: "10.0.1.1", Name: "thing", ExternalPort: true, TimeToClean: 100}})

	if err != nil {
		t.Errorf("Bad register: %v", err)
	}

	if r.Service.Port > 50054 {
		t.Errorf("Oh dead: %v", r)
	}
}

func TestBadRegister(t *testing.T) {
	s := InitTestServer()
	_, err := s.RegisterService(context.Background(), &pb.RegisterRequest{Service: &pb.RegistryEntry{Name: "blah", Identifier: "alsoblah"}})
	if err == nil {
		t.Errorf("Register without clean did not fail")
	}
}

func TestHasClash(t *testing.T) {
	if findClash(SEP) {
		t.Errorf("Hash Clash")
	}
}

func findClash(sep string) bool {
	s := InitTestServer()

	services := []string{
		"alerter",
		"backup",
		"beerserver",
		"buildserver",
		"cdprocessor",
		"datacollector",
		"dataviewer",
		"dropboxsync",
		"executor",
		"filecopier",
		"githubcard",
		"gobuildmaster",
		"gobuildslave",
		"keystore",
		"location",
		"monitor",
		"printer",
		"proxy",
		"recordalerting",
		"recordbackup",
		"recordcollection",
		"recordgetter",
		"recordmatcher",
		"recordmover",
		"recordprinter",
		"recordprocess",
		"recordsales",
		"recordsorganiser",
		"recordwants",
		"reminders",
		"solver",
		"tracer",
		"versionserver",
		"wantslist",
	}

	space := make(map[int32]int)

	for _, service := range services {
		space[s.hashPortNumber("server", service, sep)]++
	}

	for _, val := range space {
		if val > 1 {
			return true
		}
	}

	return false
}

func TestFind(t *testing.T) {
	options := ":,./abcdefghijklmnopqrstuvwxyz"

	for count := 1; count < 10; count++ {
		for _, c := range options {
			str := ""
			for v := 0; v < count; v++ {
				str += string(c)
			}
			if !findClash(str) {
				log.Printf("FOUND %v", str)
				break
			}
		}
	}
}
