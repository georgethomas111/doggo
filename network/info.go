package network

import (
	"errors"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

// info contains the information from a packet.
type info struct {
	layerInfo map[string]bool
	eth       *layers.Ethernet
	ipv4      *layers.IPv4
	ipv6      *layers.IPv6
	tcp       *layers.TCP
	udp       *layers.UDP
	errs      []error
}

func (i *info) LayerInfo() map[string]bool {
	return i.layerInfo
}

func (i *info) TCP() *layers.TCP {
	return i.tcp
}

func (i *info) UDP() *layers.UDP {
	return i.udp
}

func (i *info) Eth() *layers.Ethernet {
	return i.eth
}

func (i *info) IPv4() *layers.IPv4 {
	return i.ipv4
}

func (i *info) IPv6() *layers.IPv6 {
	return i.ipv6
}

func packetInfo(p gopacket.Packet) *info {
	pLayers := p.Layers()

	var eth *layers.Ethernet
	var ipv4 *layers.IPv4
	var ipv6 *layers.IPv6
	var tcp *layers.TCP
	var udp *layers.UDP
	var errs []error

	layerInfo := make(map[string]bool)

	var ok bool
	for _, l := range pLayers {
		lt := l.LayerType()
		switch lt {
		case layers.LayerTypeTCP:
			tcp, ok = l.(*layers.TCP)
		case layers.LayerTypeUDP:
			udp, ok = l.(*layers.UDP)
		case layers.LayerTypeIPv4:
			ipv4, ok = l.(*layers.IPv4)
		case layers.LayerTypeIPv6:
			ipv6, ok = l.(*layers.IPv6)
		case layers.LayerTypeEthernet:
			eth, ok = l.(*layers.Ethernet)
		}

		if !ok {
			errs = append(errs, errors.New("Decode error for "+lt.String()))
			// updating for the next iteration.
		}
		layerInfo[lt.String()] = true
	}

	return &info{
		layerInfo: layerInfo,
		eth:       eth,
		ipv4:      ipv4,
		ipv6:      ipv6,
		tcp:       tcp,
		udp:       udp,
		errs:      errs,
	}
}
