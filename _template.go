package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Open the input file
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}

	// Create scanner object to read in lines of file
	fileScanner := bufio.NewScanner(readFile)


	// Iterate over each line
	for fileScanner.Scan() {
		line := fileScanner.Text()
	}
}