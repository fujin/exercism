package scrabble

import (
	"strings"
)

// Score computes the basic scrabble score for given input string.
func Score(input string) (score int) {
	input = strings.ToUpper(input)

	for _, c := range input {
		switch c {
		case 'A', 'E', 'I', 'O', 'U', 'N', 'L', 'R', 'S', 'T':
			score += 1
		case 'D', 'G':
			score += 2
		case 'B', 'C', 'M', 'P':
			score += 3
		case 'F', 'H', 'V', 'W', 'Y':
			score += 4
		case 'K':
			score += 5
		case 'J', 'X':
			score += 8
		case 'Q', 'Z':
			score += 10
		}
	}

	return score
}
