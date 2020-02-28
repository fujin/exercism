package hamming

import (
	"errors"
	"unicode/utf8"
)

// Distance function calculates the Hamming distance between two strings of equal length: a,b
// It will return either the distance, or error in case of a,b not having equal length.
func Distance(a, b string) (int, error) {
	if utf8.RuneCountInString(a) != utf8.RuneCountInString(b) {
		return 0, errors.New("a,b must have same length to calculate Hamming distance")
	}

	distance := 0
	// for i, w := 0; i <= len(a); i += w {
	// 	runeValue, width := utf8.DecodeRuneInString(a[i:])
	// 	fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
	// 	w = width
	// }

	for i := range a {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}
