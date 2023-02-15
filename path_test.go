package main

import (
	"fmt"
	"testing"
)

func BenchmarkSearchAllNotIntersectedPathes(b *testing.B) {
	start := room{name: "start"}
	r1 := room{name: "1"}
	r2 := room{name: "2"}
	r3 := room{name: "3"}
	r4 := room{name: "4"}
	r5 := room{name: "5"}
	r6 := room{name: "6"}
	r7 := room{name: "7"}
	r8 := room{name: "8"}
	end := room{name: "end"}

	start.links = map[string]*room{
		"1": &r1,
		"7": &r7,
	}
	r1.links = map[string]*room{
		"2": &r2,
		"5": &r5,
		"6": &r6,
		"start": &start,
	}
	r2.links = map[string]*room{
		"1": &r1,
		"3": &r3,
	}
	r3.links = map[string]*room{
		"2": &r2,
		"4": &r4,
		"5": &r5,
	}
	r4.links = map[string]*room{
		"3": &r3,
		"5": &r5,
		"6": &r6,
		"end": &end,
	}
	r5.links = map[string]*room{
		"1": &r1,
		"3": &r3,
		"4": &r4,
		"5": &r5,
	}
	r6.links = map[string]*room{
		"1": &r1,
		"4": &r4,
		"8": &r8,
		"end": &end,
	}
	r7.links = map[string]*room{
		"start": &start,
		"end": &end,
	}
	r8.links = map[string]*room{
		"6": &r6,
	}
	end.links = map[string]*room{
		"4": &r4,
		"6": &r6,
	}

	pathes:=searchAllNotIntersectedPathes(&start,&end)
	for i,p:=range pathes{
		fmt.Printf("path# %d: %s",i,p)
	}
	
}
