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
		case "zone":
			var host = buildFlags.String("host", utils.Discover, "dicsover host")
			var zone = buildFlags.String("zone", "unknown", "zone")
			if err := buildFlags.Parse(os.Args[2:]); err == nil {
				conn, _ := grpc.Dial(*host, grpc.WithInsecure())
				defer conn.Close()

				registry := pbdi.NewDiscoveryServiceV2Client(conn)
				ctx, cancel := context.WithTimeout(context.Background(), time.Second)
				defer cancel()

				_, err := registry.SetZone(ctx, &pbdi.SetZoneRequest{Zone: *zone})
				if err != nil {
					log.Fatalf("Bad set: %v", err)
				}
			}
		case "friends":
			friendsFlags := flag.NewFlagSet("Friends", flag.ExitOnError)
			var friend = friendsFlags.String("friend", "", "Friend")
			if err := friendsFlags.Parse(os.Args[2:]); err == nil {
				conn, err := utils.LFDial(fmt.Sprintf("%v:50055", *friend))
				if err != nil {
					log.Fatalf("Dial err: %v", err)
				}
				defer conn.Close()

				registry := pbdi.NewDiscoveryServiceV2Client(conn)
				ctx, cancel := utils.ManualContext("discovery_cli-friends", time.Minute)
				defer cancel()
				friends, err := registry.GetFriends(ctx, &pbdi.GetFriendsRequest{})
				if err != nil {
					log.Fatalf("Error on get friends: %v", err)
				}
				for _, friend := range friends.GetFriends() {
					fmt.Printf("%v\n", friend)
				}
			}
		case "config":
			var host = buildFlags.String("host", utils.Discover, "dicsover host")
			if err := buildFlags.Parse(os.Args[2:]); err == nil {
				conn, _ := grpc.Dial(*host, grpc.WithInsecure())
				defer conn.Close()

				registry := pbdi.NewDiscoveryServiceV2Client(conn)
				ctx, cancel := context.WithTimeout(context.Background(), time.Second)
				defer cancel()
				state, err := registry.GetConfig(ctx, &pbdi.GetConfigRequest{})
				if err != nil {
					log.Fatalf("Error getting state: %v", err)
				}
				fmt.Printf("Config: \n")
				for f, v := range state.GetConfig().GetFriendState() {
					fmt.Printf("%v -> %v @ %v\n", f, v.GetState(), time.Unix(v.GetLastSeen(), 0))
				}
			}
		case "istate":
			conn, _ := grpc.Dial(utils.Discover, grpc.WithInsecure())
			defer conn.Close()

			registry := pbdi.NewDiscoveryServiceV2Client(conn)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			state, err := registry.GetInternalState(ctx, &pbdi.GetStateRequest{})
			if err != nil {
				log.Fatalf("Error getting state: %v", err)
			}
			fmt.Printf("STATE: %v\n", state)
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
			if err := buildFlags.Parse(os.Args[2:]); err == nil {
				conn, err := grpc.Dial(*host, grpc.WithInsecure())
				if err == nil {
					defer conn.Close()
					client := pbdi.NewDiscoveryServiceV2Client(conn)
					ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
					defer cancel()
					a, err := client.RegisterV2(ctx, &pbdi.RegisterRequest{Service: &pbdi.RegistryEntry{Name: *name, Identifier: *server, Version: pbdi.RegistryEntry_V2}})
					fmt.Printf("Registered: %v -> %v\n", err, a)
				}
			}
		case "unreg":
			var host = buildFlags.String("host", utils.Discover, "dicsover host")
			var name = buildFlags.String("name", "blah", "dicsover host")
			if err := buildFlags.Parse(os.Args[2:]); err == nil {
				conn, err := grpc.Dial(*host, grpc.WithInsecure())
				if err == nil {
					defer conn.Close()
					client := pbdi.NewDiscoveryServiceV2Client(conn)
					ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
					defer cancel()
					a, err := client.Unregister(ctx, &pbdi.UnregisterRequest{Service: &pbdi.RegistryEntry{Identifier: *name}})
					fmt.Printf("Unregistered: %v -> %v\n", err, a)
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
				conn, err := grpc.Dial(*host, grpc.WithInsecure())
				if err != nil {
					log.Fatalf("Cannot dial: %v", err)
				}
				defer conn.Close()

				registry := pbdi.NewDiscoveryServiceV2Client(conn)
				ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
				defer cancel()
				bits, err := registry.Get(ctx, &pbdi.GetRequest{})
				if err != nil {
					log.Fatalf("Error building job: %v", err)
				}
				slaves := make(map[string][]*pbdi.RegistryEntry)
				for _, bit := range bits.GetServices() {
					slaves[bit.Name] = append(slaves[bit.Name], bit)
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
		case "find":
			if err := buildFlags.Parse(os.Args[2:]); err == nil {
				for _, ip := range []string{
					"192.168.86.20:50055",
					"192.168.86.24:50055",
					"192.168.86.249:50055",
					"192.168.86.28:50055",
					"192.168.86.32:50055",
					"192.168.86.40:50055",
					"192.168.86.42:50055",
					"192.168.86.43:50055",
					"192.168.86.44:50055",
					"192.168.86.49:50055",
					"192.168.86.53:50055",
					"73.162.90.182:50055",
				} {
					conn, err := grpc.Dial(ip, grpc.WithInsecure())
					defer conn.Close()

					registry := pbdi.NewDiscoveryServiceV2Client(conn)
					ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
					defer cancel()
					bits, err := registry.Get(ctx, &pbdi.GetRequest{}, grpc.FailFast(false))

					if err != nil {
						log.Fatalf("Error in get: %v", err)
					}

					for _, bit := range bits.GetServices() {
						fmt.Printf("%v -> %v\n", ip, bit)
					}

				}
			}

		}
	}
}
