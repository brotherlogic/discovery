package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc"

	pb "github.com/brotherlogic/discovery/proto"
	pbg "github.com/brotherlogic/goserver/proto"
	pbm "github.com/brotherlogic/monitor/monitorproto"
)

type prodHealthChecker struct {
	logger func(logd string)
}

func (s *Server) recordTime(fName string, t time.Duration) {
	for _, e := range s.entries {
		if e.GetName() == "monitor" && e.GetMaster() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			conn, err := grpc.DialContext(ctx, e.Ip+":"+strconv.Itoa(int(e.Port)), grpc.WithInsecure())
			if err == nil {
				defer conn.Close()

				client := pbm.NewMonitorServiceClient(conn)
				client.WriteFunctionCall(ctx, &pbm.FunctionCall{Binary: "discovery", Name: fName, Time: int32(t.Nanoseconds() / 1000000)}, grpc.FailFast(false))
			}
		}
	}
}

func (s *Server) recordLog(logDetail string) {
	go func() {
		for _, e := range s.entries {
			if e.GetName() == "monitor" && e.GetMaster() {
				ctx, cancel := context.WithTimeout(context.Background(), time.Second)
				defer cancel()
				conn, err := grpc.DialContext(ctx, e.Ip+":"+strconv.Itoa(int(e.Port)), grpc.WithInsecure())
				if err == nil {
					defer conn.Close()

					client := pbm.NewMonitorServiceClient(conn)
					client.WriteMessageLog(ctx, &pbm.MessageLog{Entry: &pb.RegistryEntry{Name: "discovery"}, Message: logDetail}, grpc.FailFast(false))
				}
			}
		}
	}()
}

func (healthChecker prodHealthChecker) Check(count int, entry *pb.RegistryEntry) int {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*200)
	defer cancel()
	conn, err := grpc.DialContext(ctx, entry.Ip+":"+strconv.Itoa(int(entry.Port)), grpc.WithInsecure())
	if err != nil {
		return count + 1
	}
	defer conn.Close()

	registry := pbg.NewGoserverServiceClient(conn)
	resp, err := registry.IsAlive(ctx, &pbg.Alive{}, grpc.FailFast(false))
	if err != nil {
		healthChecker.logger(fmt.Sprintf("Alive fail: %v", err))
		return count + 1
	}

	if resp.Name != entry.GetName() {
		return count + 1
	}

	return 0
}

// Serve main server function
func Serve() {
	fmt.Printf("Logging is onn")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Unable to get tcp port %v -> %v", port, err)
	}
	s := grpc.NewServer()
	server := InitServer()
	pb.RegisterDiscoveryServiceServer(s, &server)

	go func() {
		for true {
			time.Sleep(time.Second * 5)
			server.cleanEntries()
		}
	}()

	server.setExternalIP(prodHTTPGetter{})

	err = s.Serve(lis)
	fmt.Printf("Failed to serve: %v\n", err)
}

func main() {
	Serve()
}
