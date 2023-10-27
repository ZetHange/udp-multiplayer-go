package data

import (
	"sync"

	"github.com/E4/box2d"
)

type MapsType struct {
	sync.RWMutex
	Maps []*Map
}

var Maps MapsType

func (m *MapsType) GetMaps() []*Map {
	m.RLock()
	defer m.RUnlock()

	return m.Maps
}

func (m *MapsType) GetMapById(mapId int) (*Map, bool) {
	m.RLock()
	defer m.RUnlock()

	for _, world := range m.Maps {
		if world.Id == mapId {
			return world, true
		}
	}
	return nil, false
}

func (m *MapsType) JoinUser(mapId int, user *User) {
	gettedMap, ok := m.GetMapById(mapId)

	if ok {
		body := CreateBodyFromUser(gettedMap.World, user)
		user.Body = body
		gettedMap.Users = append(gettedMap.Users, user)
	} else {
		world := box2d.MakeB2World(box2d.MakeB2Vec2(0, 0))
		body := CreateBodyFromUser(&world, user)
		user.Body = body

		Maps.Maps = append(Maps.Maps, &Map{
			Id:    mapId,
			World: &world,
			Users: []*User{user},
		})
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

func UpdateUser(user *User) {
	// TODO
}
