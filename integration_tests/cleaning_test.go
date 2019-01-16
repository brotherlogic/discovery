package integration

import (
	"context"
	"log"
	"net"
	"strconv"
	"testing"
	"time"

	"google.golang.org/grpc"

	discovery "github.com/brotherlogic/discovery/core"
	pb "github.com/brotherlogic/discovery/proto"
)

// Get - return free open TCP port
func Get() (port string, err error) {
	ln, err := net.Listen("tcp", "[::]:0")
	if err != nil {
		return "", err
	}
	portn := ln.Addr().(*net.TCPAddr).Port
	err = ln.Close()
	return ":" + strconv.Itoa(portn), err
}

func runServer(testPort string) *grpc.Server {
	s := discovery.Serve(testPort)

	// Give the binary 10 seconds to become alive
	time.Sleep(time.Second * 2)

	log.Printf("Listening on port %v", testPort)

	return s
}

func register(port string, binary string, timeToLive int64) error {
	log.Printf("Registering on port %v", port)
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		return err
	}

	client := pb.NewDiscoveryServiceClient(conn)
	resp, err := client.RegisterService(context.Background(), &pb.RegisterRequest{Service: &pb.RegistryEntry{Name: binary, Identifier: "testing", TimeToClean: timeToLive}})

	log.Printf("Got register response: %v", resp)

	return err
}

func registerExternal(port string, binary string, timeToLive int64) error {
	log.Printf("Registering on port %v", port)
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		return err
	}

	client := pb.NewDiscoveryServiceClient(conn)
	resp, err := client.RegisterService(context.Background(), &pb.RegisterRequest{Service: &pb.RegistryEntry{Name: binary, Identifier: "testing", TimeToClean: timeToLive, ExternalPort: true}})

	log.Printf("Got register response: %v", resp)

	return err
}

func list(port string) (*pb.ListResponse, error) {
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		return nil, err
	}

	client := pb.NewDiscoveryServiceClient(conn)
	return client.ListAllServices(context.Background(), &pb.ListRequest{})
}

func TestPortClean(t *testing.T) {
	port, err := Get()
	if err != nil {
		log.Fatalf("Unable to find port: %v", err)
	}

	s := runServer(port)
	err = registerExternal(port, "test-binary", 3000)

	time.Sleep(time.Second)
	servers, err := list(port)
	if err != nil || len(servers.GetServices().GetServices()) == 0 {
		t.Fatalf("Unable to list: %v", err)
	}

	// Run the clean
	time.Sleep(time.Second * 4)

	servers, err = list(port)
	if err != nil || len(servers.GetServices().GetServices()) != 0 {
		t.Fatalf("Cleaning failed %v and %v", err, servers)
	}

	//Re-register
	err = registerExternal(port, "test-binary", 3000)
	servers, err = list(port)
	if err != nil || len(servers.GetServices().GetServices()) != 1 {
		t.Fatalf("Failed to bring server back up %v and %v", err, servers)
	}

	s.Stop()
}

func TestNormalClean(t *testing.T) {
	port, err2 := Get()
	if err2 != nil {
		log.Fatalf("Unable to find port: %v", err2)
	}

	s := runServer(port)

	err := register(port, "test-binary", 0)
	if err != nil {
		t.Fatalf("Unable to register: %v", err)
	}

	//Wait 1 second (no cleaning should happen)
	time.Sleep(time.Second)
	servers, err := list(port)
	if err != nil {
		t.Fatalf("Unable to list: %v", err)
	}
	if len(servers.GetServices().GetServices()) != 1 {
		t.Fatalf("Error running test, server is not registered: %v", servers)
	}

	//Wait for 4 seconds to allow time for cleaning
	time.Sleep(time.Second * 4)

	servers, err = list(port)
	if err != nil {
		t.Fatalf("Unable to list: %v", err)
	}
	if len(servers.GetServices().Services) > 0 {
		t.Errorf("Server has not been cleaned after 4 seconds: %v", servers)
	}

	log.Printf("Stopping discover service")
	s.Stop()
}

func TestLongClean(t *testing.T) {
	port, err2 := Get()
	if err2 != nil {
		log.Fatalf("Unable to find port: %v", err2)
	}
	s := runServer(port)

	err := register(port, "test-binary", 5000)
	if err != nil {
		t.Fatalf("Unable to register: %v", err)
	}

	//Wait 1 second (no cleaning should happen)
	time.Sleep(time.Second)
	servers, err := list(port)
	if err != nil {
		t.Fatalf("Unable to list: %v", err)
	}
	if len(servers.GetServices().GetServices()) != 1 {
		t.Fatalf("Error running test, server is not registered: %v", servers)
	}

	//Wait for 4 seconds and still should not be cleaned
	time.Sleep(time.Second * 3)

	servers, err = list(port)
	if err != nil {
		t.Fatalf("Unable to list: %v", err)
	}
	if len(servers.GetServices().Services) != 1 {
		t.Errorf("Server has been cleaned after 4 seconds: %v", servers)
	}

	//Wait for 4 seconds to allow time for cleaning
	time.Sleep(time.Second * 2)

	servers, err = list(port)
	if err != nil {
		t.Fatalf("Unable to list: %v", err)
	}
	if len(servers.GetServices().Services) > 0 {
		t.Errorf("Server has not been cleaned after 5 seconds: %v", servers)
	}

	s.Stop()
}
