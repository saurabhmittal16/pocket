package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
)

const WORKER_PID_PATH string = "./.pocket/worker.pid"

var start bool
var stop bool

func init() {
	flag.BoolVar(&start, "start", false, "Start the server")
	flag.BoolVar(&stop, "exit", false, "Stop the server")
	flag.Parse()
}

func startWorkerAndDetach() {
	cmd := exec.Command("go", "run", "./server/controller")
	log.Printf("Running worker node and detaching!")
	err := cmd.Start()

	if err != nil {
		log.Fatal("cmd.Start failed: ", err)
	}

	err = cmd.Process.Release()
	if err != nil {
		log.Fatal("cmd.Process.Release failed: ", err)
	}
}

func stopWorker() {
	log.Printf("Stopping worker node!")
	cmd := exec.Command("fuser", "-k", "3000/tcp")
	err := cmd.Run()
	if err != nil {
		log.Fatal("Stop server failed: ", err)
	}
}

func main() {
	fmt.Print("Welcome to Pocket!\n")

	if start {
		startWorkerAndDetach()
	} else if stop {
		stopWorker()
	}
}
