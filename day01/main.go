package main

import (
	"adventofcode2024/common"
	"bufio"
	"fmt"
	"os"
	"regexp"
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

	r, _ := regexp.Compile(`(\d|one|two|three|four|five|six|seven|eight|nine)`)
	rr,  _ := regexp.Compile(`(\d|eno|owt|eerht|ruof|evif|xis|neves|thgie|enin)`)

	total := 0
	// Iterate over each line
	for fileScanner.Scan() {
		line := fileScanner.Text()

		// Find all digits in the line
		numbers := r.FindAllString(line, -1)

		// Lets do it backwards now because apparently overlapping word numbers
		// count as two numbers, eg "twone" would be 2 and 1
		enil := common.ReverseString(line)
		srebmun := rr.FindAllString(enil, -1)
		
		// Grab the first and last number, convert them to ints

		first := common.NumToInt(numbers[0])
		last := common.NumToInt(common.ReverseString(srebmun[0])) // don't forget to reverse it back

		// Concat them together

		combined := fmt.Sprintf("%d%d", first, last)

		// Convert the combined string to an int
		combinedInt, _ := strconv.Atoi(combined)

		// Add it to total
		total += combinedInt
	}

	fmt.Println(total)
}