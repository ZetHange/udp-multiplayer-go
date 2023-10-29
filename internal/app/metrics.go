package app

import (
	"log"
	"sync"
	"time"
)

type Counter struct {
	sync.Mutex
	c        int64
	requests chan int
}

var counter = Counter{
	requests: make(chan int),
}
var LastRPS int64

func MetricsInit() {
	log.Println("Start metrics")
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	go func() {
		for {
			select {
			case <-counter.requests:
				counter.Lock()
				counter.c += 1
				counter.Unlock()
			}
		}
	}()

	go func() {
		for {
			counter.Lock()
			LastRPS = counter.c
			counter.c = 0
			counter.Unlock()

			time.Sleep(time.Second)
		}
	}()
}

func HandleRequest() {
	counter.requests <- 1
}
