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
		{Room: &r1,Flow: 0},
		{Room: &r7,Flow: 0},
		{Room: &end,Flow: 0},
	}
	r1.Links = []*room.Link{
		{Room: &r2,Flow: 0},
		{Room: &r5,Flow: 0},
		{Room: &r6,Flow: 0},
		{Room: &start,Flow: 0},
	}
	r2.Links = []*room.Link{
		{Room: &r1,Flow: 0},
		{Room: &r3,Flow: 0},
	}
	r3.Links = []*room.Link{
		{Room: &r2,Flow: 0},
		{Room: &r4,Flow: 0},
		{Room: &r5,Flow: 0},
	}
	r4.Links = []*room.Link{
		{Room: &r3,Flow: 0},
		{Room: &r5,Flow: 0},
		{Room: &r6,Flow: 0},
		{Room: &end,Flow: 0},
	}
	r5.Links = []*room.Link{
		{Room: &r1,Flow: 0},
		{Room: &r3,Flow: 0},
		{Room: &r4,Flow: 0},
		{Room: &r5,Flow: 0},
	}
	r6.Links = []*room.Link{
		{Room: &r1,Flow: 0},
		{Room: &r4,Flow: 0},
		{Room: &r8,Flow: 0},
		{Room: &end,Flow: 0},
	}
	r7.Links = []*room.Link{
		{Room: &start,Flow: 0},
		{Room: &end,Flow: 0},
	}
	r8.Links = []*room.Link{
		{Room: &r6,Flow: 0},
	}
	end.Links = []*room.Link{
		{Room: &r4,Flow: 0},
		{Room: &r6,Flow: 0},
		{Room: &r7,Flow: 0},
		{Room: &start,Flow: 0},
	}

	pathes := SearchAllPaths(&farm, &pathHandler{isAllowedToGo, changeFlow, UNEXPLORED})
	for i, p := range pathes {
		fmt.Printf("path# %d: %s\n", i, p)
	}
}

func TestSearchAllNotIntersectedPathes1(t *testing.T) {
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
		{Room: &r1,Flow: 0},
		{Room: &r7,Flow: 0},
		{Room: &end,Flow: 0},
	}
	r1.Links = []*room.Link{
		{Room: &r2,Flow: 0},
		{Room: &r5,Flow: 0},
		{Room: &r6,Flow: 0},
		{Room: &start,Flow: 0},
	}
	r2.Links = []*room.Link{
		{Room: &r1,Flow: 0},
		{Room: &r3,Flow: 0},
	}
	r3.Links = []*room.Link{
		{Room: &r2,Flow: 0},
		{Room: &r4,Flow: 0},
		{Room: &r5,Flow: 0},
	}
	r4.Links = []*room.Link{
		{Room: &r3,Flow: 0},
		{Room: &r5,Flow: 0},
		{Room: &r6,Flow: 0},
		{Room: &end,Flow: 0},
	}
	r5.Links = []*room.Link{
		{Room: &r1,Flow: 0},
		{Room: &r3,Flow: 0},
		{Room: &r4,Flow: 0},
		{Room: &r5,Flow: 0},
	}
	r6.Links = []*room.Link{
		{Room: &r1,Flow: 0},
		{Room: &r4,Flow: 0},
		{Room: &r8,Flow: 0},
		{Room: &end,Flow: 0},
	}
	r7.Links = []*room.Link{
		{Room: &start,Flow: 0},
		{Room: &end,Flow: 0},
	}
	r8.Links = []*room.Link{
		{Room: &r6,Flow: 0},
	}
	end.Links = []*room.Link{
		{Room: &r4,Flow: 0},
		{Room: &r6,Flow: 0},
		{Room: &r7,Flow: 0},
		{Room: &start,Flow: 0},
	}

	pathes := SearchAllPaths(&farm, &pathHandler{isAllowedToGo, changeFlow, UNEXPLORED})
	for i, p := range pathes {
		fmt.Printf("path# %d: %s\n", i, p)
	}
}
func TestSearchAllNotIntersectedPathes02(t *testing.T) {
	start := room.Room{Name: "0"}
	r1 := room.Room{Name: "1"}
	r2 := room.Room{Name: "2"}
	end := room.Room{Name: "3"}
	farm := room.AntFarm{
		Rooms: []*room.Room{
			&start, &r1, &r2, &end,
		},
		Start: &start,
		End:   &end,
	}

	start.Links = []*room.Link{
		{Room: &r1,Flow: 0},
		{Room: &end,Flow: 0},
	}
	r1.Links = []*room.Link{
		{Room: &start,Flow: 0},
		{Room: &r2,Flow: 0},
	}
	r2.Links = []*room.Link{
		{Room: &r1,Flow: 0},
		{Room: &end,Flow: 0},
	}
	end.Links = []*room.Link{
		{Room: &start,Flow: 0},
		{Room: &r2,Flow: 0},
	}

	pathes := SearchAllPaths(&farm, &pathHandler{isAllowedToGo, changeFlow, UNEXPLORED})
	for i, p := range pathes {
		fmt.Printf("path# %d: %s\n", i, p)
	}
}
