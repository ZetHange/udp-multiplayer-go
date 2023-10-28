package data

import (
	"log"
	"time"

	"github.com/E4/box2d"
)

func B2Init() {
	log.Println("Start loop box2d")
	for {
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
		time.Sleep(20 * time.Millisecond)
	}
}

func CreateBodyFromUser(world *box2d.B2World, user *User) *box2d.B2Body {
	bodyDef := box2d.MakeB2BodyDef()
	bodyDef.Type = box2d.B2BodyType.B2_dynamicBody
	bodyDef.Position.Set(user.X, user.Y)

	dynamicBox := box2d.B2PolygonShape{}
	dynamicBox.SetAsBox(1.0, 1.0)

	fixtureDef := box2d.MakeB2FixtureDef()
	fixtureDef.Shape = &dynamicBox
	fixtureDef.Density = 1.0
	fixtureDef.Friction = 0.3

	body := world.CreateBody(&bodyDef)
	body.CreateFixture(fixtureDef.Shape, fixtureDef.Density)
	body.SetUserData(user.Id)
	body.SetLinearVelocity(box2d.B2Vec2{X: user.Dx, Y: user.Dy})

	return body
}
