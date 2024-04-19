package main

import (
	"flag"
	"fmt"

	"github.com/saurabhmittal16/pocket/client"
)

var start bool
var stop bool
var check bool
var spin bool

func init() {
	flag.BoolVar(&start, "start", false, "Start the controller node")
	flag.BoolVar(&stop, "stop", false, "Stop the controller node")
	flag.BoolVar(&check, "check", false, "Check controller node status")
	flag.BoolVar(&spin, "spin", false, "Spin worker nodes")

	flag.Parse()
}

func main() {
	fmt.Print("Welcome to Pocket!\n")

	if start {
		client.Start()
	} else if stop {
		client.Stop()
	} else if check {
		client.CheckStatus()
	} else if spin {
		client.SpinNodes(3)
	}
}
