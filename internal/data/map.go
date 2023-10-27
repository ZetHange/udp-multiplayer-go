package data

import (
	"sync"
	"udp-multiplayer-go/proto/pb"

	"github.com/E4/box2d"
)

type Map struct {
	Id    int            `json:"id"`
	World *box2d.B2World `json:"-"`
	Users []*User        `json:"users,omitempty"`
}

type MapsType struct {
	sync.RWMutex
	MapList []*Map
}

var MapList MapsType
var UserList UserListType

func (m *MapsType) GetMaps() []*Map {
	m.RLock()
	defer m.RUnlock()

	return m.MapList
}

func (m *MapsType) GetMapById(mapId int) (*Map, bool) {
	m.RLock()
	defer m.RUnlock()

	for _, world := range m.MapList {
		if world.Id == mapId {
			return world, true
		}
	}
	return nil, false
}

func (m *MapsType) ToProto(mapId int) []*pb.User {
	var users []*pb.User
	world, _ := m.GetMapById(mapId)
	for _, user := range world.Users {
		users = append(users, &pb.User{
			Login:  user.Login,
			Health: int64(user.Health),
			X:      user.X,
			Y:      user.Y,
			Dx:     user.Dx,
			Dy:     user.Dy,
		})
	}

	return users
}

func (m *MapsType) JoinUser(mapId int, user *User) {
	gettedMap, ok := m.GetMapById(mapId)

	if ok {
		body := CreateBodyFromUser(gettedMap.World, user)
		user.Body = body

		UserList.Users = append(UserList.Users, user)
		gettedMap.Users = append(gettedMap.Users, user)
	} else {
		world := box2d.MakeB2World(box2d.MakeB2Vec2(0, 0))
		body := CreateBodyFromUser(&world, user)
		user.Body = body

		UserList.Users = append(UserList.Users, user)
		MapList.MapList = append(MapList.MapList, &Map{
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

func (m *MapsType) GetMapIdByUserId(uuid string) int {
	for _, world := range m.MapList {
		for _, user := range world.Users {
			if user.Id == uuid {
				return world.Id
			}
		}
	}
	return 0
}

func (m *MapsType) UpdateUser(uuid string, x, y, dx, dy float64) {
	user, _ := UserList.GetUserByUUID(uuid)

	UserList.Lock()
	user.Body.SetLinearVelocity(*box2d.NewB2Vec2(dx, dy))
	user.Body.SetTransform(box2d.B2Vec2{
		X: x,
		Y: y,
	}, user.Body.GetAngle())
	user.X = x
	user.Y = y
	user.Dx = dx
	user.Dy = dy
	UserList.Unlock()
}
