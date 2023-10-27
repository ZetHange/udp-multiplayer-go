package data

import (
	"log"
	"time"
)

func B2Init() {
	log.Println("Start loop box2d")
	for {
		for _, world := range Maps.GetMaps() {
			world := world
			func() {
				world.World.Step(1.0/60.0, 6, 2)

				for _, user := range world.Users {
					pos := user.Body.GetPosition()

					Maps.Lock()
					user.X = pos.X
					user.Y = pos.Y
					Maps.Unlock()
				}
			}()
		}
		time.Sleep(20 * time.Millisecond)
	}
}
