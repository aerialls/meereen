package core

import (
	"time"

	log "github.com/sirupsen/logrus"

	c "github.com/aerialls/meereen/pkg/check"
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
	log.Debugf("scheduling all checks to run")
	for _, check := range s.container.GetChecks() {
		if check.IsRunning() {
			log.WithField("title", check.GetTitle()).Warn("check is already running, skipping this run")
			continue
		}

		go func(check c.Check) {
			check.Run()
		}(check)
	}
}
