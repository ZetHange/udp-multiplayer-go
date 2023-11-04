package app

import (
	"log"
	"sync"
	"time"
	"udp-multiplayer-go/internal/data"
)

type Counter struct {
	sync.Mutex
	c        int64
	requests chan int
}

var counter = Counter{
	requests: make(chan int),
}

type metrics struct {
	sync.Mutex
	Rps        int64 `json:"rps"`
	TotalUsers int64 `json:"total_users"`
}

var Metrics metrics

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
			Metrics.Rps = counter.c

			data.UserList.Lock()
			Metrics.TotalUsers = int64(len(data.UserList.Users))
			data.UserList.Unlock()

			counter.c = 0
			counter.Unlock()

			time.Sleep(time.Second)
		}
	}()
}

func HandleRequest() {
	counter.requests <- 1
}
