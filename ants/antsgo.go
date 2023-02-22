package ants

import (
	"fmt"
	"log"

	"lemin/path"
	"lemin/room"
)

type PathForAnt struct {
	Path          *path.Path
	StartingRound int
}


// TODO Later don't foget to check the case when there are no paths at all
// TODO fix func AntsGo after getting calculateNumberOfAntsOnPaths(numberOfAnts int, paths []*path.Path)

func AntsGo (numberOfAnts int, paths []*path.Path) {
	numberOfAntsOnPaths:=numberOfAntsOnPaths{} // must be a function
	pathsForAnts:=assigntAntsPathsAndOrder(numberOfAnts,paths,numberOfAntsOnPaths)
	printAntMoving(numberOfAnts,pathsForAnts)
}

/*
leads ants trough their paths and write the result using given function
*/
func printAntMoving(numberOfAnts int, antsByPaths []PathForAnt) {
	// everything starts with zero
	antNext := 0
	round := 0
	antsInTheEnd := 0
	// ... and doesn't finish until reach the end
	for antsInTheEnd < numberOfAnts {
		// 1st: let's all ants who already are on their way to end move to the next room
		for ant, antWay := range antsByPaths[:antNext] {
			passedRoomsInPath := round - antWay.StartingRound
			if passedRoomsInPath < antWay.Path.Len() {
				printAntInRoom(ant, antWay.Path.GetRoom(passedRoomsInPath))
			} else {
				antsInTheEnd++
			}
		}

		// 2d: the next band of ants is starting (if they still are in the start room)
		for antNext < numberOfAnts && antsByPaths[antNext].StartingRound == round {
			printAntInRoom(antNext, antsByPaths[antNext].Path.GetRoom(0))
			antNext++
		}
		fmt.Println()
		round++
	}
}

/*
assigns a path and a round to start moving to each ant. 
It supposes to work with correct data, e.g. length of numberOfAntsOnPaths must be not greater than length of paths.
In the other case it will panic. 
The function changes slice numberOfAntsOnPaths, each item of slice will be equal to 0
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
			numberOfAntsOnPaths[pathNumber]--
			antNext++

		}
		round++
	}

	return antsByPaths
}

/*
prints ant's position in Stdout in the format L#ant-room.Name
*/
func printAntInRoom(ant int, room *room.Room) {
	fmt.Printf("L%d-%s ", ant+1, room.Name)
}
