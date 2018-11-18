package service

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

type StopApp interface {
	Close()
}

type Stop struct {
	apps []StopApp
}

func (s *Stop) Add(app StopApp) {
	s.apps = append(s.apps, app)
}

func (s *Stop) Wait() {
	stop := make(chan os.Signal)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Waiting for interrupt")

	<-stop
	for _, app := range s.apps {
		app.Close()
	}

	fmt.Println("Received interrupt. Bye use me again.")
}
