package data

import (
	"github.com/E4/box2d"
	"log"
	"sync"
	"time"
)

func B2Init() {
	for {
		wg := sync.WaitGroup{}
		for _, mapa := range Maps {
			go func() {
				wg.Add(1)

				world := box2d.MakeB2World(box2d.B2Vec2{X: 0, Y: 0})

				for _, user := range mapa.Users {
					bodyDef := box2d.MakeB2BodyDef()
					bodyDef.Position.Set(user.X, user.Y)
					body := world.CreateBody(&bodyDef)
					body.SetLinearVelocity(box2d.B2Vec2{X: user.Dx, Y: user.Dy})

					timeStep := 1.0 / 60.0
					velocityIterations := 6
					positionIterations := 2

					world.Step(timeStep, velocityIterations, positionIterations)
					position := body.GetPosition()

					m.Lock()
					user.X = position.X
					user.Y = position.Y
					m.Unlock()
					log.Println("update position,", position)
				}
				wg.Done()
			}()
		}
		wg.Wait()
		time.Sleep(20 * time.Millisecond)
	}
}
