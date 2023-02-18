package path

import (
	"fmt"
	"math"

	"lemin/path/room"
)

type path []*Room.Room

type queuesNode struct {
	room *Room
	next *queuesNode
}

type queue struct {
	head, tail *queuesNode
}

type exploredRoom struct {
	room   *Room
	parent *Room
	label  byte // 0 - unexplored, 1 -  explored in the current phase, 2 - is a part of already found paths
}

type exploredRooms map[string]*exploredRoom

/*
searches all not intersectons paths from start to end using breadth-first search algorithm. Returned paths are sorted in accending oreder
Start is not include into the path.
*/
func searchAllNotIntersectedPaths(farm *room.AntFarm) (allNotIntersectedPaths []path) {
	exploredRooms := New(farm)
	var queue queue
	queue.add(farm.Start)
	for !queue.isEmpty() {
		for _, nextByStartRoom := range farm.Start.Links {
			// mark the room as explored
			if !exploredRooms.isExplored(nextByStartRoom) {
				exploredRooms.setStatusExplored(nextByStartRoom, farm.Start)
				queue.add(nextByStartRoom)
			}
		}
	}
	shortestPath := make([]path, len(start.links))
	countStatrtsLinks := 0

	for _, roomNextToStart := range start.links {
		shortestPath[countStatrtsLinks] = path{len: math.MaxInt}
		var searchPath func(i int, currentPath path, currentRoom *Room)

		searchPath = func(i int, currentPath path, currentRoom *Room) {
			defer wg.Done()
			// it doesn't need to compare all path.
			// If the current path's lenth is equal to the shortest path's length-1 it means that even the current room is the end
			// the current path will have the same length as the shortest and will intersect with it
			if currentPath.len < shortestPath[i].len-1 {

				currentPath.add(currentRoom)
				if currentRoom == end {
					shortestPath[i] = currentPath
					return
				}

				for _, room := range currentRoom.links {
					if room != start && !currentPath.isVisited(room) {
						wg.Add(1)
						searchPath(i, currentPath, room)
					}
				}
			}
		}
		// search the shortest path for every room linked to start
		currentPath := path{len: 0}
		wg.Add(1)
		go searchPath(countStatrtsLinks, currentPath, roomNextToStart)
		countStatrtsLinks++

	}
	wg.Wait()
	for _, shP := range shortestPath {
		if shP.len < math.MaxInt {
			notIntersectedPaths = append(notIntersectedPaths, shP)
		}
	}

	return
}

func New(farm *antFarm) exploredRooms {
	ers := make(exploredRooms)
	for _, room := range farm.rooms {
		ers[room.name] = &exploredRoom{room: room, label: 0}
	}

	ers[farm.Start.name].label = 2
	return ers
}

func (ers exploredRooms) isExplored(r *Room) bool {
	if ers[r.name].label == 0 {
		return false
	}
	return true
}

/*
sets status of exploring for the given room
*/
func (ers exploredRooms) setStatus(r *Room, stat byte) {
	ers[r.name].label = stat
}

/*
mark the given room as unexplored
*/
func (ers exploredRooms) setStatusUnexplored(r *Room) {
	ers.setStatus(r, 0)
	ers[r.name].parent = nil
}

/*
mark the given room as explored
*/
func (ers exploredRooms) setStatusExplored(r *Room, parent *Room) {
	ers.setStatus(r, 1)
	ers[r.name].parent = parent
}

/*
mark the given room as places in a path
*/
func (ers exploredRooms) setStatusInPath(r *Room) {
	ers.setStatus(r, 2)
}



func (p *path) add(r *Room) {
	*p = append(*p, r)
}

func (p path) String() string {
	res := fmt.Sprintf("len:%d, cap:%d addr: %p. strt  ", len(p.path), cap(p.path), p.path)
	for _, r := range p.path {
		res += fmt.Sprintf("-%s- ", r.name)
	}
	res += fmt.Sprintf("\nlen: %d\n", p.len)
	return res
}
