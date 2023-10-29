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

func MetricsInit() {
	log.Println("Start metrics")

	go func() {
		for {
			select {
			case <-time.After(time.Second):
				counter.Lock()
				counter.c = 0
				counter.Unlock()
			case <-counter.requests:
				counter.Lock()
				counter.c += 1
				counter.Unlock()
			}
		}
	}()

}

func HandleRequest() {
	counter.requests <- 1
}
