package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Open the input file
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}

	// Create scanner object to read in lines of file
	fileScanner := bufio.NewScanner(readFile)

	total := 0
	// Iterate over each line
	for fileScanner.Scan() {
		line := fileScanner.Text()

		digits := ""
		// Loop through each character
		for i := 0; i < len(line); i++ {
			// Check if the character is a number
			if num, err := strconv.Atoi(string(line[i])); err == nil {
				// If so, add it to line's digits
				digits += strconv.Itoa(num)
			}
		}

		// Grab the first and last digits (bytes), convert to string, then to int - crazy
		first, _ := strconv.Atoi(string(digits[0]))
		last, _ := strconv.Atoi(string(digits[len(digits)-1]))

		// Concatenate the two ints into a string
		combined := strconv.Itoa(first) + strconv.Itoa(last)

		// Convert the combined string to an int
		combinedInt, _ := strconv.Atoi(combined)

		// Add it to total
		total += combinedInt
	}

	fmt.Println(total)
}