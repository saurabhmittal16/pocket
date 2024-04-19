package client

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/saurabhmittal16/pocket/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var PORT int = 3000
var ADDR string = fmt.Sprintf(":%d", PORT)
var conn *grpc.ClientConn
var client service.ControllerClient
var ctx context.Context
var cancel context.CancelFunc

func setup() {
	// Set up a connection to the server.
	var err error
	conn, err = grpc.NewClient(ADDR, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to Controller: %v", err)
	}
	client = service.NewControllerClient(conn)

	// Contact the server and print out its response.
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
}

func teardown() {
	conn.Close()
	cancel()
}

func SpinNodes(count int32) {
	setup()
	r, err := client.StartWorkers(ctx, &service.WorkerRequest{NumWorkers: count})
	if err != nil {
		log.Fatalf("Could not spin up nodes: %v", err)
	}
	log.Printf("Response: %s", r.GetMessage())
	teardown()
}

func Start() {
	f, err := os.OpenFile("./logs/controller", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Unable to open log file")
	}

	cmd := exec.Command("go", "run", "./controller/server")
	cmd.Stderr = f
	cmd.Stdout = f

	log.Printf("Running worker node and detaching!")
	err = cmd.Start()

	if err != nil {
		log.Fatal("cmd.Start failed: ", err)
	}

	err = cmd.Process.Release()
	if err != nil {
		log.Fatal("cmd.Process.Release failed: ", err)
	}
}

func Stop() {
	log.Printf("Stopping worker node!")
	cmd := exec.Command("fuser", "-k", "3000/tcp")
	err := cmd.Run()
	if err != nil {
		log.Fatal("Stop server failed: ", err)
	}
}
