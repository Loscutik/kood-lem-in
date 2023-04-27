package room

import (
	"fmt"
)

type AntFarm struct {
	Rooms      []*Room
	Start, End *Room
}

/*
AddRoom adds a room with the given name and coordinates to the farm.
If the farm already has a room with the same name, an error will be returned.
Returns a pointer to the added room
*/
func (f *AntFarm) addRoom(name string, x, y int) (*Room, error) {
	room := Room{Name: name, Coord: coord{x, y}}

	if f.contains(&room) {
		return nil, fmt.Errorf("duplicated room")
	}

	f.Rooms = append(f.Rooms, &room)

	return &room, nil
}

/*
checks if the slice of rooms containes the given room
*/
func (f *AntFarm) contains(room *Room) bool {
	for _, r := range f.Rooms {
		if r == room {
			return true
		}
	}
	return false
}

/*
search a room with the given name in the farm. A return value of nil indicate no found.
*/
func (f *AntFarm) findRoom(name string) *Room {
	for _, r := range f.Rooms {
		if r.Name == name {
			return r
		}
	}
	return nil
}