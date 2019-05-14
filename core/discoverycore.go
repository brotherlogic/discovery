package discovery

import (
	"hash/fnv"
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
	entries         []*pb.RegistryEntry
	checkFile       string
	external        string
	lastGet         time.Time
	masterMap       map[string]*pb.RegistryEntry
	mm              *sync.RWMutex
	countM          *sync.Mutex
	counts          map[string]int
	longest         int64
	countRegister   int64
	countDiscover   int64
	countList       int64
	taken           []bool
	extTaken        []bool
	portMap         []*pb.RegistryEntry
	portMemory      map[string]int32
	portMemoryMutex *sync.Mutex
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
	s.mm = &sync.RWMutex{}
	s.masterMap = make(map[string]*pb.RegistryEntry)
	s.counts = make(map[string]int)
	s.countM = &sync.Mutex{}
	s.longest = -1
	s.taken = make([]bool, 65536-50052)
	s.extTaken = make([]bool, 2)
	s.portMap = make([]*pb.RegistryEntry, 65536-50052)
	s.portMemory = make(map[string]int32)
	s.portMemoryMutex = &sync.Mutex{}
	return s
}

func (s *Server) cleanEntries(t time.Time) {
	for i, entry := range s.portMap {
		if entry != nil {
			//Clean if we haven't seen this entry in the time to clean window
			if t.UnixNano()-entry.GetLastSeenTime() > entry.GetTimeToClean()*1000000 {
				if entry.GetMaster() {
					s.mm.Lock()
					delete(s.masterMap, entry.GetName())
					s.mm.Unlock()
				}
				s.portMap[i] = nil
			}
		}
	}
}

func conv(v1 uint32) int32 {
	v2 := int32(v1)
	if v2 < 0 {
		return -v2
	}
	return v2
}

const (
	//SEP eperates out for hashing port number
	SEP = ".."
)

func (s *Server) hashPortNumber(identifier, job string) int32 {
	s.portMemoryMutex.Lock()
	defer s.portMemoryMutex.Unlock()
	if val, ok := s.portMemory[identifier+SEP+job]; ok {
		return val
	}
	//Gets a port number between 50056 and 65535
	portRange := int32(65535 - 50056)
	h := fnv.New32a()
	h.Write([]byte(identifier + SEP + job))

	portNumber := 50056 + conv(h.Sum32())%portRange
	s.portMemory[identifier+SEP+job] = portNumber
	return portNumber
}
