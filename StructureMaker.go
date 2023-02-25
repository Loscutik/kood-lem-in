package main

import "fmt"

type AntFarm struct {
	Rooms      []*Room
	Start, End *Room
}
type Room struct {
	Name  string
	Links []*Room
	x, y  int // coordinates
}

func StructureMaker() {
	fmt.Println("Number of ants:", NumberOfAnts)
	fmt.Println("Room name:", RoomName)
	fmt.Println(TunnelList)
	fmt.Println(RoomList)
}
