package room

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type coord struct {
	x, y int
}

type Link struct {
	Room *Room
	Flow byte 
}

type Room struct {
	Name  string
	Links []*Link
	Coord coord
}

func (r *Room)GetLinkTo(to *Room) *Link{
	for _, l := range r.Links  {
		if l.Room==to{
			return l
		}
	}
	return nil
}

func (r *Room) AddLink(neibor *Room) {
	// if the Links already has a connection to the same room don't duplikate the link (to not allow rooms to contain more than 1 ant)
	for _, r := range r.Links {
		if r.Room == neibor {
			return
		}
	}
	r.Links = append(r.Links, &Link{neibor,0})
}

func CreateFarm(readFile *os.File) (int, *AntFarm, error) {
	var TunnelList []string // The slice of tunnels for printing
	var numberOfAnts int    // The number of ants from the source file
	farm := &AntFarm{
		Rooms: []*Room{},
		Start: &Room{},
		End:   &Room{},
	}

	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines) // Sets the split function for the scanning operation.
	isStart, isEnd := 0, 0
	rowCounter := 0

	// reading file row by row and creating a farm
	for scanner.Scan() { // This scanner gets the file content line by line
		txt := scanner.Text()
		rowCounter++
		if len(txt) == 0 {
			continue
		}

		// check comments
		if strings.HasPrefix(txt, "#") {
			if scanner.Text() == "##start" {
				if isStart>0 {
					return 0, nil, fmt.Errorf("row %d: there must be only one start room", rowCounter)
				}
				isStart = 1
				continue
			}
			if scanner.Text() == "##end" {
				if isEnd>0 {
					return 0, nil, fmt.Errorf("row %d: there must be only one end room", rowCounter)
				}
				isEnd = 1
				continue
			}
			continue
		}

		// check rooms
		reg := regexp.MustCompile(`^([^-]+) (\d{1,3}) (\d{1,3})$`)
		submatches := reg.FindStringSubmatch(scanner.Text())
		if submatches != nil {
			// convert the coordinates
			x, err := strconv.Atoi(submatches[2])
			if err != nil {
				return 0, nil, fmt.Errorf("row %d: incorect coordinates", rowCounter)
			}

			y, err := strconv.Atoi(submatches[3])
			if err != nil {
				return 0, nil, fmt.Errorf("row %d: incorect coordinates", rowCounter)
			}

			room, err := farm.AddRoom(submatches[1], x, y)
			if err != nil {
				return 0, nil, fmt.Errorf("row %d: %v", rowCounter, err)
			}

			if isStart==1 {
				farm.Start = room
				isStart++
			}
			if isEnd==1 {
				farm.End = room
				isEnd++
			}
			continue
		}

		// check tunnels
		reg = regexp.MustCompile(`^([^-]+)-([^-]+)$`)
		submatches = reg.FindStringSubmatch(scanner.Text())
		if submatches != nil {
			room1 := farm.findRoom(submatches[1])
			room2 := farm.findRoom(submatches[2])
			if room1 != nil && room2 != nil {
				room1.AddLink(room2)
				room2.AddLink(room1)
			}
			TunnelList = append(TunnelList, scanner.Text())
			continue
		}

		// check ants' number
		reg = regexp.MustCompile(`^\d+$`)
		submatches = reg.FindStringSubmatch(scanner.Text())
		if submatches != nil {
			n, err := strconv.Atoi(submatches[0])
			if err != nil {
				return 0, nil, fmt.Errorf("row %d: incorect quantity of ants", rowCounter)
			}
			numberOfAnts = n
		}
	}

	// printing the farm
	fmt.Println(numberOfAnts)
	for _, r := range farm.Rooms {
		if r == farm.Start {
			fmt.Println("##start")
		}
		if r == farm.End {
			fmt.Println("##end")
		}
		fmt.Printf("%s %d %d\n", r.Name, r.Coord.x, r.Coord.y)
	}
	for _, t := range TunnelList {
		fmt.Println(t)
	}

	fmt.Println()

	return numberOfAnts, farm, nil
}
