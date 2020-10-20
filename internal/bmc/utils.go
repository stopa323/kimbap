package bmc

import (
	pb "github.com/stopa323/kimbap/api/bmc"
)

func ConvertGofishPowerStateToProto(status string) pb.PowerStatus {
	switch status {
	case "On":
		return pb.PowerStatus_ON
	case "Off":
		return pb.PowerStatus_OFF
	case "PoweringOn":
		return pb.PowerStatus_POWERING_ON
	case "PoweringOff":
		return pb.PowerStatus_POWERING_OFF
	default:
		return pb.PowerStatus_UNKNOWN
	}
}
