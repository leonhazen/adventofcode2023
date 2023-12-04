package main

import (
	"adventofcode2023/common"
	"bufio"
	"fmt"
	"os"
	"strings"
)


func main() {
	// Open the input file
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}

	// Create scanner object to read in lines of file
	fileScanner := bufio.NewScanner(readFile)

	//part 1 total
	total1 := 0

	//track which card we're playing
	card := 0

	// map of multipliers for part 2
	cardmultipliers := make(map[int]int)

	// Iterate over each line
	for fileScanner.Scan() {
		line := fileScanner.Text()
		card++
		if cardmultipliers[card] == 0 {
			cardmultipliers[card] = 1
		}
		
		fmt.Printf("#%d (x%d plays) ", card, cardmultipliers[card])

		cards := strings.Split(strings.Split(line, ": ")[1], " | ")

		var winners []string = strings.Split(cards[0], " ")
		var numbers []string = strings.Split(cards[1], " ")

		reward := 0
		matches := 0
		// fmt.Printf(" %s | %s | ", winners, numbers)
		for i := 0; i < len(numbers); i++ {
			if common.StringInSlice(numbers[i], winners) && numbers[i] != "" {
				// fmt.Printf("x")
				if reward == 0 {
					reward++
				} else {
					reward = reward * 2
				}

				//part 2
				matches++
			}
		}
		// fmt.Printf(" %d\n", reward)
		total1 += reward


		//part 2
		fmt.Printf("%dx winners\n", matches)
		for i := 0; i < matches; i++ {
			if cardmultipliers[card+1+i] == 0 {
				cardmultipliers[card+1+i] = 1
			}
			cardmultipliers[card+1+i] += 1 * cardmultipliers[card]
			
		}
	}
	fmt.Printf("\n%d\n", total1)

	//part 2 - count how many cards in total
	total2 := 0
	for _, v := range cardmultipliers {
		total2 += v
	}
	fmt.Printf("%d\n", total2)
}