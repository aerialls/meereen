package core

import (
	"time"
)

// Scheduler is responsible to launch checks when needed
type Scheduler struct {
	container *Container
}

// NewScheduler returns a new scheduler
func NewScheduler(
	container *Container,
) *Scheduler {
	return &Scheduler{
		container: container,
	}
}

// Start the scheduler
func (s *Scheduler) Start() chan bool {
	stopped := make(chan bool, 1)
	ticker := time.NewTicker(1 * time.Minute)

	go func() {
		for {
			select {
			case <-ticker.C:
				s.run()
			case <-stopped:
				ticker.Stop()
				return
			}
		}
	}()

	return stopped
}

func (s *Scheduler) run() {
	for _, check := range s.container.GetChecks() {
		check.Run()
	}
}
