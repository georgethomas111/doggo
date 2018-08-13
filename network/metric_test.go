package network

import (
	"log"
	"testing"
	"time"
)

type testStatsClient struct {
}

func (t *testStatsClient) Receive(d interface{}) {
	log.Println(d)
}

func TestMetric(t *testing.T) {
	// create a stats client that can be used.
	m, err := New("wlan0", &testStatsClient{})
	if err != nil {
		t.Errorf("error while creating metric %s", err.Error())
		return
	}
	time.Sleep(time.Second)
	m.Trigger()
}
