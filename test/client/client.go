package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/stopa323/kimbap/api/bmc"
	"google.golang.org/grpc"
)

var bmcAddress string
var operation string
var serverAddress string

func init() {
	flag.StringVar(&bmcAddress, "bmc-address", "http://127.0.0.1:8001",
		"Connection string to BMC server")
	flag.StringVar(&operation, "operation", "get-power-status",
		"Operation to invoke. One of [get-power-status | power-on | power-off]")
	flag.StringVar(&serverAddress, "server-address", "127.0.0.1:8000",
		"Connection string to gRPC server")
}

func main() {
	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.Dial(
		serverAddress,
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithTimeout(time.Duration(10)*time.Second))

	if err != nil {
		log.Fatalf("Cloud not connect to %v: %v", serverAddress, err)
	}

	defer conn.Close()
	c := pb.NewBMCClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	switch operation {
	case "get-power-status":
		log.Printf("Fetching power state...")
		r, err := c.GetServerPowerStatus(
			ctx, &pb.BMCAccess{ConnectionString: bmcAddress})
		if err != nil {
			log.Fatalf("Err: %v", err)
		}
		log.Printf("Result: %v", r.GetStatus())
	case "power-off":
		log.Printf("Powering off...")
		_, err := c.PowerServerOff(
			ctx, &pb.BMCAccess{ConnectionString: bmcAddress})
		if err != nil {
			log.Fatalf("Err: %v", err)
		}
	case "power-on":
		log.Printf("Powering on...")
		_, err := c.PowerServerOn(
			ctx, &pb.BMCAccess{ConnectionString: bmcAddress})
		if err != nil {
			log.Fatalf("Err: %v", err)
		}
	default:
		log.Fatalf("Unknown operation: %s", operation)
	}
}
