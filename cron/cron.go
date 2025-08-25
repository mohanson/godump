// Package cron provides a simple way to schedule tasks at fixed intervals.
package cron

import (
	"time"

	"github.com/libraries/go/doa"
)

// Cron schedules tasks at regular intervals.
func Cron(e time.Duration, d time.Duration) <-chan struct{} {
	// Ensure that the delay is less than or equal to the elapsed time between events.
	doa.Doa(d < e)
	r := make(chan struct{})
	go func() {
		for {
			n := time.Now()
			// Wait for the next scheduled event by adding the elapsed time and then waiting for the delay.
			time.Sleep(n.Add(e).Truncate(e).Sub(n) + d)
			r <- struct{}{}
		}
	}()
	return r
}
