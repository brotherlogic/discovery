package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/brotherlogic/goserver/utils"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pbdi "github.com/brotherlogic/discovery/proto"
)

func main() {

	fails := 0
	arr := []int{1, 2, 3, 4}
	for i, val := range arr {
		if val == 4 {
			arr = append(arr[:(i-fails)], arr[(i-fails)+1:]...)
		}
	}

	buildFlags := flag.NewFlagSet("BuildServer", flag.ExitOnError)

	if len(os.Args) <= 1 {
		fmt.Printf("Commands: list\n")
	} else {
		switch os.Args[1] {
		case "list":
			if err := buildFlags.Parse(os.Args[2:]); err == nil {
				conn, _ := grpc.Dial(utils.RegistryIP+":"+strconv.Itoa(utils.RegistryPort), grpc.WithInsecure())
				defer conn.Close()

				registry := pbdi.NewDiscoveryServiceClient(conn)
				bits, err := registry.ListAllServices(context.Background(), &pbdi.Empty{}, grpc.FailFast(false))
				if err != nil {
					log.Fatalf("Error building job: %v", err)
				}
				fmt.Printf("MASTERS\n-------\n")
				for _, bit := range bits.Services {
					if bit.GetMaster() {
						fmt.Printf("%v\n", bit)
					}
				}
				fmt.Printf("SLAVES\n-------\n")
				for _, bit := range bits.Services {
					if !bit.GetMaster() {
						fmt.Printf("%v\n", bit)
					}
				}
			}
		}
	}
}
