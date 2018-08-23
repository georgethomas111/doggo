package stats

import "github.com/google/gopacket/layers"

type Packet interface {
	LayerInfo() map[string]bool
	Eth() *layers.Ethernet
	IPv4() *layers.IPv4
	IPv6() *layers.IPv6
	TCP() *layers.TCP
	UDP() *layers.UDP
}
