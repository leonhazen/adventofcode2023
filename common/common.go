package common

import (
	"strconv"
	"unicode/utf8"
)


func NumToInt(num string) int {
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
func ReverseString(s string) string {
    size := len(s)
    buf := make([]byte, size)
    for start := 0; start < size; {
        r, n := utf8.DecodeRuneInString(s[start:])
        start += n
        utf8.EncodeRune(buf[size-start:], r)
    }
    return string(buf)
}


func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}