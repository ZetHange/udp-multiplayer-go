package data

type User struct {
	Id     string  `json:"id"`
	Login  string  `json:"login"`
	Health int     `json:"health"`
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Dx     float64 `json:"dx"`
	Dy     float64 `json:"dy"`
}

type Map struct {
	Id    int     `json:"id"`
	Users []*User `json:"users,omitempty"`
}
