// Package gool offers a high-level API for running tasks asynchronously, restricting concurrent executions to the
// number of cpu cores or a specified limit.
package gool

import (
	"runtime"
	"sync"
)

var cCpu = make(chan struct{}, runtime.NumCPU())

// Gool manages a pool of goroutines for asynchronous task execution.
type Gool struct {
	C chan struct{}
	M *sync.Mutex
	W *sync.WaitGroup
}

// Call submits a function f for asynchronous execution in a new goroutine, respecting the concurrency limit.
func (g *Gool) Call(f func()) {
	g.C <- struct{}{}
	g.W.Add(1)
	go func() {
		f()
		g.W.Done()
		<-g.C
	}()
}

// Lock executes function f with exclusive access, synchronizing via the mutex, typically for aggregating results.
func (g *Gool) Lock(f func()) {
	g.M.Lock()
	defer g.M.Unlock()
	f()
}

// Wait blocks until all submitted tasks have completed.
func (g *Gool) Wait() {
	g.W.Wait()
}

// Cpu initializes a Gool instance with a global concurrency limit specified by cpu cores.
func Cpu() *Gool {
	return &Gool{
		C: cCpu,
		M: &sync.Mutex{},
		W: &sync.WaitGroup{},
	}
}

// New initializes a Gool instance with a custom concurrency limit specified by n.
func New(n int) *Gool {
	return &Gool{
		C: make(chan struct{}, n),
		M: &sync.Mutex{},
		W: &sync.WaitGroup{},
	}
}
