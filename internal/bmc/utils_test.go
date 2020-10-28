package bmc

import (
	"testing"

	pb "github.com/stopa323/kimbap/api/bmc"
)

func TestGofishPowerStateConversion(t *testing.T) {
	tables := []struct {
		status   string
		expected pb.PowerStatus
	}{
		{"On", pb.PowerStatus_ON},
		{"Off", pb.PowerStatus_OFF},
		{"PoweringOn", pb.PowerStatus_POWERING_ON},
		{"PoweringOff", pb.PowerStatus_POWERING_OFF},
		{"Everybody be cool", pb.PowerStatus_UNKNOWN},
	}

	for _, table := range tables {
		status := ConvertGofishPowerStateToProto(table.status)
		if status != table.expected {
			t.Errorf("Incorrect conversion of %s, got: %d, want: %d.",
				table.status, status, table.expected)
		}
	}
}
