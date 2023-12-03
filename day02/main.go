package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const maxred = 12
const maxgreen = 13
const maxblue = 14

func main() {
	// Open the input file
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}

	// Create scanner object to read in lines of file
	fileScanner := bufio.NewScanner(readFile)

	total := 0
	power := 0
	
	cubes := make(map[string]int)

	// Iterate over each line
	for fileScanner.Scan() {
		line := fileScanner.Text()

		cubes["red"] = 0
		cubes["green"] = 0
		cubes["blue"] = 0

		// Split the line to get the game number
		game, _ := strconv.Atoi(strings.Split(strings.Split(line, ": ")[0], " ")[1])
		fmt.Printf("#%d ", game)

		turns := strings.Split(strings.Split(line, ": ")[1], "; ")
		// Iterate over each turn in the game, separated by ;
		for _, t := range turns {
			types := strings.Split(t, ", ")

			// Iterate over each type revealed in that turn, separated by ,
			for _, y := range types {
				// Split the type string into the number and the colour eg "1 red"
				number, _ := strconv.Atoi(strings.Split(y, " ")[0])
				colour := strings.Split(y, " ")[1]

				// fmt.Printf("%d %s \n", number, colour)

				// If this is the highest number we've seen for that colour, store the max
				if cubes[colour] < number {
					cubes[colour] = number
				}
			}
		}

		// Add the required "power" of this game to the total power.
		// Power is the product of multiplying the maximum number of cubes that would've been required

		power += cubes["red"] * cubes["green"] * cubes["blue"]

		fmt.Printf("red: %d, green: %d, blue: %d - ", cubes["red"], cubes["green"], cubes["blue"])
		// See if game is possible given our maxes

		if (cubes["red"] > maxred) || (cubes["green"] > maxgreen) || (cubes["blue"] > maxblue) {
			fmt.Println("impossible")
			continue
		} else {
			fmt.Println("possible")
			total += game
		}
	}

	fmt.Printf("Total sub of possible games: %d\n", total)
	fmt.Printf("Total power: %d\n", power)
}