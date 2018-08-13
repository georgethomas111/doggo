package network

import (
	"reflect"
	"testing"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func TestPacketInfo(t *testing.T) {
	tests := []struct {
		packet  gopacket.Packet
		expInfo *info
	}{
		{
			gopacket.NewPacket([]byte{}, layers.LinkTypeEthernet, gopacket.DecodeOptions{Lazy: true}),
			&info{
				pMap: make(map[string]int),
			},
		},
	}

	for _, tc := range tests {
		i := packetInfo(tc.packet)
		if !reflect.DeepEqual(i, tc.expInfo) {
			t.Errorf("Expected %v, got %v.", tc.expInfo, i)
		}
	}
}
