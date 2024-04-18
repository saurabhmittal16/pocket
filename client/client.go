package client

import (
	"context"
	"log"
	"os/exec"
	"time"

	"github.com/saurabhmittal16/pocket/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func setup() (*grpc.ClientConn, service.ControllerClient, context.Context, context.CancelFunc) {
	// Set up a connection to the server.
	conn, err := grpc.NewClient("localhost:3000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to Controller: %v", err)
	}
	client := service.NewControllerClient(conn)

	// Contact the server and print out its response.
	context, cancel := context.WithTimeout(context.Background(), time.Second)

	return conn, client, context, cancel
}

func teardown(conn *grpc.ClientConn, cancel context.CancelFunc) {
	conn.Close()
	cancel()
}

func SpinNodes(count int32) {
	conn, client, context, cancel := setup()
	r, err := client.StartWorkers(context, &service.WorkerRequest{NumWorkers: count})
	if err != nil {
		log.Fatalf("Could not spin up nodes: %v", err)
	}
	log.Printf("Response: %s", r.GetMessage())
	teardown(conn, cancel)
}

func Start() {
	cmd := exec.Command("go", "run", "../server/controller")
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

func Stop() {
	log.Printf("Stopping worker node!")
	cmd := exec.Command("fuser", "-k", "3000/tcp")
	err := cmd.Run()
	if err != nil {
		log.Fatal("Stop server failed: ", err)
	}
}
