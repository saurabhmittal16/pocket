package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/saurabhmittal16/pocket/service"
	"google.golang.org/grpc"
)

var PORT int = 3000
var ADDR string = fmt.Sprintf(":%d", PORT)

type server struct {
	service.UnimplementedControllerServer
}

func main() {
	// create tcp listener at PORT
	lis, err := net.Listen("tcp", ADDR)
	if err != nil {
		log.Fatalf("[ERROR] failed to listen: %v", err)
	}

	s := grpc.NewServer()
	service.RegisterControllerServer(s, &server{})
	log.Printf("[LOG] controller server listening at: %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("[ERROR] failed to serve: %v", err)
	}
}

func (s *server) StartWorkers(ctx context.Context, in *service.WorkerRequest) (*service.WorkerReply, error) {
	log.Printf("[CALL] StartWorkers")
	message := fmt.Sprintf("Spinning up %d workers", in.GetNumWorkers())
	return &service.WorkerReply{Message: message}, nil
}
