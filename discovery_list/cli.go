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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	resp, err := registry.Get(ctx, &pbdi.GetRequest{})
	if err == nil {
		var entries []*Entry

		for _, res := range resp.GetServices() {
			found := false
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
		b, err := json.Marshal(entries)
		if err == nil {
			err = ioutil.WriteFile(os.Args[1], b, 0644)
			log.Printf("Written: %v", err)
		} else {
			log.Printf("Bad: %v", err)
		}
	} else {
		log.Printf("Bad: %v", err)
	}
}
