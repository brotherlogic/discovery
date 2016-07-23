package main

import (
	"errors"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"

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
	return s
}

// ListAllServices returns a list of all the services
func (s *Server) ListAllServices(ctx context.Context, in *pb.Empty) (*pb.ServiceList, error) {
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
					log.Printf("TAKEN")
					taken = true
				}
				// If we've already registered this service, return immediately
				if service.Identifier == in.Identifier && service.Name == in.Name {
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
			taken := false
			for _, service := range s.entries {
				if service.Port == portNumber {
					taken = true
				}

				// If we've already registered this service, return immediately
				if service.Identifier == in.Identifier && service.Name == in.Name {
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
		if entry.Name == in.Name {
			return entry, nil
		}
	}

	return &pb.RegistryEntry{}, errors.New("Cannot find service called " + in.Name)
}

// Serve main server function
func Serve() {
	lis, _ := net.Listen("tcp", port)
	s := grpc.NewServer()
	server := InitServer()
	server.loadCheckFile("checkfile")
	pb.RegisterDiscoveryServiceServer(s, &server)
	err := s.Serve(lis)
	log.Printf("Failed to serve: %v", err)
}
