package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func CreateFarmStruct() {
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
		for _, r := range scanner.Text() {
			if (len(scanner.Text()) == 3) && (r == 45) { // Searches for tunnel pairs
				fmt.Println(scanner.Text())                     // Prints tunnel pairs to the terminal
				TunnelList = append(TunnelList, scanner.Text()) // Ready list of tunnels made as a string array
				break
			}
			// Need to add check for int coordinates
			if (len(scanner.Text()) > 0) && (r == 32) { // Searches for rooms & coordinates
				if CheckRoomList(scanner.Text()) == true { // Checks if the string with spaces contains the room name and coordinates
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
}

func CheckRoomList(a string) bool {
	var aa string
	var RoomName string
	wasSpace := false
	for i := 0; i < len(a)-1; i++ {
		if a[0] != 76 && a[0] != 35 && a[i] != 32 && wasSpace == false {
			RoomName = RoomName + string(a[i])
		}
		if a[i] == 32 && wasSpace == false {
			wasSpace = true // Checks only the string after the first space because room name can be not int
		}
		if a[i] != 32 && wasSpace == true {
			aa = string(a[i])
			_, err := strconv.Atoi(aa)
			if err != nil {
				return false
			}
		}
	}
	fmt.Println("Room name:", RoomName)
	return true
}
