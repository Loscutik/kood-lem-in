package main

import (
	"fmt"
	"math"
)

type path struct {
	path []*room
	len  int
}

func searchAllNotIntersectedPathes(start *room, end *room) (notIntersectedPathes[]path) {
	for _, roomNextToStart := range start.links {
		shortestPath := path{len: math.MaxInt}
		
		var searchPath func(currentPath path, currentRoom *room)
		searchPath = func(currentPath path, currentRoom *room) {
			// it doesn't need to compare all path.
			// If the current path's lenth is equal to the shortest path's length-1 it means that even the current room is the end
			// the current path will have the same length as the shortest and will intersect with it
			//fmt.Printf("!!!!\nfunc's start. currentPath: %sshortest len: %d\ncurrentRoom: %s\n\n", currentPath,shortestPath.len, currentRoom.name)
			if currentPath.len < shortestPath.len-1{
				
				currentPath.add(currentRoom)
				if currentRoom == end {
					shortestPath = currentPath
					return
				}
				
				for _, room := range currentRoom.links {
					//fmt.Printf("??visited. room's name: %s \ncurrentPath: %s *\n",room.name, currentPath)
					if room !=start && !currentPath.isVisited(room) {
						//fmt.Printf("()before the recurse. currentPath: %sroom's name: %s \n***\n\n", currentPath,room.name)
						searchPath(currentPath, room)
					}
				}
			}
		}
		// search the shortest path for every room linked to start
		currentPath := path{len: 0}
		searchPath(currentPath, roomNextToStart)

		if shortestPath.len<math.MaxInt{
			notIntersectedPathes=append(notIntersectedPathes, shortestPath)
		}
	}
	return
}

func (p *path) add(r *room) {
	p.path = append(p.path, r)
	p.len++
}

func (p *path) isVisited(r *room) bool{
	for _,roomOnPath:= range p.path{
		if r.name==roomOnPath.name {return true}
	}
	return false
}

func (p path)String() string{
	res:=fmt.Sprintf("len:%d, cap:%d addr: %p. strt  ",len(p.path),cap(p.path),p.path)
	for _,r:=range p.path{
		res+=fmt.Sprintf("-%s- ",r.name)
	}
	res+=fmt.Sprintf("\nlen: %d\n", p.len)
	return res
}


