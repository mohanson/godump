package balloc

import (
	"math/rand/v2"
	"testing"

	"github.com/mohanson/godump/doa"
)

func TestBalloc(t *testing.T) {
	maxAlive := 1024
	maxAlloc := 1024
	maxRange := 1024 * 1024
	maxTotal := 1024 * 1024
	minBlock := 64
	b := New(maxTotal, minBlock)
	g := []Block{}
	for range maxRange {
		i := rand.Int() % maxAlive
		if i > len(g) {
			g = append(g, b.Malloc(max(1, rand.Int()%maxAlloc)))
		}
		if i < len(g) {
			b.Free(g[i])
			g = append(g[:i], g[i+1:]...)
		}
	}
	for _, e := range g {
		b.Free(e)
	}
	for _, e := range b.freeList[:len(b.freeList)-1] {
		doa.Doa(e.Len() == 0)
	}
	doa.Doa(b.freeList[len(b.freeList)-1].Len() == 1)
}
