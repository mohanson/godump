// Package gool provides a high-level interface for asynchronously executing callables. Gool limits the number of
// concurrently executed tasks to be less than or equal to the number of CPU cores.
package gool

import (
	"runtime"
	"sync"
)

var (
	defaultChan = func() chan struct{} {
		n := runtime.NumCPU()
		c := make(chan struct{}, n)
		return c
	}()
	defaultGool = NewGool()
)

// Gool is an executor that uses a pool of goroutines to execute calls asynchronously.
type Gool struct {
	W *sync.WaitGroup
	M *sync.Mutex
}

// Call the callable f.
func (g *Gool) Call(f func()) {
	defaultChan <- struct{}{}
	g.W.Add(1)
	go func() {
		f()
		g.W.Done()
		<-defaultChan
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

// NewGool creates a new Gool.
func NewGool() *Gool {
	return &Gool{
		W: &sync.WaitGroup{},
		M: &sync.Mutex{},
	}
}

// Call the callable f.
func Call(f func()) {
	defaultGool.Call(f)
}

// Lock locks m and call f. This function is usually used to aggregate calculation results.
func Lock(f func()) {
	defaultGool.Lock(f)
}

// Wait blocks until all tasks is done.
func Wait() {
	defaultGool.Wait()
}
