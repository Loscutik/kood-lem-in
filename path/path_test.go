package path

import (
	"fmt"
	"testing"

	"lemin/room"
)

func BenchmarkSearchAllNotIntersectedPathes(b *testing.B) {
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

	start.Links = []*room.Link{
		{Room: &r1, Flow: 0},
		{Room: &r7, Flow: 0},
		{Room: &end, Flow: 0},
	}
	r1.Links = []*room.Link{
		{Room: &r2, Flow: 0},
		{Room: &r5, Flow: 0},
		{Room: &r6, Flow: 0},
		{Room: &start, Flow: 0},
	}
	r2.Links = []*room.Link{
		{Room: &r1, Flow: 0},
		{Room: &r3, Flow: 0},
	}
	r3.Links = []*room.Link{
		{Room: &r2, Flow: 0},
		{Room: &r4, Flow: 0},
		{Room: &r5, Flow: 0},
	}
	r4.Links = []*room.Link{
		{Room: &r3, Flow: 0},
		{Room: &r5, Flow: 0},
		{Room: &r6, Flow: 0},
		{Room: &end, Flow: 0},
	}
	r5.Links = []*room.Link{
		{Room: &r1, Flow: 0},
		{Room: &r3, Flow: 0},
		{Room: &r4, Flow: 0},
		{Room: &r5, Flow: 0},
	}
	r6.Links = []*room.Link{
		{Room: &r1, Flow: 0},
		{Room: &r4, Flow: 0},
		{Room: &r8, Flow: 0},
		{Room: &end, Flow: 0},
	}
	r7.Links = []*room.Link{
		{Room: &start, Flow: 0},
		{Room: &end, Flow: 0},
	}
	r8.Links = []*room.Link{
		{Room: &r6, Flow: 0},
	}
	end.Links = []*room.Link{
		{Room: &r4, Flow: 0},
		{Room: &r6, Flow: 0},
		{Room: &r7, Flow: 0},
		{Room: &start, Flow: 0},
	}

	pathes := SearchAllPaths(&farm, &pathHandler{isAllowedToGo, changeFlow, UNEXPLORED})
	for i, p := range pathes {
		fmt.Printf("path# %d: %s\n", i, p)
	}
}

func TestTrickedFarmExm01(tt *testing.T) {
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

	fmt.Println("--just BFS--")
	pathes := SearchAllPaths(&farm, &pathHandler{isJustBFS, noFlows, ON_PATH})
	for i, p := range pathes {
		fmt.Printf("path# %d: %s\n", i, p)
	}

	fmt.Println("--preparing: mark links usin Edmound-Karp algorithm--")
	pathes = SearchAllPaths(&farm, &pathHandler{isAllowedToGo, changeFlow, UNEXPLORED})
	for i, p := range pathes {
		fmt.Printf("path# %d: %s\n", i, p)
	}

	fmt.Println("--make intersected paths--")
	pathes = SearchAllPaths(&farm, &pathHandler{isPartOfPaths, resetFlow, ON_PATH})
	for i, p := range pathes {
		fmt.Printf("path# %d: %s\n", i, p)
	}
}

func TestSearchAllNotIntersectedPathes2(tt *testing.T) {
	start := room.Room{Name: "start"}
	a := room.Room{Name: "a"}
	b := room.Room{Name: "b"}
	c := room.Room{Name: "c"}
	d := room.Room{Name: "d"}
	e := room.Room{Name: "e"}
	f := room.Room{Name: "f"}
	end := room.Room{Name: "end"}
	farm := room.AntFarm{
		Rooms: []*room.Room{
			&start, &a, &b, &c, &d, &e, &f, &end,
		},
		Start: &start,
		End:   &end,
	}

	start.Links = []*room.Link{
		{Room: &a, Flow: 0},
		{Room: &b, Flow: 0},
	}
	a.Links = []*room.Link{
		{Room: &start, Flow: 0},
		{Room: &c, Flow: 0},
		{Room: &e, Flow: 0},
	}
	b.Links = []*room.Link{
		{Room: &start, Flow: 0},
		{Room: &d, Flow: 0},
	}
	c.Links = []*room.Link{
		{Room: &a, Flow: 0},
		{Room: &f, Flow: 0},
	}
	d.Links = []*room.Link{
		{Room: &b, Flow: 0},
		{Room: &e, Flow: 0},
	}
	e.Links = []*room.Link{
		{Room: &a, Flow: 0},
		{Room: &d, Flow: 0},
		{Room: &end, Flow: 0},
	}
	f.Links = []*room.Link{
		{Room: &c, Flow: 0},
		{Room: &end, Flow: 0},
	}
	end.Links = []*room.Link{
		{Room: &e, Flow: 0},
		{Room: &f, Flow: 0},
	}

	fmt.Println("--preparing: mark links usin Edmound-Karp algorithm--")
	pathes := SearchAllPaths(&farm, &pathHandler{isAllowedToGo, changeFlow, UNEXPLORED})
	for i, p := range pathes {
		fmt.Printf("path# %d: %s\n", i, p)
	}

	fmt.Println("--make intersected paths--")
	pathes = SearchAllPaths(&farm, &pathHandler{isPartOfPaths, resetFlow, ON_PATH})
	for i, p := range pathes {
		fmt.Printf("path# %d: %s\n", i, p)
	}
}

func TestPathesMustBeIntersected(t *testing.T) {
	start := room.Room{Name: "s"}
	r1 := room.Room{Name: "1"}
	r2 := room.Room{Name: "2"}
	r3 := room.Room{Name: "3"}
	r4 := room.Room{Name: "4"}
	end := room.Room{Name: "e"}
	farm := room.AntFarm{
		Rooms: []*room.Room{
			&start, &r1, &r2, &r3, &r4,&end,
		},
		Start: &start,
		End:   &end,
	}

	start.Links = []*room.Link{
		{Room: &r1, Flow: 0},
		{Room: &r2, Flow: 0},
	}
	r1.Links = []*room.Link{
		{Room: &start, Flow: 0},
		{Room: &r3, Flow: 0},
		{Room: &r4, Flow: 0},
		{Room: &end, Flow: 0},
	}
	r2.Links = []*room.Link{
		{Room: &start, Flow: 0},
		{Room: &r3, Flow: 0},
	}
	r3.Links = []*room.Link{
		{Room: &r1, Flow: 0},
		{Room: &r2, Flow: 0},
	}
	r4.Links = []*room.Link{
		{Room: &r1, Flow: 0},
		{Room: &end, Flow: 0},
	}
	end.Links = []*room.Link{
		{Room: &r1, Flow: 0},
		{Room: &r4, Flow: 0},
	}

	pathes := SearchAllPaths(&farm, &pathHandler{isAllowedToGo, changeFlow, UNEXPLORED})
	for i, p := range pathes {
		fmt.Printf("path# %d: %s\n", i, p)
	}

	fmt.Println("--del intersected--")
	pathes = SearchAllPaths(&farm, &pathHandler{isPartOfPaths, resetFlow, ON_PATH})
	for i, p := range pathes {
		fmt.Printf("path# %d: %s\n", i, p)
	}
}
