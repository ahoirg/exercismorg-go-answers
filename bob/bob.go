//Package regexp is the solution to https://exercism.io/my/solutions/f0b0cb464131434bb3527654f9fdad27
package bob

import (
	"strings"
)

// Hey return bob's answer
func Hey(input string) string {
	input = strings.TrimSpace(input)

	switch {
	case input == "":
		return "Fine. Be that way!"
	case IsQuestion(input):
		if IsCapialized(input) {
			return "Calm down, I know what I'm doing!"
		}
		return "Sure."
	case IsCapialized(input):
		return "Whoa, chill out!"
	default:
		return "Whatever."
	}
}

func IsQuestion(input string) bool {
	questionMarkIndex := strings.Index(input, "?")
	return questionMarkIndex == len(input)-1
}

func IsCapialized(input string) bool {
	if strings.Compare(strings.ToUpper(input), input) == 0 {
		for _, r := range input {
			if (r > 'A' && r < 'Z') || r == 'A' || r == 'Z' {
				return true
			}
		}
	}

	return false
}
