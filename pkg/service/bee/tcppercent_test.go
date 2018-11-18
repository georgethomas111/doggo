package bee

import (
	"testing"

	"github.com/google/gopacket/layers"
)

func TestUpdate(t *testing.T) {
	tcpAck := &layers.TCP{
		ACK: true,
	}

	tcpSyn := &layers.TCP{
		SYN: true,
	}

	tests := []struct {
		tcp    *layers.TCP
		expPer float64
	}{
		{tcpAck, 100.00},
		{tcpSyn, 0.00},
	}

	for _, test := range tests {
		p := new(TCPPercent)
		p.Update(test.tcp)
		if test.expPer != p.Percentage {
			t.Errorf("Expected %v, got %v", test.expPer, p.Percentage)
		}
	}
}
