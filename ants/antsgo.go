package ants

import (
	"log"

	"lemin/path"
	"lemin/room"
)

type PathForAnt struct {
	Path          *path.Path
	StartingRound int
}

// TODO funcs write...
// TODO test file 
// TODO Later don't foget to check the case when there are no paths at all
// TODO func AntsGo (numberOfAnts int, paths []*path.Path, writeAntInRoom func(ant int, room *room.Room), writeRoundStart func(round int))
// NEED func calculateNumberOfAntsOnPaths(numberOfAnts int, paths []*path.Path)

/*
leads ants trough their paths and write the result using given function
*/
func writeAntMoving(numberOfAnts int, antsByPaths []PathForAnt, writeAntInRoom func(ant int, room *room.Room), writeRoundStart func(round int)) {
	// everything starts with zero
	antNext := 0
	round := 0
	antsInTheEnd := 0
	// ... and doesn't finish until reach the end
	for antsInTheEnd < numberOfAnts {
		writeRoundStart(round)
		// 1st: let's all ants who already are on their way to end move to the next room
		for ant, antWay := range antsByPaths[:antNext] {
			passedRoomsInPath := round - antWay.StartingRound
			if passedRoomsInPath < antWay.Path.Len() {
				writeAntInRoom(ant, antWay.Path.GetRoom(passedRoomsInPath))
			} else {
				antsInTheEnd++
			}
		}

		// 2d: the next band of ants is starting (if they still are in the start room)
		for antNext < numberOfAnts && antsByPaths[antNext].StartingRound == round {
			writeAntInRoom(antNext, antsByPaths[antNext].Path.GetRoom(0))
			antNext++
		}

		round++
	}
}

/*
assigns a path and a round to start moving to each ant
*/
func assigntAntsPathsAndOrder(numberOfAnts int, paths []*path.Path, numberOfAntsOnPaths numberOfAntsOnPaths) []PathForAnt {
	// everything starts with zero
	antNext := 0
	round := 0
	// it keeps a path and round which an ant will start from. The index of the slice is equal to the ant number
	antsByPaths := make([]PathForAnt, numberOfAnts)
	for antNext < numberOfAnts {

		// walk through the list keeping the number of ants for each path
		// and save the path and the start round for each ant in slice antsByPaths
		for pathNumber, ants := range numberOfAntsOnPaths {
			// if ants run out for the path, stop using that path
			if ants == 0 {
				numberOfAntsOnPaths = numberOfAntsOnPaths[:pathNumber]
				break
			}
			if antNext == numberOfAnts {
				log.Fatalln("ERROR during assigning ants to paths: there are some extra ants waiting to go")
			}

			// the next ant gets its own path and the round to start
			antsByPaths[antNext] = PathForAnt{paths[pathNumber], round}
			antNext++

		}
		round++
	}

	return antsByPaths
}
