package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pbdi "github.com/brotherlogic/discovery/proto"
	"github.com/brotherlogic/goserver/utils"

	_ "google.golang.org/grpc/encoding/gzip"
)

func repEntry(entry *pbdi.RegistryEntry) string {
	return fmt.Sprintf("%v - %v", entry.Name, entry.Identifier)
}

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
			conn, _ := grpc.Dial(utils.Discover, grpc.WithInsecure())
			defer conn.Close()

			registry := pbdi.NewDiscoveryServiceClient(conn)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			state, err := registry.State(ctx, &pbdi.StateRequest{})
			if err != nil {
				log.Fatalf("Error getting state: %v", err)
			}
			fmt.Printf("STATE: %v\n", state)
		case "statechange":
			conn, _ := grpc.Dial(utils.Discover, grpc.WithInsecure())
			defer conn.Close()

			registry := pbdi.NewDiscoveryServiceClient(conn)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()
			state1, err := registry.State(ctx, &pbdi.StateRequest{})
			if err != nil {
				log.Fatalf("Error getting state: %v", err)
			}

			time.Sleep(time.Second * 5)
			state2, err := registry.State(ctx, &pbdi.StateRequest{})
			if err != nil {
				log.Fatalf("Error getting state: %v", err)
			}

			fmt.Printf("STATE: %v\n", state1)
			fmt.Printf("STATE: %v\n", state2)
		case "reg":
			var host = buildFlags.String("host", utils.Discover, "dicsover host")
			var server = buildFlags.String("server", "192.168.86.32", "dicsover host")
			var name = buildFlags.String("name", "blah", "dicsover host")
			var master = buildFlags.Bool("master", false, "master")
			if err := buildFlags.Parse(os.Args[2:]); err == nil {
				conn, err := grpc.Dial(*host, grpc.WithInsecure())
				if err == nil {
					defer conn.Close()
					client := pbdi.NewDiscoveryServiceV2Client(conn)
					ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
					defer cancel()
					if *master {
						a, err := client.MasterElect(ctx, &pbdi.MasterRequest{Service: &pbdi.RegistryEntry{Name: *name, Identifier: *server, Version: pbdi.RegistryEntry_V2}, MasterElect: true})
						fmt.Printf("Registered: %v -> %v\n", err, a)
					} else {
						a, err := client.RegisterV2(ctx, &pbdi.RegisterRequest{Service: &pbdi.RegistryEntry{Name: *name, Identifier: *server, Version: pbdi.RegistryEntry_V2}})
						fmt.Printf("Registered: %v -> %v\n", err, a)
					}
				}
			}
		case "get":
			var host = buildFlags.String("host", utils.Discover, "dicsover host")
			var name = buildFlags.String("name", "", "name")
			if err := buildFlags.Parse(os.Args[2:]); err == nil {
				conn, _ := grpc.Dial(*host, grpc.WithInsecure())
				defer conn.Close()

				registry := pbdi.NewDiscoveryServiceV2Client(conn)
				ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
				defer cancel()
				resp, err := registry.Get(ctx, &pbdi.GetRequest{Job: *name})
				if err != nil {
					fmt.Printf("Get error: %v", err)
				} else {
					for _, service := range resp.GetServices() {
						fmt.Printf("%v\n", service)
					}
				}
			}
		case "blist":
			var host = buildFlags.String("host", utils.Discover, "dicsover host")
			if err := buildFlags.Parse(os.Args[2:]); err == nil {
				conn, _ := grpc.Dial(*host, grpc.WithInsecure())
				defer conn.Close()

				registry := pbdi.NewDiscoveryServiceClient(conn)
				ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
				defer cancel()
				bits, err := registry.ListAllServices(ctx, &pbdi.ListRequest{}, grpc.FailFast(false))
				if err != nil {
					log.Fatalf("Error building job: %v", err)
				}
				for _, bit := range bits.GetServices().Services {
					fmt.Printf("%v\n", bit)
				}
			}
		case "list":
			var host = buildFlags.String("host", utils.Discover, "host")
			if err := buildFlags.Parse(os.Args[2:]); err == nil {
				conn, _ := grpc.Dial(*host, grpc.WithInsecure())
				defer conn.Close()

				registry := pbdi.NewDiscoveryServiceV2Client(conn)
				ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
				defer cancel()
				bits, err := registry.Get(ctx, &pbdi.GetRequest{}, grpc.FailFast(false))
				if err != nil {
					log.Fatalf("Error building job: %v", err)
				}
				fmt.Printf("MASTERS\n-------\n")
				masters := []string{}
				for _, bit := range bits.GetServices() {
					regtime := time.Now().Sub(time.Unix(0, bit.GetRegisterTime())).Truncate(time.Minute)
					mastertime := time.Now().Sub(time.Unix(0, bit.GetMasterTime())).Truncate(time.Minute)
					if bit.GetMaster() {
						masters = append(masters, fmt.Sprintf("%v [%v - %v]", repEntry(bit), mastertime, regtime))
					}
				}
				sort.Strings(masters)
				for _, str := range masters {
					fmt.Printf("%v\n", str)
				}

				fmt.Printf("\nSLAVES\n-------\n")
				slaves := make(map[string][]*pbdi.RegistryEntry)
				for _, bit := range bits.GetServices() {
					if !bit.GetMaster() {
						slaves[bit.Name] = append(slaves[bit.Name], bit)
					}
				}
				keys := []string{}
				for k := range slaves {
					keys = append(keys, k)
				}
				sort.Strings(keys)
				for _, key := range keys {
					sort.SliceStable(slaves[key], func(i, j int) bool {
						return slaves[key][i].Identifier < slaves[key][j].Identifier
					})
					regtime := time.Now().Sub(time.Unix(0, slaves[key][0].GetRegisterTime())).Truncate(time.Minute)
					fmt.Printf("%v {%v} [%v]\n", repEntry(slaves[key][0]), len(slaves[key]), regtime)
				}

			}
		}
	}
}
