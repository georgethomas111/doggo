package network

import (
	"time"

	"github.com/georgethomas111/doggo/pkg/stats"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

// Network registers a client which takes appropriate action when a packet is
// received from it.
type Network struct {
	Sc []stats.Client
	Ps *gopacket.PacketSource
	h  *pcap.Handle
}

// New intitializes metric capture with an interface name and the stats client
// which receives the name.
func New(iName string, clients []stats.Client) (*Network, error) {
	handle, err := pcap.OpenLive(iName, int32(65535), false, -1*time.Second)
	if err != nil {
		return nil, err
	}

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	n := &Network{
		Sc: clients,
		Ps: packetSource,
		h:  handle,
	}

	go n.PacketSource()
	return n, nil
}

func (n *Network) PacketSource() {
	for p := range n.Ps.Packets() {
		info := packetInfo(p)
		for _, c := range n.Sc {
			c.Receive(info)
		}
	}
}

func (n *Network) Close() {
	n.h.Close()
}
