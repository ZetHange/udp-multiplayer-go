package data

import (
	"sync"

	"github.com/E4/box2d"
)

var Maps []Map
var m sync.RWMutex

func JoinUser(mapId int, user *User) {
	for _, world := range Maps {
		if world.Id == mapId {
			m.RLock()
			body := CreateBodyFromUser(world.World, user)
			user.Body = body
			world.Users = append(world.Users, user)
			m.RUnlock()

			return
		}
	}

	m.Lock()
	world := box2d.MakeB2World(box2d.MakeB2Vec2(0, 0))
	body := CreateBodyFromUser(&world, user)
	user.Body = body

	Maps = append(Maps, Map{
		Id:    mapId,
		World: &world,
		Users: []*User{user},
	})
	m.Unlock()
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

func UpdateUser(user *User) {
	// TODO
}
