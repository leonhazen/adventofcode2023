package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"unicode/utf8"
)

func numToInt(num string) int {
	// Try to convert string to int
	// If it fails, try to convert word to int

	numInt, err := strconv.Atoi(num)
	if err == nil {
		return numInt
	}

	switch num {
	case "one" :
		return 1
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	default:
		return -1
	}
}

// stolen straight from https://stackoverflow.com/a/34521190
func reverse(s string) string {
    size := len(s)
    buf := make([]byte, size)
    for start := 0; start < size; {
        r, n := utf8.DecodeRuneInString(s[start:])
        start += n
        utf8.EncodeRune(buf[size-start:], r)
    }
    return string(buf)
}

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
		enil := reverse(line)
		srebmun := rr.FindAllString(enil, -1)
		
		// Grab the first and last number, convert them to ints

		first := numToInt(numbers[0])
		last := numToInt(reverse(srebmun[0])) // don't forget to reverse it back

		// Concat them together

		combined := fmt.Sprintf("%d%d", first, last)

		// Convert the combined string to an int
		combinedInt, _ := strconv.Atoi(combined)

		// Add it to total
		total += combinedInt
	}

	fmt.Println(total)
}