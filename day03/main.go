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

		// Iterate over each symbol and add to symbols array
		for _, symbol := range foundSymbols {
			symbols = append(symbols, Symbol{symbol[0], lineno})
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
	fmt.Println(total)
}