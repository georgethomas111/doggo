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
	Sc     stats.Client
	Ps     *gopacket.PacketSource
	pInfos []*info
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
	}

	go m.PacketSource()
	return m, nil
}

func (m *Metric) PacketSource() {
	for p := range m.Ps.Packets() {
		info := packetInfo(p)
		m.pInfos = append(m.pInfos, info)
	}
}

func (m *Metric) Trigger() {
	m.Sc.Receive(m.pInfos)
	m.Reset()
}

func (m *Metric) Reset() {
}
