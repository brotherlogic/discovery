package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/context"

	pb "github.com/brotherlogic/discovery/proto"

	//Needed to pull in gzip encoding init
	_ "google.golang.org/grpc/encoding/gzip"
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

func (s *Server) setExternalIP(getter httpGetter) {
	resp, err := getter.Get("http://myexternalip.com/raw")
	if err == nil {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		s.external = strings.TrimSpace(string(body))
	}
}

// GetExternalIP gets the external IP
func (s *Server) getExternalIP(getter httpGetter) string {
	if s.external == "" || time.Now().Sub(s.lastGet) > time.Hour {
		s.lastGet = time.Now()
		go s.setExternalIP(&prodHTTPGetter{})
	}
	return s.external
}

// InitServer builds a server item ready for useo
func InitServer() Server {
	s := Server{}
	s.entries = make([]*pb.RegistryEntry, 0)
	s.strikes = make(map[*pb.RegistryEntry]int)
	s.hc = prodHealthChecker{logger: s.recordLog}
	s.m = &sync.Mutex{}
	return s
}

func (s *Server) cleanEntries() {
	s.m.Lock()
	fails := 0
	for i, entry := range s.entries {
		s.strikes[entry] = s.hc.Check(s.strikes[entry], entry)
		if s.strikes[entry] > strikeCount {
			s.entries = append(s.entries[:(i-fails)], s.entries[(i-fails)+1:]...)
			fails++
		}
	}
	s.m.Unlock()
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

	s.recordTime("Discover-fail", time.Now().Sub(t))
	return &pb.RegistryEntry{}, errors.New("Cannot find service called " + in.Name + " on server (maybe): " + in.Identifier)
}
