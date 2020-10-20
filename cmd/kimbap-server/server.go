package main

import (
	"log"
	"net"

	pb "github.com/stopa323/kimbap/api/bmc"
	bmc "github.com/stopa323/kimbap/internal/bmc"
	"google.golang.org/grpc"
)

func main() {
	log.Println("Siema Heniu!")

	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterBMCServer(s, &bmc.BMCServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
