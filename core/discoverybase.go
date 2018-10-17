package discovery

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"google.golang.org/grpc"

	pb "github.com/brotherlogic/discovery/proto"
	pbm "github.com/brotherlogic/monitor/monitorproto"
)

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

// Serve main server function
func Serve(port string) *grpc.Server {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Unable to get tcp port %v -> %v", port, err)
	}
	s := grpc.NewServer()
	server := InitServer()
	pb.RegisterDiscoveryServiceServer(s, &server)

	go func() {
		for true {
			time.Sleep(time.Second)
			server.cleanEntries(time.Now())
		}
	}()

	server.setExternalIP(prodHTTPGetter{})

	go func() {
		err = s.Serve(lis)
		if err != nil {
			fmt.Printf("Server Response: %v\n", err)
		}
		os.Exit(0)
	}()

	return s
}
