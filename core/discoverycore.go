package discovery

import (
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"

	pb "github.com/brotherlogic/discovery/proto"

	//Needed to pull in gzip encoding init
	_ "google.golang.org/grpc/encoding/gzip"
)

const (
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
	masterMap map[string]*pb.RegistryEntry
	mm        *sync.Mutex
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
	s.hc = prodHealthChecker{logger: s.recordLog}
	s.m = &sync.Mutex{}
	s.mm = &sync.Mutex{}
	s.masterMap = make(map[string]*pb.RegistryEntry)
	return s
}

func (s *Server) cleanEntries(t time.Time) {
	s.m.Lock()
	fails := 0
	for i, entry := range s.entries {
		//Clean if we haven't seen this entry in the time to clean window
		if t.Sub(time.Unix(entry.GetLastSeenTime(), 0)).Nanoseconds()/1000000 > entry.GetTimeToClean() {
			if entry.GetMaster() {
				s.mm.Lock()
				delete(s.masterMap, entry.GetName())
				s.mm.Unlock()
			}
			s.entries = append(s.entries[:(i-fails)], s.entries[(i-fails)+1:]...)
			fails++
		}
	}
	s.m.Unlock()
}
