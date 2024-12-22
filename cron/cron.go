// Package cron provides functionality for generate signals at intervals timed in January 1, 1970 UTC.
package cron

import (
	"time"
)

// Cron loop generate a signal every d duration timed in January 1, 1970 UTC.
func Cron(d time.Duration) <-chan struct{} {
	r := make(chan struct{})
	go func() {
		for {
			n := time.Now().UnixNano()
			m := time.Duration(int64(d)-n%int64(d)) * time.Nanosecond
			time.Sleep(m)
			r <- struct{}{}
		}
	}()
	return r
}

// Wait function offsets the received time signal by d duration.
func Wait(c <-chan struct{}, d time.Duration) <-chan struct{} {
	r := make(chan struct{})
	go func() {
		for range c {
			time.Sleep(d)
			r <- struct{}{}
		}
	}()
	return r
}
