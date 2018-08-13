package network

import (
	"sync"
	"time"

	"github.com/georgethomas111/doggohttp/stats"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

// Metric is a struct which listens for a heartBeat and sends the information
// collected till now to the st
type Metric struct {
	L  sync.Mutex
	Sc stats.Client
	Ps *gopacket.PacketSource
}

// New intitializes metric capture with an interface name and the stats client
// which receives the name.
func New(iName string, c stats.Client) (*Metric, error) {
	handle, err := pcap.OpenLive(iName, int32(65535), false, -1*time.Second)
	if err != nil {
		return nil, err
	}

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	return &Metric{
		Sc: c,
		Ps: packetSource,
	}, nil
}

func (m *Metric) PacketSource() {
	for p := range m.Ps.Packets() {
		packetInfo(p)
	}
}

func (m *Metric) Trigger() {
	m.L.Lock()
	m.Sc.Receive(m)
	m.Reset()

	m.L.Unlock()
}

func (m *Metric) Reset() {
}
