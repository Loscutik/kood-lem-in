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

/*
prints ants moving through the farm
*/
func AntsGo(numberOfAnts int, paths []*path.Path) {
	numberOfAntsOnPaths := calculateNumberOfAntsOnPaths(numberOfAnts, paths)
	pathsForAnts := assigntAntsPathsAndOrder(numberOfAnts, paths, numberOfAntsOnPaths)
	printAntMoving(numberOfAnts, pathsForAnts)
}

/*
leads ants trough their paths and write the result using given function
*/
func printAntMoving(numberOfAnts int, antsByPaths []PathForAnt) {
	if antsByPaths == nil {
		fmt.Println("No path")
		return
	}

	if numberOfAnts == 0 {
		fmt.Println("No ant seems to be in the farm")
		return
	}
	// everything starts with zero
	antNext := 0
	round := 0
	antsInTheEnd := 0
	// ... and doesn't finish until reach the end
	for antsInTheEnd < numberOfAnts {
		//fmt.Printf("-- round=%d, antNext=%d, antsInTheEnd=%d\n",round,antNext,antsInTheEnd)
		// 1st: let's all ants who already are on their way to end move to the next room
		for ant:=antsInTheEnd;ant<antNext; ant++ {
			passedRoomsInPath := round - antsByPaths[ant].StartingRound
			if passedRoomsInPath < antsByPaths[ant].Path.Len() {
				printAntInRoom(ant, antsByPaths[ant].Path.GetRoom(passedRoomsInPath))
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
	fmt.Printf("\nTotal rounds: %d\n", round-1)
}

/*
assigns a path and a round to start moving to each ant.
It supposes to work with correct data, e.g. length of numberOfAntsOnPaths must be not greater than length of paths.
In the other case it will panic.
The function changes slice numberOfAntsOnPaths, each item of slice will be equal to 0
*/
func assigntAntsPathsAndOrder(numberOfAnts int, paths []*path.Path, numberOfAntsOnPaths numberOfAntsOnPaths) []PathForAnt {
	if paths == nil || numberOfAntsOnPaths == nil {
		return nil
	}
	// everything starts with zero
	antNext := 0
	round := 0
	// it keeps a path and a round which an ant will start from. The index of the slice is equal to the ant number
	antsByPaths := make([]PathForAnt, numberOfAnts)
	for antNext < numberOfAnts {

		// walk through the list keeping the number of ants for each path
		// and save the path and the start round for each ant in slice antsByPaths
		//fmt.Printf("numberOfAntsOnPaths=%#v, round=%d, antNext=%d\n", numberOfAntsOnPaths, round, antNext)
		for pathNumber, ants := range numberOfAntsOnPaths {
			// if ants run out for the path, stop using that path
			//fmt.Printf("--ants=%d, pathNumber=%d --\n",ants, pathNumber)
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
