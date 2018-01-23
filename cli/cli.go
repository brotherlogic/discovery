package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/brotherlogic/goserver/utils"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pbdi "github.com/brotherlogic/discovery/proto"

	_ "google.golang.org/grpc/encoding/gzip"
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
		case "state":
			conn, _ := grpc.Dial(utils.RegistryIP+":"+strconv.Itoa(utils.RegistryPort), grpc.WithInsecure())
			defer conn.Close()

			registry := pbdi.NewDiscoveryServiceClient(conn)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			state, err := registry.State(ctx, &pbdi.StateRequest{})
			if err != nil {
				log.Fatalf("Error getting state: %v", err)
			}
			fmt.Printf("STATE: %v\n", state)
		case "list":
			if err := buildFlags.Parse(os.Args[2:]); err == nil {
				conn, _ := grpc.Dial(utils.RegistryIP+":"+strconv.Itoa(utils.RegistryPort), grpc.WithInsecure())
				defer conn.Close()

				registry := pbdi.NewDiscoveryServiceClient(conn)
				ctx, cancel := context.WithTimeout(context.Background(), time.Second)
				defer cancel()
				bits, err := registry.ListAllServices(ctx, &pbdi.Empty{}, grpc.FailFast(false))
				if err != nil {
					log.Fatalf("Error building job: %v", err)
				}
				fmt.Printf("MASTERS\n-------\n")
				for _, bit := range bits.Services {
					regtime := time.Unix(bit.GetLastSeenTime(), 0).Sub(time.Unix(bit.GetRegisterTime(), 0))
					mastertime := time.Unix(bit.GetLastSeenTime(), 0).Sub(time.Unix(bit.GetMasterTime(), 0))
					if bit.GetMaster() {
						fmt.Printf("%v [%v - %v]\n", bit, mastertime, regtime)
					}
				}
				fmt.Printf("SLAVES\n-------\n")
				for _, bit := range bits.Services {
					regtime := time.Unix(bit.GetLastSeenTime(), 0).Sub(time.Unix(bit.GetRegisterTime(), 0))
					if !bit.GetMaster() {
						fmt.Printf("%v [%v]\n", bit, regtime)
					}
				}
			}
		}
	}
}
