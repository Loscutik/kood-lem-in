package main

import (
	"fmt"
	"math"
	"sync"
)

type path struct {
	path []*room
	len  int
}
/*
returns all not intersectons pathes from start to end. Start is not include into the path
*/
func searchAllNotIntersectedPathes(start *room, end *room) (notIntersectedPathes[]path) {
	wg:=&sync.WaitGroup{}

	// looking for pathes for each room conected to start
	// shortestPath keeps the path that is the shortest at the determined moment of searching 
	// create shortestPath for each start's link
	shortestPath:=make([]path, len(start.links))
	countStatrtsLinks:=0

	for _, roomNextToStart := range start.links {
		shortestPath[countStatrtsLinks] = path{len: math.MaxInt}
		var searchPath func(i int, currentPath path, currentRoom *room)

		
		searchPath = func(i int, currentPath path, currentRoom *room) {
			defer wg.Done()
			// it doesn't need to compare all path.
			// If the current path's lenth is equal to the shortest path's length-1 it means that even the current room is the end
			// the current path will have the same length as the shortest and will intersect with it
			if currentPath.len < shortestPath[i].len-1{
				
				currentPath.add(currentRoom)
				if currentRoom == end {
					shortestPath[i] = currentPath
					return
				}
				
				for _, room := range currentRoom.links {
					if room !=start && !currentPath.isVisited(room) {
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
	for _,shP:=range shortestPath{
		if shP.len<math.MaxInt{
			notIntersectedPathes=append(notIntersectedPathes, shP)
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


