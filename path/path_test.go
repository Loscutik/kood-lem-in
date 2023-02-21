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
		Rooms: map[string]*room.Room{
			"start": &start, "1": &r1, "2": &r2, "3": &r3, "4": &r4, "5": &r5, "6": &r6, "7": &r7, "8": &r8, "end": &end,
		},
		Start: &start,
		End:   &end,
	}

	start.Links = map[string]*room.Room{
		"1": &r1,
		"7": &r7,
		"end": &end,
	}
	r1.Links = map[string]*room.Room{
		"2":     &r2,
		"5":     &r5,
		"6":     &r6,
		"start": &start,
	}
	r2.Links = map[string]*room.Room{
		"1": &r1,
		"3": &r3,
	}
	r3.Links = map[string]*room.Room{
		"2": &r2,
		"4": &r4,
		"5": &r5,
	}
	r4.Links = map[string]*room.Room{
		"3":   &r3,
		"5":   &r5,
		"6":   &r6,
		"end": &end,
	}
	r5.Links = map[string]*room.Room{
		"1": &r1,
		"3": &r3,
		"4": &r4,
		"5": &r5,
	}
	r6.Links = map[string]*room.Room{
		"1":   &r1,
		"4":   &r4,
		"8":   &r8,
		"end": &end,
	}
	r7.Links = map[string]*room.Room{
		"start": &start,
		"end":     &end,
	}
	r8.Links = map[string]*room.Room{
		"6": &r6,
	}
	end.Links = map[string]*room.Room{
		"4": &r4,
		"6": &r6,
	}

	pathes := searchAllNotIntersectedPaths(&farm)
	for i, p := range pathes {
		fmt.Printf("path# %d: %s\n", i, p)
	}
}
