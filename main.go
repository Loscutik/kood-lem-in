package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"lemin/ants"
	"lemin/path"
	"lemin/room"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Please enter the source text file name as the first argument!")
		os.Exit(1)
	}
	SourceFile := os.Args[1]
	file, err := os.Open(SourceFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	numberOfAnts, farm, err := room.CreateFarm(file)
	if err != nil {
		log.Fatalf("ERROR: invalid data format: %v",err)
	}

	start:=time.Now()
	paths := path.SearchAllNotIntersectedPaths(farm)

	// for i, p := range paths {
	// 	fmt.Printf("path# %d: %s\n", i, p)
	// }

	ants.AntsGo(numberOfAnts, paths)
	fmt.Printf("finished in %v \n", time.Since(start))
}
