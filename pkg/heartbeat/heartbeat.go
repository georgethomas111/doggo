package heartbeat

import "time"

type HeartBeat struct {
	Duration time.Duration
	apps     []Application
	stop     chan bool
}

// New creates a new heartBeat given a duration and a set of apps.
func New(d time.Duration, apps []Application) *HeartBeat {
	stop := make(chan bool)
	h := &HeartBeat{
		Duration: d,
		apps:     apps,
		stop:     stop,
	}

	go h.beat()
	return h
}

func (h *HeartBeat) beat() {
	for {
		t := time.NewTimer(h.Duration)
		select {
		case <-t.C:
			for _, app := range h.apps {
				app.Trigger()
			}
		case <-h.stop:
			break
		}
	}

}

func (h *HeartBeat) Close() {
	h.stop <- true
}
