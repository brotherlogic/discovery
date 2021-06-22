package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	pb "github.com/brotherlogic/discovery/proto"
)

type Entry struct {
	Targets []string `json:"targets"`
	Labels  Label    `json:"labels"`
}

type Label struct {
	Job string `json:"job"`
}

// This writes out the file for prometheus
func (s *Server) writeFile(f string, services []*pb.RegistryEntry) error {
	var entries []*Entry
	servers := make(map[string]bool)

	for _, res := range services {
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
	entries = append(entries, nodes)

	nodes2 := &Entry{Targets: []string{}, Labels: Label{Job: "discovery"}}
	for server, _ := range servers {
		// Track discovery entries
		nodes2.Targets = append(nodes2.Targets, fmt.Sprintf("%v:50057", server))
	}
	entries = append(entries, nodes2)

	b, err := json.Marshal(entries)
	if err == nil {
		err = ioutil.WriteFile(f, b, 0644)
	}

	return err
}
