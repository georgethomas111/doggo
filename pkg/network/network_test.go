package network

import (
	"testing"

	"github.com/georgethomas111/doggo/pkg/stats"
)

type testStatsClient struct {
}

func (t *testStatsClient) Receive(s stats.Packet) {
}

func TestNetwork(t *testing.T) {
	// create a stats client that can be used.
	_, err := New("", []stats.Client{&testStatsClient{}})
	if err != nil {
		return
	}
	t.Errorf("Expected the test to return an error as the correct interface is not passed.")
}

func TestNetworkWithInterface(t *testing.T) {
	// create a stats client that can be used.
	interfaces, err := LS()
	if err != nil {
		t.Error("Error listing interfaces ", err.Error())
	}

	if len(interfaces) == 0 {
		t.Errorf("At least one interface required for testing.")
	}

	_, err = New(interfaces[0], []stats.Client{&testStatsClient{}})
	if err != nil {
		t.Errorf("Could not listen to %v, got error %v", interfaces[0], err.Error())
	}
}
