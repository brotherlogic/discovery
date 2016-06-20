package main

import (
	"testing"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/brotherlogic/discovery/proto"	
)

func TestRegisterService(t *testing.T) {
     s := InitServer()
     entry := &pb.RegistryEntry{}
     s.RegisterService(context.Background(), entry)
}

func TestDiscover(t *testing.T) {
     s := InitServer()
     entry := &pb.RegistryEntry{}
     s.Discover(context.Background(), entry)
}

func TestRunServer(t *testing.T) {
        Serve()

	go func() {
		conn, err := grpc.Dial("localhost:50055", grpc.WithInsecure())
				if err != nil {
			t.Errorf("Error connecting to port")
		}

		defer conn.Close()
		client := pb.NewDiscoveryServiceClient(conn)

		entry := pb.RegistryEntry{}

		_, err = client.RegisterService(context.Background(), &entry)
		if err != nil {
			t.Errorf("Error registering service: %v", err)
		}

		_, err = client.Discover(context.Background(), &entry)
		if err != nil {
			t.Errorf("Error performing discovery: %v", err)
		}

	}()

	time.Sleep(10 * time.Second)
}

func TestMainForCoverage(t *testing.T) {
     main()
}