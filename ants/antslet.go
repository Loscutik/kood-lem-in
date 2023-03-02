package ants

import "lemin/path"

type (
	// each item of the slice presents the number of ants
	// which are going to go by the path with an index equal to that slice index
	// so if the 1st item of an element of numberOfAntsOnPaths =3, it means that 3 ants will go by the path N3
	numberOfAntsOnPaths []int
)

// TODO  don't foget to check the case when there are no paths at all in func calculateNumberOfAntsOnPaths
func calculateNumberOfAntsOnPaths(numberOfAnts int, paths []*path.Path) numberOfAntsOnPaths {
	if paths==nil {return nil}
	var numbers numberOfAntsOnPaths

	numbers = nil // TODO if new path is added append, else numbers[i]++
	return numbers
}
