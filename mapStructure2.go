package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type AntFarm struct {
	Rooms []*Room
	Start *Room
	End   *Room
}
type Room struct {
	Name  string
	Links []*Room // The array of []*Room structures of the Name+x+y rooms linked to the current room.
	x     int     // X coordinates
	y     int     // Y coordinates
}

var TunnelList []string // The slice array of tunnels

func CreateFarmStruct() (int, *AntFarm, error) {
	room := &Room{}
	var NumberOfAnts int // The number of ants from the source file
	f := &AntFarm{
		Rooms: []*Room{},
		Start: &Room{},
		End:   &Room{},
	}
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Please enter the source text file name as the first argument!")
		os.Exit(1)
	}
	SourceFile := os.Args[1]
	readFile, _ := os.Open(SourceFile)
	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines) // Sets the split function for the scanning operation.
	isStart := false
	isEnd := false
	for scanner.Scan() { // This scanner gets the file content line by line
		txt := scanner.Text()
		if len(txt) == 0 {
			continue
		}
		switch {
		case strings.HasPrefix(txt, "#"):
			if scanner.Text() == "##start" {
				isStart = true
			}
			if scanner.Text() == "##end" {
				isEnd = true
			}
		case strings.Contains(txt, " "):
			if room.CheckRoomList(scanner.Text()) { // Checks if the string with spaces contains the room name and coordinates
				room.Links = f.CreateLinks(TunnelList, room)
				if isStart {
					f.Start = &Room{
						Name:  room.Name,
						Links: room.Links,
						x:     room.x,
						y:     room.y,
					}
					fmt.Println("START: ", f.Start.Name, f.Start.x, f.Start.y)
					isStart = false
				}
				if isEnd {
					f.End = &Room{
						Name:  room.Name,
						Links: room.Links,
						x:     room.x,
						y:     room.y,
					}
					fmt.Println("END: ", f.End.Name, f.End.x, f.End.y)
					isEnd = false
				} else {
					room = &Room{
						Name:  room.Name,
						Links: room.Links,
						x:     room.x,
						y:     room.y,
					}
				}
				fmt.Println(room.Name, room.x, room.y)
				f.Rooms = append(f.Rooms, room)
			}
		case strings.Contains(txt, "-"):
			TunnelList = append(TunnelList, scanner.Text())
		default:
			i, err := strconv.Atoi(scanner.Text()) // If there is only one number in the line - according to the rules this can be only the number of ants
			if err == nil {
				NumberOfAnts = i // Creates the variable of the number of ants from source file to use in other functions
				fmt.Println("Number of ants:", NumberOfAnts)
			}
		}
	}
	fmt.Println("Tunnellist:", TunnelList)
	for _, k := range f.Rooms {
		k.Links = f.CreateLinks(TunnelList, room) // Пока дичь, needs to be improved
		fmt.Println("Name, Links", room.Name, room.Links)
	}
	return NumberOfAnts, f, nil
}

func (f *AntFarm) CreateLinks(TunnelList []string, room *Room) []*Room {

	var r1 string // room 1
	var r2 string // room 2
	fmt.Printf("TL: %v", TunnelList)
	fmt.Printf("TL: %v", len(TunnelList))
	for _, j := range TunnelList {

		// find  rooms for tunnels in the slice f.Room (r1:=f.findRoomByName(name1)). if there is no the room -error
		// add tunnel to the both rooms
		//r1.Links = append([]*Room, r2)
		//r2.Links = append([]*Room, r1)
		re := regexp.MustCompile("-")
		split := re.Split(j, -1) // split is the room pair
		for i := range split {
			r1 = split[i]
			r2 = split[i+1]
			break
		}
		fmt.Println("Room 1; Room 2:", r1, r2)
		for _, g := range f.Rooms {
			if g.Name == r2 && f.IsThereRoom(r1) {
				g = room
				room.Links = append(room.Links, g)
			}
		}
		// counter add
	}
	return room.Links
}

func (f *AntFarm) IsThereRoom(name string) bool {
	for _, r := range f.Rooms {
		if r.Name == name {
			return true
		}
	}
	return false
}

func (r *Room) CheckRoomList(a string) bool {
	StringArray := strings.Fields(a)
	for h, k := range StringArray {
		if len(StringArray) > 3 { // Checking if there are 3 parts of the string - Name, x and y
			return false
		}
		switch h {
		case 0: // The first part - the name of the room
			for i := 0; i < len(k); i++ {
				if k[0] != '#' && k[0] != 'L' {
					r.Name = k
					break
				} else {
					return false
				}
			}
		case 1: // The second part - coordinates, x
			j, err := strconv.Atoi(k) // Checking if it is int
			if err == nil {
				r.x = j
			}
			if err != nil {
				return false
			}
		case 2: // The third part - coordinates, y
			j, err := strconv.Atoi(k) // Checking if it is int
			if err == nil {
				r.y = j
			}
			if err != nil {
				return false
			}
		}
	}
	return true
}
