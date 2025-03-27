package main

import (
	"crypto/sha256"
	"log"
	"math/rand/v2"
	"runtime"
	"time"

	"github.com/mohanson/godump/gool"
)

func once() {
	buffer := make([]byte, 256)
	chacha := rand.NewChaCha8([32]byte{})
	hasher := sha256.New()
	for range 1024 {
		chacha.Read(buffer)
		hasher.Write(buffer)
		hasher.Sum(nil)
		hasher.Reset()
	}
}

// Execution on different CPUs.
//
// Intel(R) Core(TM) m3-7Y30 --- 590976
// Intel(R) Xeon(R) Gold 6133 -- 706816
// Intel(R) Core(TM) i7-9700 -- 1020672
// AMD EPYC 7K62 -------------- 1861888
func mainLoop() int {
	done := 0
	time.AfterFunc(8*time.Second, func() {
		done += 1
	})
	cnts := 0
	for done != 1 {
		once()
		cnts += 1
	}
	rate := cnts * 128
	return rate
}

func mainGool() int {
	done := 0
	time.AfterFunc(8*time.Second, func() {
		done += 1
	})
	cnts := 0
	for done != 1 {
		gool.Call(func() {
			once()
			gool.Lock(func() {
				cnts += 1
			})
		})
	}
	gool.Wait()
	rate := cnts * 128
	return rate
}

func main() {
	log.Println("main:", runtime.NumCPU(), "logical cpus usable by the current process")
	log.Println("main: sha256 by loop")
	log.Println("main: sha256 by loop rate", mainLoop())
	log.Println("main: sha256 by gool")
	log.Println("main: sha256 by gool rate", mainGool())
}
