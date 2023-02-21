package ants

import "lemin/path"

type (
	pathesInRound  int // the number of used pathes in a round
	pathesByRounds []pathesInRound
)

type antOnPath struct {
	ant      int
	path     *path.Path
	position int
}

type antsAtRound []antOnPath

type antsByRounds []antsAtRound
