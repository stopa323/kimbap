package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/stopa323/kimbap/api/bmc"
	bmc "github.com/stopa323/kimbap/internal/bmc"
	"google.golang.org/grpc"
)

var address string
var port string

func init() {
	flag.StringVar(&address, "address", "127.0.0.1",
		"gRPC server listen address")
	flag.StringVar(&port, "port", "8000", "gRPC server listen port")
}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", address, port))
	if err != nil {
		log.Fatalf("failed to listen on %s:%s: %v", address, port, err)
	}
	log.Printf("Listen on %s:%s", address, port)

	s := grpc.NewServer()
	pb.RegisterBMCServer(s, &bmc.BMCServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC: %v", err)
	}
}
