package main

type (
	pathesInRound  int // the number of used pathes in a round
	pathesByRounds []pathesInRound
)

type antOnPath struct {
	ant      int
	path     *path
	position int
}

type antsAtRound []antOnPath

type antsByRounds []antsAtRound
