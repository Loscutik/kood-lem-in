package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type AntFarm struct {
	Rooms      []*Room
	Start, End *Room
}
type Room struct {
	Name  string   // Done
	Links []string // Done
	x, y  int      // coordinates
}

func (r *Room) CreateFarmStruct() {
	TunnelList := []string{} // The slice array of tunnels
	RoomList := []string{}   // The slice array of room names and coordinates
	var NumberOfAnts int     // The number of ants from the source file
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Please enter the source text file name as the first argument!")
		os.Exit(1)
	}
	SourceFile := os.Args[1]
	readFile, _ := os.Open(SourceFile)
	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines) // Sets the split function for the scanning operation.
	count := 0                     // Counts the lines.
	for scanner.Scan() {           // This scanner gets the file content line by line
		count++
		for _, h := range scanner.Text() {
			if (len(scanner.Text()) == 3) && (h == 45) { // Searches for tunnel pairs
				fmt.Println(scanner.Text())                     // Prints tunnel pairs to the terminal
				TunnelList = append(TunnelList, scanner.Text()) // Ready list of tunnels made as a string array
				break
			}
			if (len(scanner.Text()) > 0) && (h == 32) { // Searches for rooms & coordinates
				if r.CheckRoomList(scanner.Text()) { // Checks if the string with spaces contains the room name and coordinates
					fmt.Println(scanner.Text())                 // Prints room coordinates to the terminal
					RoomList = append(RoomList, scanner.Text()) // Ready list of rooms with coordinates made as a string array
					break
				}
			}
			i, err := strconv.Atoi(scanner.Text()) // If there is only one number in the line - according to the rules this can be only the number of ants
			if err == nil {
				NumberOfAnts = i // Creates the variable of the number of ants from source file to use in other functions
				fmt.Println("Number of ants:", NumberOfAnts)
				break
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Name:", r.Name)
	//fmt.Println("Coordinates:", r.x, r.y)
	r.Links = TunnelList
	fmt.Println("Tunnels:", r.Links)
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
				if k[0] != 76 && k[0] != 35 {
					r.Name = r.Name + k + " " // NB, NEEDS TO BE CONVERTED TO MAP
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
			fmt.Println("X:", k)
		case 2: // The third part - coordinates, y
			j, err := strconv.Atoi(k) // Checking if it is int
			if err == nil {
				r.y = j
			}
			if err != nil {
				return false
			}
			fmt.Println("Y:", k)
		}
	}
	return true
}
