package data

import "github.com/E4/box2d"

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

type Map struct {
	Id    int            `json:"id"`
	World *box2d.B2World `json:"-"`
	Users []*User        `json:"users,omitempty"`
}
