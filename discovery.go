package main

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/context"

	pb "github.com/brotherlogic/discovery/proto"
)

const (
	port        = ":50055"
	strikeCount = 3
)

var externalPorts = map[string][]int32{"main": []int32{50052, 50053}}

// Server the central server object
type Server struct {
	entries   []*pb.RegistryEntry
	checkFile string
	hc        healthChecker
	m         *sync.Mutex
	external  string
	lastGet   time.Time
	strikes   map[*pb.RegistryEntry]int
}

type healthChecker interface {
	Check(count int, entry *pb.RegistryEntry) int
}
type httpGetter interface {
	Get(url string) (*http.Response, error)
}
type prodHTTPGetter struct{}

func (httpGetter prodHTTPGetter) Get(url string) (*http.Response, error) {
	return http.Get(url)
}

func (s *Server) getExternalIP(getter httpGetter) string {
	if s.external == "" || time.Now().Sub(s.lastGet) > time.Hour {
		resp, err := getter.Get("http://myexternalip.com/raw")
		if err != nil {
			return ""
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		s.external = strings.TrimSpace(string(body))
		s.lastGet = time.Now()
	}
	return s.external
}

// InitServer builds a server item ready for useo
func InitServer() Server {
	s := Server{}
	s.entries = make([]*pb.RegistryEntry, 0)
	s.strikes = make(map[*pb.RegistryEntry]int)
	s.hc = prodHealthChecker{}
	s.m = &sync.Mutex{}
	return s
}

func (s *Server) cleanEntries() {
	s.m.Lock()
	fails := 0
	for i, entry := range s.entries {
		s.strikes[entry] = s.hc.Check(s.strikes[entry], entry)
		log.Printf("Cleaning %v -> %v", entry, s.strikes[entry])
		if s.strikes[entry] > strikeCount {
			log.Printf("Removing %v", entry)
			s.entries = append(s.entries[:(i-fails)], s.entries[(i-fails)+1:]...)
			fails++
		}
	}
	s.m.Unlock()
}

// ListAllServices returns a list of all the services
func (s *Server) ListAllServices(ctx context.Context, in *pb.Empty) (*pb.ServiceList, error) {
	s.cleanEntries()
	return &pb.ServiceList{Services: s.entries}, nil
}

// RegisterService supports the RegisterService rpc end point
func (s *Server) RegisterService(ctx context.Context, in *pb.RegistryEntry) (*pb.RegistryEntry, error) {
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
					service.Master = in.Master
					return service, nil
				}
			}

			if !taken {
				in.Port = port
				in.RegisterTime = time.Now().Unix()
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
			taken := false
			for _, service := range s.entries {
				if service.Port == portNumber {
					taken = true
				}

				// If we've already registered this service, return immediately
				if service.Identifier == in.Identifier && service.Name == in.Name {
					//Refresh the IP and store the checkfile
					service.Ip = in.Ip
					service.Master = in.Master
					return service, nil
				}
			}
			if !taken {
				//Only set the port if it's not set
				if in.Port <= 0 {
					in.Port = portNumber
					in.RegisterTime = time.Now().Unix()
				}
				break
			}
		}
	}

	s.entries = append(s.entries, in)
	return in, nil
}

// Discover supports the Discover rpc end point
func (s *Server) Discover(ctx context.Context, in *pb.RegistryEntry) (*pb.RegistryEntry, error) {
	t := time.Now()
	var nonmaster *pb.RegistryEntry
	for _, entry := range s.entries {
		if entry.Name == in.Name && (in.Identifier == "" || in.Identifier == entry.Identifier) {
			if entry.Master || in.Identifier != "" {
				s.recordTime("Discover-foundmaster", time.Now().Sub(t))
				return entry, nil
			}
			nonmaster = entry
		}
	}

	//Return the non master if possible
	if nonmaster != nil {
		s.recordTime("Discover-nonmaster", time.Now().Sub(t))
		return nil, errors.New("Cannot find a master for service called " + in.Name + " on server (maybe): " + in.Identifier)
	}

	log.Printf("No such service %v -> %v", in, s.entries)
	s.recordTime("Discover-fail", time.Now().Sub(t))
	return &pb.RegistryEntry{}, errors.New("Cannot find service called " + in.Name + " on server (maybe): " + in.Identifier)
}
