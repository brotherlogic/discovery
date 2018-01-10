package main

import (
	"time"

	discovery "github.com/brotherlogic/discovery/core"
)

const (
	port = ":50055"
)

func main() {
	discovery.Serve(port)
	for true {
		time.Sleep(time.Minute)
	}
}
