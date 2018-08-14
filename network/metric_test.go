package network

import (
	"fmt"
	"testing"
	"time"

	"github.com/georgethomas111/doggohttp/stats"
)

type testStatsClient struct {
}

func (t *testStatsClient) Receive(s stats.PacketStats) {
	for key, _ := range s.LayerInfo() {
		fmt.Println(key)
	}
	fmt.Println()
}

func TestMetric(t *testing.T) {
	// create a stats client that can be used.
	m, err := New("wlan0", &testStatsClient{})
	if err != nil {
		t.Errorf("error while creating metric %s", err.Error())
		return
	}

	time.Sleep(10 * time.Second)
	m.Close()
}
