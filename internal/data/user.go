package data

import (
	"sync"

	"github.com/E4/box2d"
)

type User struct {
	Id     string        `json:"id"`
	Login  string        `json:"login"`
	Health int           `json:"health"`
	Body   *box2d.B2Body `json:"-"`
	X      float64       `json:"x"`
	Y      float64       `json:"y"`
	Dx     float64       `json:"dx"`
	Dy     float64       `json:"dy"`
}

type UserListType struct {
	sync.RWMutex
	Users []*User
}

var UserList UserListType

func (u *UserListType) GetUserByUUID(uuid string) (*User, bool) {
	u.Lock()
	defer u.Unlock()

	for _, user := range u.Users {
		if user.Id == uuid {
			return user, true
		}
	}
	return nil, false
}

func UpdateUser(uuid string, dx, dy float64) bool {
	user, ok := UserList.GetUserByUUID(uuid)
	if !ok {
		return false
	}

	UserList.Lock()

	user.Body.ApplyLinearImpulse(box2d.B2Vec2{
		X: user.Dx * 30.0,
		Y: user.Dy * 30.0,
	}, box2d.B2Vec2{
		X: user.X,
		Y: user.Y,
	}, true)

	user.Dx = dx
	user.Dy = dy

	UserList.Unlock()
	return true
}

func Leave(uuid string) (*User, bool) {
	user, ok := UserList.GetUserByUUID(uuid)
	if !ok {
		return nil, ok
	}

	for _, world := range MapList.GetMaps() {
		for i, user := range world.Users {
			if user.Id == uuid {
				world.World.DestroyBody(user.Body)
				world.Users = append(world.Users[:i], world.Users[i+1:]...)
				break
			}
		}
	}

	for i, user := range UserList.Users {
		if user.Id == uuid {
			UserList.Users = append(UserList.Users[:i], UserList.Users[i+1:]...)
			break
		}
	}

	return user, ok
}
