package heartbeat

import (
	"testing"
	"time"
)

type TApp struct {
	TestData string
}

func (t *TApp) Trigger() {
	t.TestData = "abcd"
}

func TestHeartBeat(t *testing.T) {
	tApp := new(TApp)

	New(time.Millisecond, []Application{tApp})
	ti := time.NewTimer(2 * time.Millisecond)
	<-ti.C
	if tApp.TestData != "abcd" {
		t.Errorf("Trigger not called in %v.", time.Millisecond)
	}
}
