package ants

import (
	"fmt"
	"testing"

	"lemin/path"
	"lemin/room"
)

func TestPrintAnts(tt *testing.T) {
	start := room.Room{Name: "start"}
	r0 := room.Room{Name: "0"}
	o := room.Room{Name: "o"}
	n := room.Room{Name: "n"}
	e := room.Room{Name: "e"}
	t := room.Room{Name: "t"}
	E := room.Room{Name: "E"}
	a := room.Room{Name: "a"}
	m := room.Room{Name: "m"}
	h := room.Room{Name: "h"}
	A := room.Room{Name: "A"}
	c := room.Room{Name: "c"}
	k := room.Room{Name: "k"}
	end := room.Room{Name: "end"}
	farm := room.AntFarm{
		Rooms: []*room.Room{
			&start, &r0, &o, &n, &e, &t, &E, &a, &m, &h, &A, &c, &k, &end,
		},
		Start: &start,
		End:   &end,
	}

	start.Links = []*room.Link{
		{Room: &t, Flow: 0},
		{Room: &h, Flow: 0},
		{Room: &r0, Flow: 0},
	}
	r0.Links = []*room.Link{
		{Room: &o, Flow: 0},
		{Room: &start, Flow: 0},
	}
	o.Links = []*room.Link{
		{Room: &r0, Flow: 0},
		{Room: &n, Flow: 0},
	}
	n.Links = []*room.Link{
		{Room: &e, Flow: 0},
		{Room: &o, Flow: 0},
		{Room: &m, Flow: 0},
		{Room: &h, Flow: 0},
	}
	e.Links = []*room.Link{
		{Room: &n, Flow: 0},
		{Room: &end, Flow: 0},
	}
	t.Links = []*room.Link{
		{Room: &start, Flow: 0},
		{Room: &E, Flow: 0},
	}
	E.Links = []*room.Link{
		{Room: &a, Flow: 0},
		{Room: &t, Flow: 0},
	}
	a.Links = []*room.Link{
		{Room: &m, Flow: 0},
		{Room: &E, Flow: 0},
	}
	m.Links = []*room.Link{
		{Room: &a, Flow: 0},
		{Room: &n, Flow: 0},
		{Room: &end, Flow: 0},
	}
	h.Links = []*room.Link{
		{Room: &start, Flow: 0},
		{Room: &A, Flow: 0},
		{Room: &n, Flow: 0},
	}
	A.Links = []*room.Link{
		{Room: &c, Flow: 0},
		{Room: &h, Flow: 0},
	}
	c.Links = []*room.Link{
		{Room: &A, Flow: 0},
		{Room: &k, Flow: 0},
	}
	k.Links = []*room.Link{
		{Room: &end, Flow: 0},
		{Room: &c, Flow: 0},
	}
	end.Links = []*room.Link{
		{Room: &k, Flow: 0},
		{Room: &m, Flow: 0},
		{Room: &e, Flow: 0},
	}

	paths := path.SearchAllNotIntersectedPaths(&farm)

	ants := 5
	numberOfAntsOnPaths := numberOfAntsOnPaths{2, 2, 1}

	pathsForAnts := assigntAntsPathsAndOrder(ants, paths, numberOfAntsOnPaths)
	fmt.Printf("after func slice numberOfAntsOnPaths %#v\n", numberOfAntsOnPaths)
	printAntMoving(ants, pathsForAnts)
}
