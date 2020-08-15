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
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/brotherlogic/discovery/proto"
	pbg "github.com/brotherlogic/goserver/proto"
)

const (
	port = 50055
)

const (
	strikeCount = 3
)

var (
	//Friends discovery chums
	Friends = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "discovery_friends",
		Help: "The number of friends we have",
	})
)

var externalPorts = map[string][]int32{"main": []int32{50052, 50053}}

// Server the central server object
type Server struct {
	*goserver.GoServer
	friends         []string
	entries         []*pb.RegistryEntry
	checkFile       string
	external        string
	lastGet         time.Time
	masterMap       map[string]*pb.RegistryEntry
	masterTime      map[string]time.Time
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
	friendTime      time.Duration
	locks           map[string]int64
	lockNames       map[string]string
	failAcquire     bool
	lastError       string
	lastRemove      string
	getMap          sync.Map
	countMap        map[int]string
	getLoad         int
	getMapB         map[string]int
	mapLock         *sync.Mutex
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
	s.masterTime = make(map[string]time.Time)
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
	s.friends = make([]string, 0)
	s.locks = make(map[string]int64)
	s.lockNames = make(map[string]string)
	s.elector = &prodElector{dial: s.DoDial}
	s.countMap = make(map[int]string)
	s.getMapB = make(map[string]int)
	s.mapLock = &sync.Mutex{}
	return s
}

func (s *Server) cleanEntries(t time.Time) {
	nPortMap := []*pb.RegistryEntry{}
	for _, entry := range s.portMap {
		//Clean if we haven't seen this entry in the time to clean window
		// Don't clean V2 entries
		if t.UnixNano()-entry.GetLastSeenTime() > entry.GetTimeToClean()*1000000 && entry.Version == pb.RegistryEntry_V1 {
			if entry.GetMaster() {
				s.mm.Lock()
				delete(s.masterMap, entry.GetName())
				delete(s.masterTime, entry.GetName())
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
	unelect(ctx context.Context, entry *pb.RegistryEntry) error
}

type prodElector struct {
	dial func(entry *pb.RegistryEntry) (*grpc.ClientConn, error)
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
func (p *prodElector) unelect(ctx context.Context, entry *pb.RegistryEntry) error {
	if entry == nil {
		return nil
	}
	entry.Master = false

	conn, err := p.dial(entry)
	if err != nil {
		return err
	}
	defer conn.Close()

	server := pbg.NewGoserverServiceClient(conn)
	_, err = server.Mote(ctx, &pbg.MoteRequest{Master: false})

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

func (s *Server) getCMaster(in *pb.RegistryEntry) (*pb.RegistryEntry, time.Time) {
	s.mm.RLock()
	defer s.mm.RUnlock()
	return s.masterMap[in.GetName()], s.masterTime[in.GetName()]
}

func (s *Server) getJob(in *pb.RegistryEntry) (*pb.RegistryEntry, *pb.RegistryEntry) {
	// Setup the port info
	s.setupPort(in)

	m, _ := s.getCMaster(in)
	return s.getCurr(in), m
}

func (s *Server) addMaster(in *pb.RegistryEntry) {
	//Register as master if there is none
	in.MasterTime = time.Now().UnixNano()
	s.mm.Lock()
	if val, ok := s.masterMap[in.GetName()]; ok {
		val.Master = false
	}
	s.masterMap[in.GetName()] = in
	s.masterTime[in.GetName()] = time.Now()
	s.mm.Unlock()
}

func (s *Server) removeMaster(in *pb.RegistryEntry) {
	// Remove if we're re-registering without master
	s.mm.Lock()
	delete(s.masterMap, in.GetName())
	delete(s.masterTime, in.GetName())
	s.mm.Unlock()

}

func (s *Server) addToPortMap(in *pb.RegistryEntry) {
	s.portMap = append(s.portMap, in)
}

func (s *Server) removeFromPortMap(in *pb.RegistryEntry) {
	s.lastRemove = fmt.Sprintf("%v", in)

	if in == nil {
		return
	}

	newPortMap := make([]*pb.RegistryEntry, 0)

	for _, entry := range s.portMap {
		if entry.GetIdentifier() != in.GetIdentifier() ||
			(len(in.GetName()) > 0 && in.GetName() != entry.GetName()) {
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
		if in.ExternalPort && in.Name == "proxy" {
			in.Port = 50053
		} else if in.ExternalPort && in.Name == "secureproxy" {
			in.Port = 50040
		} else if in.ExternalPort {
			in.Port = 50054
		} else {
			in.Port = s.hashPortNumber(in.Identifier, in.Name, SEP)
		}
	}

	return nil
}

func (s *Server) findFriend(host int) {
	hostStr := fmt.Sprintf("192.168.86.%v:50055", host)
	if fmt.Sprintf("%v:50055", s.Registry.Ip) == hostStr {
		s.countMap[host] = fmt.Sprintf("%v SELF", time.Now())
		return
	}
	for _, f := range s.friends {
		if f == hostStr {
			return
		}
	}
	conn, err := s.FDial(fmt.Sprintf("192.168.86.%v:50055", host))
	if err == nil {
		defer conn.Close()
		client := pbg.NewGoserverServiceClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		_, err := client.IsAlive(ctx, &pbg.Alive{})
		if err == nil {
			s.countMap[host] = fmt.Sprintf("%v FOUND_FRIEND", time.Now())
			s.friends = append(s.friends, hostStr)
			Friends.Set(float64(len(s.friends)))
			s.readFriend(hostStr)
		} else {
			c := status.Convert(err)
			if c.Code() != codes.DeadlineExceeded {
				s.lastError = fmt.Sprintf("%v", err)
			}
			s.countMap[host] = fmt.Sprintf("%v", err)
		}

	}
}

func (s *Server) validateFriends() {
	for _, f := range s.friends {
		s.readFriend(f)
	}
}

func (s *Server) checkFriend(addr string) {
	//Don't friend ourselves
	if addr == s.Registry.GetIp() {
		return
	}

	newaddr := addr + ":50055"

	for _, f := range s.friends {
		if f == newaddr {
			return
		}
	}

	s.friends = append(s.friends, newaddr)
	Friends.Set(float64(len(s.friends)))
}

var (
	etcreg = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "discovery_etcreg",
		Help: "ETCD Registry Attempts",
	}, []string{"error"})
)

func (s *Server) readFriend(host string) {
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err == nil {
		defer conn.Close()
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		client := pb.NewDiscoveryServiceV2Client(conn)
		regs, err := client.Get(ctx, &pb.GetRequest{Friend: fmt.Sprintf("%v:%v", s.Registry.Ip, s.Registry.Port)})
		if err == nil {
			for _, entry := range regs.GetServices() {
				if entry.GetVersion() != pb.RegistryEntry_V1 {
					s.RegisterV2(ctx, &pb.RegisterRequest{Fanout: true, Service: entry})
				}
			}

		} else {
			s.lastError = fmt.Sprintf("%v", err)
		}
	}
}

// GetState gets the state of the server
func (s *Server) GetState() []*pbg.State {
	topCaller := ""
	topR := 0
	topRequest := ""

	bad := int64(0)
	for _, blah := range s.portMap {
		if blah == nil {
			bad++
		}
	}

	s.mm.RLock()
	defer s.mm.RUnlock()
	return []*pbg.State{
		&pbg.State{Key: "friend", Text: fmt.Sprintf("%v", s.friends)},
		&pbg.State{Key: "friend_map", Text: fmt.Sprintf("%v", s.countMap)},
		&pbg.State{Key: "last_remove", Text: s.lastRemove},
		&pbg.State{Key: "ports", Value: bad},
		&pbg.State{Key: "locks", Text: fmt.Sprintf("%v", s.locks)},
		&pbg.State{Key: "locks_name", Text: fmt.Sprintf("%v", s.lockNames)},
		&pbg.State{Key: "last_error", Text: s.lastError},
		&pbg.State{Key: "ftime", TimeDuration: s.friendTime.Nanoseconds()},
		&pbg.State{Key: "master_map", Text: fmt.Sprintf("%v", s.masterMap)},
		&pbg.State{Key: "discover_peer", Text: s.discoverPeer},
		&pbg.State{Key: "register_peer", Text: s.registerPeer},
		&pbg.State{Key: "top_caller", Text: topCaller},
		&pbg.State{Key: "peer_fail", Text: s.peerFail},
		&pbg.State{Key: "top_requests", Text: fmt.Sprintf("%v (%v)", topRequest, topR)},
		&pbg.State{Key: "version", Text: fmt.Sprintf("%v", s.version)},
		&pbg.State{Key: "calls", Text: fmt.Sprintf("%v", s.getMap)},
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

var (
	fanout = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "discovery_fanout",
		Help: "The size of the print queue",
	}, []string{"service", "origin", "error"})
)

func (s *Server) fanoutRegister(ctx context.Context, req *pb.RegisterRequest) {
	for _, f := range s.friends {
		conn, err := s.FDial(f)
		if err == nil {
			defer conn.Close()
			client := pb.NewDiscoveryServiceV2Client(conn)
			_, err := client.RegisterV2(ctx, req)
			if err != nil {
				s.Log(fmt.Sprintf("register error: %v", err))
				fanout.With(prometheus.Labels{"service": req.GetService().GetName(), "origin": f, "error": fmt.Sprintf("%v", err)}).Inc()
			}
		} else {
			s.Log(fmt.Sprintf("Dial error in fanout register: %v", err))
			fanout.With(prometheus.Labels{"service": req.GetService().GetName(), "origin": f, "error": fmt.Sprintf("%v", err)}).Inc()
		}
	}
}
func (s *Server) fanoutMaster(ctx context.Context, req *pb.MasterRequest) {
	for _, f := range s.friends {
		conn, err := s.FDial(f)
		if err == nil {
			defer conn.Close()
			client := pb.NewDiscoveryServiceV2Client(conn)
			client.MasterElect(ctx, req)
		}
	}
}

func (s *Server) fanoutUnregister(ctx context.Context, req *pb.UnregisterRequest) {
	for _, f := range s.friends {
		conn, err := s.FDial(f)
		if err == nil {
			defer conn.Close()
			client := pb.NewDiscoveryServiceV2Client(conn)
			client.Unregister(ctx, req)
		}
	}
}

func (s *Server) acquireMasterLock(ctx context.Context, job string, lk int64) error {
	if s.failAcquire {
		return fmt.Errorf("Built to fail")
	}

	//Local lock first
	_, err := s.Lock(ctx, &pb.LockRequest{Job: job, LockKey: lk, Requestor: s.Registry.Identifier})
	if err != nil {
		return err
	}

	for _, f := range s.friends {
		conn, err := grpc.Dial(f, grpc.WithInsecure())
		if err == nil {
			defer conn.Close()
			client := pb.NewDiscoveryServiceV2Client(conn)
			_, err = client.Lock(ctx, &pb.LockRequest{Job: job, LockKey: lk, Requestor: s.Registry.Identifier})
			if err != nil {
				return err
			}
		}
	}
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

	server.RegisterServerV2("discovery", false, false)
	server.Registry.IgnoresMaster = true
	server.SendTrace = false

	// Find friends
	go func() {
		t := time.Now()
		time.Sleep(time.Second)
		for i := 1; i < 255; i++ {
			server.findFriend(i)
		}

		// Double check that we have everything
		server.validateFriends()
		server.friendTime = time.Now().Sub(t)
	}()
	server.Serve()
}
