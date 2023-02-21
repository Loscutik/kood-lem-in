package ants



type (
	//pathesInRound  int // the number of used pathes in a round

	// each item of the slice presents the number of ants
	// which are going to go by the path with an index equal to that slice index
	// so if the 1st item of an element of numberOfAntsOnPaths =3, it means that 3 ants will go by the path N3
	numberOfAntsOnPaths []int
) 