package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"golang.org/x/net/context"
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
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
					entry.Targets = append(entry.Targets, fmt.Sprintf("%v:%v", res.GetIdentifier(), res.GetPort()+2))
				}
			}

			if !found {
				entries = append(entries, &Entry{Targets: []string{fmt.Sprintf("%v:%v", res.GetIdentifier(), res.GetPort()+2)}, Labels: Label{Job: res.GetName()}})
			}
		}

		nodes := &Entry{Targets: []string{}, Labels: Label{Job: "node"}}
		for server, _ := range servers {
			// This is the raspberry pi export
			nodes.Targets = append(nodes.Targets, fmt.Sprintf("%v:9100", server))
		}
		nodes2 := &Entry{Targets: []string{}, Labels: Label{Job: "discovery"}}
		for server, _ := range servers {

			// Track discovery
			nodes2.Targets = append(nodes2.Targets, fmt.Sprintf("%v:50056", server))
		}
		entries = append(entries, nodes)
		entries = append(entries, nodes2)

		b, err := json.Marshal(entries)
		if err == nil {
			err = ioutil.WriteFile(os.Args[1], b, 0644)
			if err != nil {
				log.Printf("Write error: %v", err)
			}
		} else {
			log.Printf("Bad: %v", err)
		}
	} else {
		log.Printf("Bad: %v", err)
	}
}
