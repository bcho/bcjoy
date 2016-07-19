package team

import (
	"log"
	"time"
)

var (
	DEFAULT_POLLING_INTERVAL = time.Duration(2 * time.Second)
)

type Poller struct {
	Interval *time.Ticker
	Team     *Team
}

func (p Poller) Start() error {
	go func() {
		interval := p.Interval
		if interval == nil {
			interval = time.NewTicker(DEFAULT_POLLING_INTERVAL)
		}
		defer interval.Stop()

		for {
			select {
			case <-interval.C:
				p.Poll()
			}
		}
	}()

	return nil
}

func (p Poller) Poll() error {
	log.Printf("polling team")

	if err := p.Team.UpdateMembers(); err != nil {
		return err
	}

	return nil
}
