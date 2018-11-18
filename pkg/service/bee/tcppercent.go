package bee

import (
	"sync"

	"github.com/google/gopacket/layers"
)

type TCPPercent struct {
	sync.Mutex
	ackCount   int     `json:"-"`
	packCount  int     `json:"-"`
	Percentage float64 `json:"percentage"`
}

func (t *TCPPercent) Update(tcp *layers.TCP) {
	if tcp == nil {
		return
	}

	t.Lock()

	t.packCount++
	if tcp.ACK {
		t.ackCount++
	}
	t.Percentage = (float64(t.ackCount) / float64(t.packCount)) * 100

	t.Unlock()
}

func (t *TCPPercent) Clear() {
	t.Lock()
	t.ackCount = 0
	t.packCount = 0
	t.Unlock()
}
