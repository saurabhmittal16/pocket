package main

import (
	"flag"
	"fmt"

	"github.com/saurabhmittal16/pocket/client"
)

var start bool
var stop bool

func init() {
	flag.BoolVar(&start, "start", false, "Start the server")
	flag.BoolVar(&stop, "exit", false, "Stop the server")
	flag.Parse()
}

func main() {
	fmt.Print("Welcome to Pocket!\n")

	// if start {
	// 	client.Start()
	// } else if stop {
	// 	client.Stop()
	// }
	client.Start()
	// client.SpinNodes(3)
}
