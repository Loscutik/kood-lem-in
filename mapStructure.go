package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
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
				fmt.Println(scanner.Text()) // Should be attached to the map instead
			}
			if (len(scanner.Text()) > 0) && (r == 32) { // Searches for rooms & coordinates
				fmt.Println(scanner.Text()) // Should be attached to the map instead
				break
			}
			i, err := strconv.Atoi(scanner.Text()) // If there is only one number in the line - according to the rules this can be only the number of ants
			if err == nil {
				fmt.Println("Number of ants:", i)
				break
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
