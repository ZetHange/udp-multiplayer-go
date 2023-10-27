package data

import (
	"github.com/E4/box2d"
	"sync"
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

func (u *UserListType) GetUserByUUID(uuid string) (*User, bool) {
	u.RLock()
	defer u.RUnlock()

	for _, user := range u.Users {
		if user.Id == uuid {
			return user, true
		}
	}
	return nil, false
}
