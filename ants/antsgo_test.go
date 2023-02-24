package ants

import (
	"fmt"
	"testing"

	"lemin/path"
	"lemin/room"
)

func BenchmarkPrintAnts(b *testing.B) {
	start := room.Room{Name: "start"}
	r1 := room.Room{Name: "1"}
	r2 := room.Room{Name: "2"}
	r3 := room.Room{Name: "3"}
	r4 := room.Room{Name: "4"}
	r5 := room.Room{Name: "5"}
	r6 := room.Room{Name: "6"}
	r7 := room.Room{Name: "7"}
	r8 := room.Room{Name: "8"}
	end := room.Room{Name: "end"}
	farm := room.AntFarm{
		Rooms: []*room.Room{
			&start, &r1, &r2, &r3, &r4, &r5, &r6, &r7, &r8, &end,
		},
		Start: &start,
		End:   &end,
	}

	start.Links = []*room.Room{
		&r1,
		&r7,
		&end,
	}
	r1.Links = []*room.Room{
		&r2,
		&r5,
		&r6,
		&start,
	}
	r2.Links = []*room.Room{
		&r1,
		&r3,
	}
	r3.Links = []*room.Room{
		&r2,
		&r4,
		&r5,
	}
	r4.Links = []*room.Room{
		&r3,
		&r5,
		&r6,
		&end,
	}
	r5.Links = []*room.Room{
		&r1,
		&r3,
		&r4,
		&r5,
	}
	r6.Links = []*room.Room{
		&r1,
		&r4,
		&r8,
		&end,
	}
	r7.Links = []*room.Room{
		&start,
		&end,
	}
	r8.Links = []*room.Room{
		&r6,
	}
	end.Links = []*room.Room{
		&r4,
		&r6,
	}

	paths := path.SearchAllNotIntersectedPaths(&farm)

	ants := 5
	numberOfAntsOnPaths := numberOfAntsOnPaths{2, 2, 1}

	pathsForAnts := assigntAntsPathsAndOrder(ants, paths, numberOfAntsOnPaths)
	fmt.Printf("after func slice numberOfAntsOnPaths %#v\n", numberOfAntsOnPaths)
	printAntMoving(ants, pathsForAnts)
}
