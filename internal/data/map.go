package data

import "sync"

var Maps []Map
var m sync.RWMutex

func JoinUser(mapId int, user *User) {
	var exists = false
	for _, world := range Maps {
		if world.Id == mapId {
			exists = true

			m.Lock()
			world.Users = append(world.Users, user)
			m.Unlock()

			return
		}
	}

	if !exists {
		m.Lock()
		Maps = append(Maps, Map{
			Id:    mapId,
			Users: []*User{user},
		})
		m.Unlock()
		return
	}
}
