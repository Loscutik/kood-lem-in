package path

import (
	"fmt"

	"lemin/room"
)

type Path []*room.Room

type pathHandler struct {
	flowCheck        func(*room.Link) bool
	setFlow          func(*room.Room, *room.Room)
	roomOnPathStatus byte
}

func isAllowedToGo(l *room.Link) bool {
	return l.Flow != 1
}

func isPartOfPaths(l *room.Link) bool {
	return l.Flow == 1
}

func isJustBFS(l *room.Link) bool {
	return l.Flow == 0
}

func noFlows(from *room.Room, to *room.Room) {}

func changeFlow(from *room.Room, to *room.Room) {
	link := from.GetLinkTo(to)
	if link == nil {
		panic("incorrect link")
	}
	link.Flow++

	link = to.GetLinkTo(from)
	if link == nil {
		panic("incorrect link")
	}
	link.Flow--
}

func resetFlow(from *room.Room, to *room.Room) {
	link := from.GetLinkTo(to)
	if link == nil {
		panic("incorrect link")
	}
	link.Flow = 0

	link = to.GetLinkTo(from)
	if link == nil {
		panic("incorrect link")
	}
	link.Flow = 0
}

/*
searches all not intersectons paths from start to end using breadth-first search algorithm. Returned paths are sorted in accending oreder
Start is not include into the path.
*/
func SearchAllPaths(farm *room.AntFarm, pathHandler *pathHandler) (paths []*Path) {
	if farm.Start == farm.End {
		paths = append(paths, nil)
		return
	}
	// create list of visited rooms. At start all rooms except of the start is unvisited
	exploredRooms := New(farm)

	var queue queue

	// push all  start's link to the queue
	for _, link := range farm.Start.Links {
		queue.pushToBack(link.Room)
		exploredRooms.setExplored(link.Room, nil)
	}
	for !queue.isEmpty() {
		currentRoom := queue.popFromFront()

		if currentRoom == farm.End {
			// create a new path and mark its rooms as part of the path
			paths = append(paths, exploredRooms.createPath(farm.Start, farm.End, pathHandler))
			// the others room as unexplored
			exploredRooms.switchExploredintoUnexplored()
			// create a new start queue
			// if the end room is linked to the start it will not be added to the queue (in that it is marked as a part of the path).
			// If there is a path start-end, it will have been handled at the start of the loop, so it mustn't be added to start queue
			queue.clear()
			for _, link := range farm.Start.Links {
				if exploredRooms.isUnexplored(link.Room) && pathHandler.flowCheck(link) { 
					queue.pushToBack(link.Room)
					exploredRooms.setExplored(link.Room, nil)
				}
			}
			exploredRooms.setUnexplored(farm.End)
			continue
		}

		for _, link := range currentRoom.Links {
			if exploredRooms.isUnexplored(link.Room) && pathHandler.flowCheck(link) {
				queue.pushToBack(link.Room)
				exploredRooms.setExplored(link.Room, currentRoom)
				// fmt.Printf("curr room: %s, parrent: %s\n", link.Room.Name, exploredRooms[link.Room.Name].parent.Name)
			}
		}
	}

	return
}

func SearchAllNotIntersectedPaths(farm *room.AntFarm) []*Path {
	SearchAllPaths(farm, &pathHandler{isAllowedToGo, changeFlow, UNEXPLORED})
	return SearchAllPaths(farm, &pathHandler{isPartOfPaths, resetFlow, ON_PATH})
}

func (p *Path) Len() int {
	return len(*p)
}

func (p *Path) GetRoom(number int) *room.Room {
	return (*p)[number]
}

func (p *Path) String() string {
	res := fmt.Sprintf("len:%d, addr: %p. strt - ", p.Len(), p)

	for i := 0; i < p.Len(); i++ {
		res += fmt.Sprintf("-%s- ", p.GetRoom(i).Name)
	}
	return res
}
