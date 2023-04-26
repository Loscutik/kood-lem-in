package path

import "lemin/room"

/*
keeps rooms exploring information during broad first searching: is the room explored,
if yes to which room(parent) it was connected
*/
type exploredRoom struct {
	room   *room.Room
	parent *room.Room
	label  byte // 0 - unexplored, 1 -  explored in the current phase, 2 - is a part of already found paths
}

const(
	UNEXPLORED=iota
	EXPLORED
	ON_PATH	
)

type exploredRooms map[*room.Room]*exploredRoom

/*
creates a new maps with all rooms, marks all the rooms as unexplored, apart from the start which is marked as part of a path
*/
func New(farm *room.AntFarm) exploredRooms {
	ers := make(exploredRooms)
	for _, room := range farm.Rooms {
		ers[room] = &exploredRoom{room: room, label: UNEXPLORED}
	}

	ers[farm.Start].label = ON_PATH
	return ers
}

/*
returns true if the given room was marked as unexplored
*/
func (ers exploredRooms) isUnexplored(r *room.Room) bool {
	return (ers)[r].label == UNEXPLORED
}

/*
sets status of exploring for the given room
*/
func (ers exploredRooms) setStatus(r *room.Room, stat byte) {
	(ers)[r].label = stat
}

/*
marks the given room as unexplored
*/
func (ers exploredRooms) setUnexplored(r *room.Room) {
	ers.setStatus(r, UNEXPLORED)
	(ers)[r].parent = nil
}

/*
marks the given room as explored
*/
func (ers exploredRooms) setExplored(r *room.Room, parent *room.Room) {
	ers.setStatus(r, EXPLORED)
	(ers)[r].parent = parent
}

/*
marks the given room as placed in a path
*/
func (ers exploredRooms) setStatusInPath(r *room.Room) {
	ers.setStatus(r, ON_PATH)
}

/*
returns a parent of the given room
*/
func (ers exploredRooms) getParent(r *room.Room) *room.Room {
	return (ers)[r].parent
}

/*
creates path from a given room
*/
func (ers exploredRooms) createPath(start, end *room.Room, pathHandler *pathHandler) *Path {
	var temPath queue
	len := 0
	// push room to the queue and then its parent, then grandparent e.c.
	for end != nil {
		temPath.pushToFront(end)
		len++
		ers.setStatus(end,pathHandler.roomOnPathStatus)
		end = ers.getParent(end)
	}
	path := make(Path, len)
	path[0] = temPath.popFromFront()
	pathHandler.setFlow(start, path[0])
	for i := 1; i < len; i++ {
		path[i] = temPath.popFromFront()
		pathHandler.setFlow(path[i-1], path[i])
	}

	return &path
}

/*
switches status for all explored but not placed in a path rooms to unexplored status
*/
func (ers exploredRooms) switchExploredintoUnexplored() {
	for _, er := range ers {
		if er.label == EXPLORED {
			er.label = UNEXPLORED
		}
	}
}
