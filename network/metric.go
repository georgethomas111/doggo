package network

import (
	"time"

	"github.com/georgethomas111/doggohttp/stats"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

// Metric is a struct which listens for a heartBeat and sends the information
// collected till now to the st
type Metric struct {
	Sc stats.Client
	Ps *gopacket.PacketSource
	h  *pcap.Handle
}

// New intitializes metric capture with an interface name and the stats client
// which receives the name.
func New(iName string, c stats.Client) (*Metric, error) {
	handle, err := pcap.OpenLive(iName, int32(65535), false, -1*time.Second)
	if err != nil {
		return nil, err
	}

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	m := &Metric{
		Sc: c,
		Ps: packetSource,
		h:  handle,
	}

	go m.PacketSource()
	return m, nil
}

func (m *Metric) PacketSource() {
	for p := range m.Ps.Packets() {
		info := packetInfo(p)
		m.Sc.Receive(info.pMap)
	}
}

func (m *Metric) Close() {
	m.h.Close()
}
