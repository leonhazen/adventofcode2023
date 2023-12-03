package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Part struct {
	partno string
	x, y int
	relevant bool
}

type Symbol struct {
	symbol string
	x, y int
}

func (p *Part) isRelevant() bool {
	return p.relevant
}

func main() {
	// Open the input file
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}

	lineno := 0
	total := 0

	// Create arrays to store parts and symbols
	// Each is a struct with x and y coordinates, y being the line number, x being the character index from 0
	var parts []Part
	var symbols []Symbol

	r, _ := regexp.Compile(`(\d+)`)
	r2, _ := regexp.Compile(`[^.\d]`)

	// Create scanner object to read in lines of file
	fileScanner := bufio.NewScanner(readFile)

	// Iterate over each line
	for fileScanner.Scan() {
		line := fileScanner.Text()

		// Find all part numbers in the line
		foundParts := r.FindAllString(line, -1)
		foundPartIndexes := r.FindAllStringIndex(line, -1)
		
		// Iterate over each part number and add to parts array
		for i, p := range foundParts {
			parts = append(parts, Part{p, foundPartIndexes[i][0], lineno, false})
		}

		// Find all symbols in the line
		foundSymbols := r2.FindAllStringIndex(line, -1)
		foundSymbolChars := r2.FindAllString(line, -1)

		// Iterate over each symbol and add to symbols array
		for i, symbol := range foundSymbols {
			symbols = append(symbols, Symbol{foundSymbolChars[i], symbol[0], lineno})
		}
		lineno++
	}

	// Iterate over each part to see if they're relevant
	for i, p := range parts {
		relevant := false

		// Check if surrounding x y coordinates are symbols
		minX := max(p.x - 1, 0)
		maxX := p.x + len(p.partno)
		minY := max(p.y - 1, 0)
		maxY := p.y + 1

		// Iterate over symbols and check if any are in the surrounding coordinates
		for _, s := range symbols {
			if s.x >= minX && s.x <= maxX && s.y >= minY && s.y <= maxY {
				relevant = true
			}
		}

		if relevant {
			parts[i].relevant = true
			partnoint, _ := strconv.Atoi(p.partno)
			// fmt.Printf("Adding relevant part %d to total: %d\n",partnoint,total)
			total += partnoint
		}
	}

	// fmt.Println(len(parts))
	// fmt.Println(len(symbols))
	fmt.Printf("Part 1 answer - total of relevant parts: %d\n",total)

	// Part 2
	// Find gear ratios. Gears are '*' symbols adjacent to exactly 2 parts
	// Iterate over each symbol and find gears

	total = 0
	for _, s := range symbols {
		if s.symbol == "*" {
			// Create slice of matching parts
			var adjparts []string

			// Check if there are 2 parts adjacent to the gear
			for _, p := range parts {
				minX := max(p.x - 1, 0)
				maxX := p.x + len(p.partno)
				minY := max(p.y - 1, 0)
				maxY := p.y + 1

				if s.x >= minX && s.x <= maxX && s.y >= minY && s.y <= maxY {
					adjparts = append(adjparts, p.partno)
				}
			}
			if len(adjparts) == 2 {
				part1, _ := strconv.Atoi(adjparts[0])
				part2, _ := strconv.Atoi(adjparts[1])
				total += part1 * part2
			}
		}
	}
	fmt.Printf("Part 2 answer - total of gear ratios: %d\n",total)
}