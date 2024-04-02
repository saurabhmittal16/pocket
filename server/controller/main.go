package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/saurabhmittal16/pocket/service"
	"google.golang.org/grpc"
)

var PORT int = 3000
var ADDR string = fmt.Sprintf(":%d", PORT)

type server struct {
	pb.UnimplementedControllerServer
}

func main() {
	// create tcp listener at PORT
	lis, err := net.Listen("tcp", ADDR)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterControllerServer(s, &server{})
	log.Printf("controller server listening at: %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// func CheckControlNodeRunning() bool {
// 	_, err := net.Listen("tcp", ADDR)
// 	return err != nil
// }
