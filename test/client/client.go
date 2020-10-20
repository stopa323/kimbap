package main

import (
	"context"
	"log"
	"time"

	pb "github.com/stopa323/kimbap/api/bmc"
	"google.golang.org/grpc"
)

const (
	address = "localhost:8000"
)

func main() {
	log.Print("Client start...")
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewBMCClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetServerPowerStatus(ctx, &pb.BMCAccess{ConnectionString: "http://localhost:8001"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Status: %v", r.GetStatus())
}
