package main

import (
	"context"
	"flag"
	"fmt"
	"hash/fnv"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/brotherlogic/goserver"
	"google.golang.org/grpc"

	pb "github.com/brotherlogic/discovery/proto"
	pbg "github.com/brotherlogic/goserver/proto"
)

const (
	port = 50055
)

const (
	strikeCount = 3
)

var externalPorts = map[string][]int32{"main": []int32{50052, 50053}}

// Server the central server object
type Server struct {
	*goserver.GoServer
	entries         []*pb.RegistryEntry
	checkFile       string
	external        string
	lastGet         time.Time
	masterMap       map[string]*pb.RegistryEntry
	mm              *sync.RWMutex
	callerCountM    *sync.Mutex
	callerCount     map[string]int
	reqCountM       *sync.Mutex
	reqCount        map[string]int
	longest         int64
	countRegister   int64
	countDiscover   int64
	countList       int64
	taken           []bool
	extTaken        []bool
	portMap         []*pb.RegistryEntry
	portMemory      map[string]int32
	portMemoryMutex *sync.Mutex
	countV2Register int64
	masterv2Mutex   *sync.Mutex
	masterv2        map[string]*pb.RegistryEntry
	elector         elector
	version         sync.Map
	peerFail        string
	discoverPeer    string
	registerPeer    string
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
func InitServer() *Server {
	s := &Server{}
	s.GoServer = &goserver.GoServer{}
	s.entries = make([]*pb.RegistryEntry, 0)
	s.mm = &sync.RWMutex{}
	s.masterMap = make(map[string]*pb.RegistryEntry)
	s.callerCount = make(map[string]int)
	s.callerCountM = &sync.Mutex{}
	s.reqCount = make(map[string]int)
	s.reqCountM = &sync.Mutex{}
	s.longest = -1
	s.taken = make([]bool, 65536-50052)
	s.extTaken = make([]bool, 2)
	s.portMap = make([]*pb.RegistryEntry, 0)
	s.portMemory = make(map[string]int32)
	s.portMemoryMutex = &sync.Mutex{}
	s.masterv2 = make(map[string]*pb.RegistryEntry)
	s.masterv2Mutex = &sync.Mutex{}
	return s
}

func (s *Server) cleanEntries(t time.Time) {
	nPortMap := []*pb.RegistryEntry{}
	for _, entry := range s.portMap {
		//Clean if we haven't seen this entry in the time to clean window
		if t.UnixNano()-entry.GetLastSeenTime() > entry.GetTimeToClean()*1000000 {
			if entry.GetMaster() {
				s.mm.Lock()
				delete(s.masterMap, entry.GetName())
				s.mm.Unlock()
			}
		} else {
			nPortMap = append(nPortMap, entry)
		}
	}
	s.portMap = nPortMap
}

// DoRegister does RPC registration
func (s *Server) DoRegister(server *grpc.Server) {
	pb.RegisterDiscoveryServiceServer(server, s)
	pb.RegisterDiscoveryServiceV2Server(server, s)
}

func (s *Server) clean(ctx context.Context) error {
	s.cleanEntries(time.Now())
	return nil
}

func main() {
	var quiet = flag.Bool("quiet", false, "Show all output")
	flag.Parse()

	//Turn off logging
	if *quiet {
		log.SetFlags(0)
		log.SetOutput(ioutil.Discard)
	}
	server := InitServer()
	server.setExternalIP(prodHTTPGetter{})
	server.PrepServerNoRegister(port)
	server.Register = server

	server.RegisterServer("discovery", false)

	server.RegisterRepeatingTaskNonMaster(server.clean, "clean", time.Second)

	server.Serve()
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
	SEP = "ww"
)

func (s *Server) hashPortNumber(identifier, job string, sep string) int32 {
	s.portMemoryMutex.Lock()
	defer s.portMemoryMutex.Unlock()
	if val, ok := s.portMemory[job]; ok {
		return val
	}
	//Gets a port number between 50056 and 65535
	portRange := int32(65535 - 50056)
	h := fnv.New32a()
	h.Write([]byte(job))

	portNumber := 50056 + conv(h.Sum32())%portRange
	s.portMemory[job] = portNumber
	return portNumber
}

type elector interface {
	elect(ctx context.Context, entry *pb.RegistryEntry) error
}

type prodElector struct {
}

func (p *prodElector) elect(ctx context.Context, entry *pb.RegistryEntry) error {
	conn, err := grpc.Dial(entry.Ip+":"+strconv.Itoa(int(entry.Port)), grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	server := pbg.NewGoserverServiceClient(conn)
	_, err = server.Mote(ctx, &pbg.MoteRequest{Master: true})

	return err
}

// Mote promotes/demotes this server
func (s *Server) Mote(ctx context.Context, master bool) error {
	return nil
}

func (s *Server) getCurr(in *pb.RegistryEntry) *pb.RegistryEntry {
	for _, val := range s.portMap {
		if val.Identifier == in.Identifier && val.Name == in.Name {
			return val
		}
	}
	return nil
}

func (s *Server) getCMaster(in *pb.RegistryEntry) *pb.RegistryEntry {
	s.mm.RLock()
	defer s.mm.RUnlock()
	return s.masterMap[in.GetName()]
}

func (s *Server) getJob(in *pb.RegistryEntry) (*pb.RegistryEntry, *pb.RegistryEntry) {
	// Setup the port info
	s.setupPort(in)

	return s.getCurr(in), s.getCMaster(in)
}

func (s *Server) addMaster(in *pb.RegistryEntry) {
	//Register as master if there is none
	in.MasterTime = time.Now().UnixNano()
	s.mm.Lock()
	s.masterMap[in.GetName()] = in
	s.mm.Unlock()
}

func (s *Server) removeMaster(in *pb.RegistryEntry) {
	// Remove if we're re-registering without master
	s.mm.Lock()
	delete(s.masterMap, in.GetName())
	s.mm.Unlock()

}

func (s *Server) addToPortMap(in *pb.RegistryEntry) {
	s.portMap = append(s.portMap, in)
}

func (s *Server) removeFromPortMap(in *pb.RegistryEntry) {
	newPortMap := make([]*pb.RegistryEntry, 0)

	for _, entry := range s.portMap {
		if entry.Identifier != in.Identifier || (len(in.Name) > 0 && in.Name != entry.Name) {
			newPortMap = append(newPortMap, entry)
		}
	}

	s.portMap = newPortMap
}

func (s *Server) setupPort(in *pb.RegistryEntry) {
	if in.Port == 0 {
		in.RegisterTime = time.Now().UnixNano()
		if in.ExternalPort {
			in.Ip = s.getExternalIP(prodHTTPGetter{})
		}

		s.setPortNumber(in)
	}
}

func (s *Server) setPortNumber(in *pb.RegistryEntry) error {
	if in.Port == 0 {
		if in.ExternalPort {
			in.Port = 50053
		} else {
			in.Port = s.hashPortNumber(in.Identifier, in.Name, SEP)
		}
	}

	return nil
}

// GetState gets the state of the server
func (s *Server) GetState() []*pbg.State {
	topC := 0
	topCaller := ""
	topR := 0
	topRequest := ""

	s.callerCountM.Lock()
	s.reqCountM.Lock()
	defer s.callerCountM.Unlock()
	defer s.reqCountM.Unlock()
	for key, value := range s.callerCount {
		if value > topC {
			topC = value
			topCaller = key
		}
	}

	for key, value := range s.reqCount {
		if value > topR {
			topR = value
			topRequest = key
		}
	}

	s.mm.RLock()
	defer s.mm.RUnlock()
	return []*pbg.State{
		&pbg.State{Key: "master_map", Text: fmt.Sprintf("%v", s.masterMap)},
		&pbg.State{Key: "discover_peer", Text: s.discoverPeer},
		&pbg.State{Key: "register_peer", Text: s.registerPeer},
		&pbg.State{Key: "top_caller", Text: topCaller},
		&pbg.State{Key: "peer_fail", Text: s.peerFail},
		&pbg.State{Key: "top_requests", Text: fmt.Sprintf("%v (%v)", topRequest, topR)},
		&pbg.State{Key: "version", Text: fmt.Sprintf("%v", s.version)},
	}
}

// ReportHealth alerts if we're not healthy
func (s *Server) ReportHealth() bool {
	return true
}

// Shutdown the server
func (s *Server) Shutdown(ctx context.Context) error {
	return nil
}
