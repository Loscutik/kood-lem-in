package ants

import "lemin/path"

type (
	// each item of the slice presents the number of ants
	// which are going to go by the path with an index equal to that slice index
	// so if the 1st item of an element of numberOfAntsOnPaths =3, it means that 3 ants will go by the path N3
	numberOfAntsOnPaths []int
)

// determine the quantity of ants for each path
func calculateNumberOfAntsOnPaths(numberOfAnts int, paths []*path.Path) numberOfAntsOnPaths {
	if paths==nil || numberOfAnts==0 {return nil}
	antsNumbers:= make(numberOfAntsOnPaths, len(paths))
	i:=0
	antsNumbers[0]=1
	for a := numberOfAnts-1; a >0; a-- {
		i++
		if i>=len(paths) || paths[i].Len()-paths[0].Len()>/*>=*/a{i=0}
		antsNumbers[i]++		
	}
	return antsNumbers
}

func calculateNumberOfAntsOnPaths2(numberOfAnts int, paths []*path.Path) numberOfAntsOnPaths {
	if paths==nil || numberOfAnts==0 {return nil}
	antsNumbers:= make(numberOfAntsOnPaths, len(paths))
	i:=0
	antsNumbers[0]=1
	for a := numberOfAnts; a >1; a-- {
		i++
		if i>=len(paths) || paths[i].Len() + antsNumbers[i]>/*>=*/paths[i-1].Len()+antsNumbers[i-1]{i=0}
		antsNumbers[i]++		
	}
	return antsNumbers
}