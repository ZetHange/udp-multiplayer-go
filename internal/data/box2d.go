package data

import (
	"log"
	"time"
)

func B2Init() {
	log.Println("Start loop box2d")
	for {
		for _, plan := range Maps {
			plan := plan
			go func() {
				plan.World.Step(1.0/60.0, 6, 2)

				for _, user := range plan.Users {
					pos := user.Body.GetPosition()

					m.Lock()
					user.X = pos.X
					user.Y = pos.Y
					m.Unlock()
				}
			}()
		}
		time.Sleep(20 * time.Millisecond)
	}
}
