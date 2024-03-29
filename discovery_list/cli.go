package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"google.golang.org/grpc"

	pbdi "github.com/brotherlogic/discovery/proto"
	"github.com/brotherlogic/goserver/utils"

	_ "google.golang.org/grpc/encoding/gzip"
)

type Entry struct {
	Targets []string `json:"targets"`
	Labels  Label    `json:"labels"`
}

type Label struct {
	Job string `json:"job"`
}

func main() {
	conn, _ := grpc.Dial(utils.Discover, grpc.WithInsecure())
	defer conn.Close()

	registry := pbdi.NewDiscoveryServiceV2Client(conn)
	ctx, cancel := utils.ManualContext("discovery-list", time.Minute)
	defer cancel()
	resp, err := registry.Get(ctx, &pbdi.GetRequest{})
	if err == nil {
		var entries []*Entry
		servers := make(map[string]bool)

		for _, res := range resp.GetServices() {
			found := false
			servers[res.GetIdentifier()] = true
			for _, entry := range entries {
				if entry.Labels.Job == res.GetName() {
					found = true
					entry.Targets = append(entry.Targets, fmt.Sprintf("%v:%v", res.GetIdentifier(), res.GetPort()+1))
				}
			}

			if !found {
				entries = append(entries, &Entry{Targets: []string{fmt.Sprintf("%v:%v", res.GetIdentifier(), res.GetPort()+1)}, Labels: Label{Job: res.GetName()}})
			}
		}

		nodes := &Entry{Targets: []string{}, Labels: Label{Job: "node"}}
		for server, _ := range servers {
			// This is the raspberry pi export
			nodes.Targets = append(nodes.Targets, fmt.Sprintf("%v:9100", server))
						nodes.Targets = append(nodes.Targets, fmt.Sprintf("%v:9110", server))	
		}
		for val := 1; val <= 4; val++ {
			nodes.Targets = append(nodes.Targets, fmt.Sprintf("kclust%v:9110", val))
		}

		nodes2 := &Entry{Targets: []string{}, Labels: Label{Job: "discovery"}}
		for server := range servers {
			// Track discovery
			nodes2.Targets = append(nodes2.Targets, fmt.Sprintf("%v:50056", server))
		}
		entries = append(entries, nodes)
		entries = append(entries, nodes2)

		b, err := json.Marshal(entries)
		if err == nil {
			err = ioutil.WriteFile(os.Args[1], b, 0644)
		}
	}
}
