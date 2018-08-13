package tcp

import (
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
			gopacket.NewPacket([]byte{}, layers.LinkTypeEthernet, gopacket.Default),
			&info{
				pMap: map[string]int{
					"Ethernet": 1,
				},
			},
		},
	}

	for _, tc := range tests {
		i := packetInfo(tc.packet)
		if i != tc.expInfo {
			t.Errorf("Expected %v, got %v.", tc.expInfo, i)
		}
	}
}
