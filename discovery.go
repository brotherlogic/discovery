package main

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"

	pb "github.com/brotherlogic/discovery/proto"
)

const (
	port = ":50055"
)

var externalPorts = map[string][]int32{"main": []int32{50052, 50053}}

// Server the central server object
type Server struct {
	entries   []*pb.RegistryEntry
	checkFile string
	hc        healthChecker
}

type healthChecker interface {
	Check(entry *pb.RegistryEntry) bool
}
type httpGetter interface {
	Get(url string) (*http.Response, error)
}
type prodHTTPGetter struct{}

func (httpGetter prodHTTPGetter) Get(url string) (*http.Response, error) {
	return http.Get(url)
}

func (s *Server) getExternalIP(getter httpGetter) string {
	resp, err := getter.Get("http://myexternalip.com/raw")
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return strings.TrimSpace(string(body))
}

func (s *Server) saveCheckFile() {
	serviceList := &pb.ServiceList{Services: s.entries}
	data, _ := proto.Marshal(serviceList)
	ioutil.WriteFile(s.checkFile, data, 0644)
}

func (s *Server) loadCheckFile(fileName string) {
	data, _ := ioutil.ReadFile(fileName)
	serviceList := &pb.ServiceList{}
	proto.Unmarshal(data, serviceList)
	s.entries = serviceList.Services
	s.checkFile = fileName
}

// InitServer builds a server item ready for useo
func InitServer() Server {
	s := Server{}
	s.entries = make([]*pb.RegistryEntry, 0)
	s.hc = prodHealthChecker{}
	return s
}

func (s *Server) cleanEntries() {
	fails := 0
	for i, entry := range s.entries {
		if !s.hc.Check(entry) {
			s.entries = append(s.entries[:(i-fails)], s.entries[(i-fails)+1:]...)
			fails++
		}
	}
}

// ListAllServices returns a list of all the services
func (s *Server) ListAllServices(ctx context.Context, in *pb.Empty) (*pb.ServiceList, error) {
	s.cleanEntries()
	return &pb.ServiceList{Services: s.entries}, nil
}

// RegisterService supports the RegisterService rpc end point
func (s *Server) RegisterService(ctx context.Context, in *pb.RegistryEntry) (*pb.RegistryEntry, error) {

	log.Printf("Registering %v", in)

	// Server is requesting an external port
	if in.ExternalPort {
		availablePorts := externalPorts["main"]
		// Reset the request IP to an external IP
		in.Ip = s.getExternalIP(prodHTTPGetter{})

		for _, port := range availablePorts {
			taken := false
			for _, service := range s.entries {
				if service.Ip == in.Ip && service.Port == port {
					taken = true
				}
				// If we've already registered this service, return immediately
				if service.Identifier == in.Identifier && service.Name == in.Name {
					//Refresh the IP and store the checkfile
					service.Ip = in.Ip
					s.saveCheckFile()
					return service, nil
				}
			}

			if !taken {
				in.Port = port
				break
			}
		}

		//Throw an error if we can't find a port number
		if in.Port <= 0 {
			return in, errors.New("Unable to allocate external port")
		}
	} else {
		var portNumber int32
		for portNumber = 50055 + 1; portNumber < 60000; portNumber++ {
			log.Printf("Assigning port number %v", portNumber)
			taken := false
			for _, service := range s.entries {
				if service.Port == portNumber {
					taken = true
				}

				// If we've already registered this service, return immediately
				if service.Identifier == in.Identifier && service.Name == in.Name {
					//Refresh the IP and store the checkfile
					service.Ip = in.Ip
					s.saveCheckFile()
					return service, nil
				}
			}
			if !taken {
				in.Port = portNumber
				break
			}
		}
	}

	s.entries = append(s.entries, in)
	s.saveCheckFile()
	return in, nil
}

// Discover supports the Discover rpc end point
func (s *Server) Discover(ctx context.Context, in *pb.RegistryEntry) (*pb.RegistryEntry, error) {
	for _, entry := range s.entries {
		if entry.Name == in.Name && (in.Identifier == "" || in.Identifier == entry.Identifier) {
			return entry, nil
		}
	}

	return &pb.RegistryEntry{}, errors.New("Cannot find service called " + in.Name + " on server (maybe): " + in.Identifier)
}
