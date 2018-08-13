package network

import (
	"fmt"
	"testing"
	"time"
)

type testStatsClient struct {
}

func (t *testStatsClient) Receive(d map[string]int) {
	for key, count := range d {
		fmt.Println(key, count)
	}
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
