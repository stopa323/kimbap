package bmc

import (
	"context"
	"fmt"
	"log"

	gf "github.com/stmcginnis/gofish"
	"github.com/stmcginnis/gofish/redfish"
	_ "github.com/stmcginnis/gofish/redfish"
	pb "github.com/stopa323/kimbap/api/bmc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BMCServer struct {
	pb.UnimplementedBMCServer
}

func (s *BMCServer) GetServerPowerStatus(
	ctx context.Context,
	bmc *pb.BMCAccess) (
	*pb.ServerPowerStatusResponse, error) {
	log.Printf("get power status: %v", bmc.GetConnectionString())

	c, err := gf.ConnectDefault(bmc.GetConnectionString())
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unable to connect to %v",
				bmc.GetConnectionString()))
	}

	computerSystems, err := c.Service.Systems()
	if err != nil {
		return nil, status.Errorf(
			codes.Internal, "Unable to fetch computer systems")
	}

	if 1 != len(computerSystems) {
		return &pb.ServerPowerStatusResponse{
				Status: pb.PowerStatus_UNKNOWN},
			nil
	}

	return &pb.ServerPowerStatusResponse{
		Status: ConvertGofishPowerStateToProto(string(
			computerSystems[0].PowerState))}, nil
}

func (s *BMCServer) PowerServerOff(
	ctx context.Context,
	bmc *pb.BMCAccess) (*pb.Empty, error) {
	c, err := gf.ConnectDefault(bmc.GetConnectionString())
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unable to connect to %v",
				bmc.GetConnectionString()))
	}
	defer c.Logout()

	// Query the computer systems
	systems, err := c.Service.Systems()
	if err != nil {
		return nil, status.Errorf(
			codes.Internal, "Unable to fetch computer systems")
	}

	// Each computer system that is found is gracefuly shut down
	for _, system := range systems {
		err = system.Reset(redfish.GracefulShutdownResetType)
		if err != nil {
			return nil, status.Errorf(
				codes.Internal,
				fmt.Sprintf("Failed to shut down computer: %v", err))
		}
	}

	// Empty message on success
	return &pb.Empty{}, nil
}
