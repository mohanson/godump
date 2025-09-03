// Package deadline provides a mechanism to execute a function after a specified time, with the ability to reset or
// stop the timer.
package deadline

import (
	"sync"
	"time"
)

// Deadline represents a timer that will execute a function after a certain time.
type Deadline struct {
	call func()     // Function to call when the deadline is reached
	done int        // Flag indicating if the deadline has been completed (1) or not (0)
	mu   sync.Mutex // Mutex to protect concurrent access
	time time.Time  // The time at which the function should be executed
}

// Reset sets a new deadline time. If the deadline was already done, it restarts the timer.
func (d *Deadline) Reset(t time.Time) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.time = t
	if d.done == 1 {
		d.done = 0
		d.wait()
	}
}

// Stop marks the deadline as done, preventing the function from being called.
func (d *Deadline) Stop() {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.done = 1
}

// Wait checks every second if the deadline has been reached. If so, it calls the function.
func (d *Deadline) wait() {
	time.AfterFunc(time.Second, func() {
		d.mu.Lock()
		defer d.mu.Unlock()
		if d.done == 1 {
			// Deadline stopped, do not continue waiting.
			return
		}
		if time.Now().After(d.time) {
			// Deadline reached, call the function.
			d.call()
			d.done = 1
			return
		}
		d.wait()
	})
}

// New creates a new Deadline that will call the provided function f after time d.
func New(d time.Time, f func()) *Deadline {
	deadline := &Deadline{
		call: f,
		time: d,
	}
	deadline.wait()
	return deadline
}
