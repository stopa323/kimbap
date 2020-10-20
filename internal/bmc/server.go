package bmc

import (
	"context"
	"fmt"
	"log"

	gf "github.com/stmcginnis/gofish"
	pb "github.com/stopa323/kimbap/api/bmc"
)

type BMCServer struct {
	pb.UnimplementedBMCServer
}

func (s *BMCServer) GetServerPowerStatus(ctx context.Context, bmc *pb.BMCAccess) (*pb.ServerPowerStatus, error) {
	log.Printf("get power status: %v", bmc.GetConnectionString())

	c, err := gf.ConnectDefault(bmc.GetConnectionString())
	if err != nil {
		return nil, fmt.Errorf("could not connect to: %v", bmc.GetConnectionString())
	}

	computerSystems, err := c.Service.Systems()
	if err != nil {
		return nil, fmt.Errorf("could not find any system")
	}

	if 1 != len(computerSystems) {
		return &pb.ServerPowerStatus{Status: "UNKNOWN"}, nil
	}

	return &pb.ServerPowerStatus{Status: string(computerSystems[0].PowerState)}, nil
}
