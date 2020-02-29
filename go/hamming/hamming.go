package hamming

import (
	"errors"
	"unicode/utf8"
)

// Distance function calculates the Hamming distance between two strings of equal length: a,b
// It will return either the distance, or error in case of a,b not having equal length.
func Distance(aStr, bStr string) (int, error) {
	if utf8.RuneCountInString(aStr) != utf8.RuneCountInString(bStr) {
		return 0, errors.New("a,b must have same length (rune count in string) to calculate Hamming distance")
	}

	distance := 0
	for _, a := range aStr {
		b, n := utf8.DecodeRuneInString(bStr)
		if a != b {
			distance++
		}
		bStr = bStr[n:]
	}

	return distance, nil
}
