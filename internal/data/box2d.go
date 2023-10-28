package data

import (
	"log"
	"time"

	"github.com/E4/box2d"
)

func B2Init() {
	log.Println("Start loop box2d")
	ticker := time.NewTicker(time.Millisecond * 20) // 20 миллисекунд = 1/60 секунды
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			for _, world := range MapList.GetMaps() {
				world := world
				func() {
					world.World.Step(1.0/60.0, 6, 2)

					for _, user := range world.Users {
						pos := user.Body.GetPosition()

						MapList.Lock()
						user.X = pos.X
						user.Y = pos.Y
						MapList.Unlock()
					}
				}()
			}
		}
	}
}

func CreateBodyFromUser(world *box2d.B2World, user *User) *box2d.B2Body {
	bodyDef := box2d.MakeB2BodyDef()
	bodyDef.Type = box2d.B2BodyType.B2_dynamicBody
	bodyDef.Position.Set(user.X, user.Y)

	body := world.CreateBody(&bodyDef)
	body.SetLinearDamping(10)
	circle := box2d.NewB2CircleShape()
	circle.SetRadius(3)

	fixtureDef := box2d.MakeB2FixtureDef()
	fixtureDef.Shape = circle
	fixtureDef.Density = 0.5
	fixtureDef.Friction = 0

	body.CreateFixture(fixtureDef.Shape, fixtureDef.Density)

	return body
}
