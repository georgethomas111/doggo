package plot

import (
	"testing"
	"time"

	"github.com/google/gopacket/layers"
)

type mockPacketStats struct {
	layerInfo map[string]bool
	eth       *layers.Ethernet
	ipv4      *layers.IPv4
	ipv6      *layers.IPv6
	tcp       *layers.TCP
	udp       *layers.UDP
	errs      []error
}

func (m *mockPacketStats) LayerInfo() map[string]bool {
	return m.layerInfo
}

func (m *mockPacketStats) TCP() *layers.TCP {
	return m.tcp
}

func (m *mockPacketStats) UDP() *layers.UDP {
	return m.udp
}

func (m *mockPacketStats) Eth() *layers.Ethernet {
	return m.eth
}

func (m *mockPacketStats) IPv4() *layers.IPv4 {
	return m.ipv4
}

func (m *mockPacketStats) IPv6() *layers.IPv6 {
	return m.ipv6
}

func TestClient(t *testing.T) {
	tcpSyn := &layers.TCP{
		SYN: true,
	}

	tcpAck := &layers.TCP{
		ACK: true,
	}

	ps := new(mockPacketStats)

	plotCli := New()
	for i := 1; i < 60*60; i++ {
		time.Sleep(1 * time.Second)
		ps.tcp = tcpSyn
		if i%4 == 0 {
			ps.tcp = tcpAck
		}
		plotCli.Receive(ps)
	}
}
