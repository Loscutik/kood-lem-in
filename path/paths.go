package path

import (
	"fmt"

	"lemin/room"
)

type Path []*room.Room

/*
searches all not intersectons paths from start to end using breadth-first search algorithm. Returned paths are sorted in accending oreder
Start is not include into the path.
*/
func searchAllNotIntersectedPaths(farm *room.AntFarm) (allNotIntersectedPaths []*Path) {
	if farm.Start == farm.End {
		allNotIntersectedPaths = append(allNotIntersectedPaths, nil)
		return
	}
	// create list of visited rooms. At start all rooms except of the start is unvisited
	exploredRooms := New(farm)

	var queue queue

	// push all  start's link to the queue
	for _, linkedRoom := range farm.Start.Links {
		queue.pushToBack(linkedRoom)
		exploredRooms.setExplored(linkedRoom, nil)
	}
	for !queue.isEmpty() {
		currentRoom := queue.popFromFront()

		if currentRoom == farm.End {
			// create a new path and mark its rooms as part of the path
			allNotIntersectedPaths = append(allNotIntersectedPaths, exploredRooms.createPath(farm.End))
			// the others room as unexplored
			exploredRooms.switchExploredintoUnexplored()
			// create new start queue
			// if the end room is linked to start it will not be added to queue (in that it is marked as part of a path).
			// If there is a path start-end, it will have been handled at start of the loop, so mustn't be added to start queue
			queue.clear()
			for _, linkedRoom := range farm.Start.Links {
				if exploredRooms.isUnexplored(linkedRoom) {
					queue.pushToBack(linkedRoom)
					exploredRooms.setExplored(linkedRoom, nil)
				}
			}
			exploredRooms.setUnexplored(farm.End)
			continue
		}

		for _, linkedRoom := range currentRoom.Links {
			if exploredRooms.isUnexplored(linkedRoom) {
				queue.pushToBack(linkedRoom)
				exploredRooms.setExplored(linkedRoom, currentRoom)
				// fmt.Printf("curr room: %s, parrent: %s\n", linkedRoom.Name, exploredRooms[linkedRoom.Name].parent.Name)
			}
		}
	}

	return
}

func (p *Path) len() int {
	return len(*p)
}

func (p *Path) getRoom(number int) *room.Room {
	return (*p)[number]
}

func (p *Path) String() string {
	res := fmt.Sprintf("len:%d, addr: %p. strt - ", p.len(), p)

	for i := 0; i < p.len(); i++ {
		res += fmt.Sprintf("-%s- ", p.getRoom(i).Name)
	}
	return res
}
