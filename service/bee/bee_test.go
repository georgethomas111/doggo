package bee

import (
	"testing"

	"github.com/google/gopacket/layers"
)

type MockDB struct {
	Written bool
}

// Write satisfies the interface
func (m *MockDB) Write(interface{}) {
	m.Written = true
}

type MockPacket struct {
}

func (m *MockPacket) LayerInfo() map[string]bool {
	return nil
}

func (m *MockPacket) Eth() *layers.Ethernet {
	return nil
}

func (m *MockPacket) IPv4() *layers.IPv4 {
	return nil
}

func (m *MockPacket) IPv6() *layers.IPv6 {
	return nil
}

func (m *MockPacket) TCP() *layers.TCP {
	return nil
}

func (m *MockPacket) UDP() *layers.UDP {
	return nil
}

func TestBee(t *testing.T) {
	db := new(MockDB)
	b := New(db)
	b.Receive(new(MockPacket))
	b.Trigger()

	if !db.Written {
		t.Errorf("Expected data to be written to the db.")
	}
}
