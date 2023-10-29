package utils

import (
	"log"
	"sync"
	"time"
	"udp-multiplayer-go/internal/data"
)

type oko struct {
	sync.Mutex
	Users map[string]time.Time
}

var Oko oko = oko{
	Mutex: sync.Mutex{},
	Users: map[string]time.Time{},
}

func InitOko() {
	log.Println("Oko initialize")

	for {
		Oko.Lock()
		for key, value := range Oko.Users {
			duration := time.Since(value)
			if duration >= time.Second*5 {
				user, ok := data.Leave(key)
				if ok {
					log.Printf("[AUTODISCONNECT](id: %s) User with login: %s autodisconnected from map", user.Id, user.Login)
				}
			}
		}
		Oko.Unlock()

		time.Sleep(1 * time.Second)
	}
}
