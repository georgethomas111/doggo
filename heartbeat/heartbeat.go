package heartbeat

import "time"

type HeartBeat struct {
	Duration time.Duration
	apps     []Application
}

// New creates a new heartBeat given a duration and a set of apps.
func New(d time.Duration, apps []Application) *HeartBeat {
	h := &HeartBeat{
		Duration: d,
		apps:     apps,
	}

	go h.beat()
	return h
}

func (h *HeartBeat) beat() {
	for {
		t := time.NewTimer(h.Duration)
		<-t.C
		for _, app := range h.apps {
			app.Trigger()
		}
	}

}
