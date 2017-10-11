package main

import (
	"context"
	"flag"
	"io/ioutil"
	"log"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc"

	pb "github.com/brotherlogic/discovery/proto"
	pbg "github.com/brotherlogic/goserver/proto"
	pbm "github.com/brotherlogic/monitor/monitorproto"
)

type prodHealthChecker struct{}

func (s *Server) recordTime(fName string, t time.Duration) {
	for _, e := range s.entries {
		if e.GetName() == "monitor" && e.GetMaster() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			conn, err := grpc.DialContext(ctx, e.Ip+":"+strconv.Itoa(int(e.Port)), grpc.WithInsecure())
			if err != nil {
				log.Printf("Can't event dial %v -> %v", e, err)
			} else {
				defer conn.Close()

				client := pbm.NewMonitorServiceClient(conn)
				client.WriteFunctionCall(ctx, &pbm.FunctionCall{Binary: "discovery", Name: fName, Time: int32(t.Nanoseconds() / 1000000)})
			}
		}
	}
}

func (healthChecker prodHealthChecker) Check(count int, entry *pb.RegistryEntry) int {
	t2 := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	conn, err := grpc.DialContext(ctx, entry.Ip+":"+strconv.Itoa(int(entry.Port)), grpc.WithInsecure())
	if err != nil {
		log.Printf("Can't event dial %v -> %v", entry, err)
		return count + 1
	}
	defer conn.Close()

	t1 := time.Now()
	registry := pbg.NewGoserverServiceClient(conn)
	resp, err := registry.IsAlive(ctx, &pbg.Alive{})
	if err != nil {
		log.Printf("Error reading health of %v -> %v given %v (%v) seconds", entry, err, time.Now().Sub(t1), time.Now().Sub(t2))
		return count + 1
	}

	if resp.Name != entry.GetName() {
		log.Printf("Wrong name on return: %v", resp)
		return count + 1
	}

	return 0
}

// Serve main server function
func Serve() {
	var quiet = flag.Bool("quiet", false, "Show all output")
	flag.Parse()
	if *quiet {
		log.SetFlags(0)
		log.SetOutput(ioutil.Discard)
	}
	log.Printf("Logging is on!")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Unable to get tcp port %v -> %v", port, err)
	}
	s := grpc.NewServer()
	server := InitServer()
	pb.RegisterDiscoveryServiceServer(s, &server)
	err = s.Serve(lis)
	log.Printf("Failed to serve: %v", err)
}

func main() {
	Serve()
}
