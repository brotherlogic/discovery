package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pbdi "github.com/brotherlogic/discovery/proto"
)

func main() {
	buildFlags := flag.NewFlagSet("BuildServer", flag.ExitOnError)

	if len(os.Args) <= 1 {
		fmt.Printf("Commands: list\n")
	} else {
		switch os.Args[1] {
		case "list":
			if err := buildFlags.Parse(os.Args[2:]); err == nil {
				conn, _ := grpc.Dial("192.168.86.64:50055", grpc.WithInsecure())
				defer conn.Close()

				registry := pbdi.NewDiscoveryServiceClient(conn)
				bits, err := registry.ListAllServices(context.Background(), &pbdi.Empty{})
				if err != nil {
					log.Printf("Error building job: %v", err)
				}
				for _, bit := range bits.Services {
					log.Printf("%v", bit)
				}
			}
		}
	}
}
