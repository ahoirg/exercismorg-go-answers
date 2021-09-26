// Package scrabble is the solution to
// https://exercism.io/my/solutions/bbc6c7c75ca54ced863eed4780b18800

package scrabble

import "strings"

// Score returns the points earned by the user for the word used.
func Score(input string) (score int) {

	input = strings.ToLower(input)
	for _, st := range input {
		switch st {
		case 'a', 'e', 'i', 'o', 'u', 'l', 'n', 'r', 's', 't':
			score += 1
		case 'd', 'g':
			score += 2
		case 'b', 'c', 'm', 'p':
			score += 3
		case 'f', 'h', 'v', 'w', 'y':
			score += 4
		case 'k':
			score += 5
		case 'j', 'x':
			score += 8
		case 'q', 'z':
			score += 10
		}
	}

	return
}
