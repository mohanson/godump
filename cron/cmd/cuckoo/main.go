package main

import (
	"log"
	"time"

	"github.com/mohanson/godump/cron"
)

func main() {
	for range cron.Cron(time.Minute, time.Second*30) {
		log.Println("main: cuckoo")
	}
}
