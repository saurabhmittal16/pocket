package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
)

const WORKER_PID_PATH string = "./.pocket/worker.pid"

var start bool
var stop bool

func init() {
	flag.BoolVar(&start, "start", false, "Start the server")
	flag.BoolVar(&stop, "exit", false, "Stop the server")
	flag.Parse()
}

func writePIDFile(pid int) {
	pidString := fmt.Sprintf("%v", pid)
	file, err := os.Create(WORKER_PID_PATH)
	if err != nil {
		log.Fatal("Failed to create pid file: ", err)
	}
	defer file.Close()

	_, err = file.WriteString(pidString)
	if err != nil {
		log.Fatal("Failed to create pid file: ", err)
	}
}

func readPIDFile() int {
	data, err := os.ReadFile(WORKER_PID_PATH)
	if err != nil {
		log.Fatal("Failed to read pid file: ", err)
	}
	pidString := string(data)
	pid, err := strconv.Atoi(pidString)
	if err != nil {
		log.Fatal("Failed to read pid file: ", err)
	}
	return pid
}

func startWorkerAndDetach() {
	cmd := exec.Command("go", "run", "./server/controller")
	log.Printf("Running worker node and detaching!")
	err := cmd.Start()

	if err != nil {
		log.Fatal("cmd.Start failed: ", err)
	}

	pid := cmd.Process.Pid
	log.Printf("Started the worker node at PID: %d\n", pid)

	writePIDFile(pid)
	_, err = cmd.Process.Wait()
	if err != nil {
		log.Fatal("cmd.Process.Release failed: ", err)
	}
}

func stopWorker() {
	pid := readPIDFile()
	proc, err := os.FindProcess(int(pid))
	if err != nil {
		log.Fatal("Failed to kill worker: ", err)
	}
	proc.Kill()
}

func main() {
	fmt.Print("Welcome to Pocket!\n")

	if start {
		startWorkerAndDetach()
	} else if stop {
		stopWorker()
	}
}
