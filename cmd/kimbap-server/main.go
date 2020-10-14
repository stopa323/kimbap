package main

import (
	"context"
	"log"
	"net"

	pb "github.com/stopa323/kimbap/api/bmc"
	"google.golang.org/grpc"
)

type kimbapServer struct {
	pb.UnimplementedBMCServer
}

func (s *kimbapServer) PowerOn(ctx context.Context, req *pb.BMCOperationRequest) (*pb.BMCOperationResponse, error) {
	log.Printf("PowerON: %v", req.GetBmcAddress())
	return &pb.BMCOperationResponse{Status: true}, nil
}

func main() {
	log.Println("Siema Heniu!")

	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterBMCServer(s, &kimbapServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
