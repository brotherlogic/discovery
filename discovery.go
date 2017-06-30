package main

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"

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
	m         *sync.Mutex
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

// InitServer builds a server item ready for useo
func InitServer() Server {
	s := Server{}
	s.entries = make([]*pb.RegistryEntry, 0)
	s.hc = prodHealthChecker{}
	s.m = &sync.Mutex{}
	return s
}

func (s *Server) cleanEntries() {
	s.m.Lock()
	log.Printf("Cleaning")
	fails := 0
	for i, entry := range s.entries {
		if !s.hc.Check(entry) {
			log.Printf("Unable to find %v", entry)
			log.Printf("Removing (%v) %v from %v with %v -> %v", entry, i, len(s.entries), fails, s.entries)
			log.Printf("WITH %v", s.entries[:(i-fails)])
			log.Printf("AND %v", s.entries[(i-fails)+1:])
			log.Printf("REMOVING %v", s.entries[i-fails])
			s.entries = append(s.entries[:(i-fails)], s.entries[(i-fails)+1:]...)
			fails++
		}
	}
	log.Printf("Cleaned")
	s.m.Unlock()
}

// ListAllServices returns a list of all the services
func (s *Server) ListAllServices(ctx context.Context, in *pb.Empty) (*pb.ServiceList, error) {
	log.Printf("Starting Clean")
	s.cleanEntries()
	log.Printf("Cleaned: %v", s.entries)
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
			log.Printf("Assigning port number for external %v", port)
			taken := false
			for _, service := range s.entries {
				if service.Ip == in.Ip && service.Port == port {
					taken = true
				}
				// If we've already registered this service, return immediately
				if service.Identifier == in.Identifier && service.Name == in.Name {
					//Refresh the IP and store the checkfile
					service.Ip = in.Ip
					log.Printf("Fast return : %v", service)
					log.Printf("Returning quick")
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
					return service, nil
				}

				//Unmaster the service if the incoming also wants to be master
				if service.Name == in.Name && service.Master && in.Master {
					service.Master = false
				}
			}
			if !taken {
				//Only set the port if it's not set
				if in.Port <= 0 {
					in.Port = portNumber
				}
				break
			}
		}
	}

	log.Printf("Added to entries %v with %v", s.entries, in)
	s.entries = append(s.entries, in)
	log.Printf("Saving Checkfile")
	log.Printf("Returning")
	return in, nil
}

// Discover supports the Discover rpc end point
func (s *Server) Discover(ctx context.Context, in *pb.RegistryEntry) (*pb.RegistryEntry, error) {
	log.Printf("DISCOVERING: %v", in)
	var nonmaster *pb.RegistryEntry
	for _, entry := range s.entries {
		if entry.Name == in.Name && (in.Identifier == "" || in.Identifier == entry.Identifier) {
			if entry.Master {
				log.Printf("Returning %v", entry)
				return entry, nil
			}
			nonmaster = entry
		}
	}

	//Return the non master if possible
	if nonmaster != nil {
		return nonmaster, nil
	}

	log.Printf("No such service %v", in)
	return &pb.RegistryEntry{}, errors.New("Cannot find service called " + in.Name + " on server (maybe): " + in.Identifier)
}
