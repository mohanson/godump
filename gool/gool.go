// Package gool provides a high-level interface for asynchronously executing callables. Gool limits the number of
// concurrently executed tasks to be less than or equal to the number of CPU cores.
package gool

import (
	"runtime"
	"sync"
)

var c = func() chan struct{} {
	n := runtime.NumCPU()
	c := make(chan struct{}, n)
	return c
}()

// Gool is an executor that uses a pool of goroutines to execute calls asynchronously.
type Gool struct {
	W *sync.WaitGroup
	M *sync.Mutex
}

// Call the callable f.
func (g *Gool) Call(f func()) {
	c <- struct{}{}
	g.W.Add(1)
	go func() {
		f()
		g.W.Done()
		<-c
	}()
}

// Lock locks m and call f. This function is usually used to aggregate calculation results.
func (g *Gool) Lock(f func()) {
	g.M.Lock()
	defer g.M.Unlock()
	f()
}

// Wait blocks until all tasks is done.
func (g *Gool) Wait() {
	g.W.Wait()
}

// Init creates a new Gool.
func Init() *Gool {
	return &Gool{
		W: &sync.WaitGroup{},
		M: &sync.Mutex{},
	}
}
