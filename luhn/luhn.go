// Package luhn is the solution to
// https://exercism.io/my/solutions/c32e18f7eada43ab9dfdb5681bfc7336
package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

// Valid checks if the number complies with the luhm rule.
func Valid(n string) bool {
	n = strings.ReplaceAll(n, " ", "")
	len := len(n)
	if len <= 1 {
		return false
	}

	sum := 0

	// Reversing the number consumes unnecessary cpu and ram.
	// It is checked whether the entered value is even or odd.
	// The starting point is chosen accordingly.
	// So it's the same as starting from the far right.
	pass := true
	if len%2 == 0 {
		pass = false
	}

	for _, r := range n {
		if unicode.IsLetter(r) {
			return false
		}

		number, _ := strconv.Atoi(string(r))
		if !pass {
			number = number * 2
			if number > 9 {
				number -= 9
			}
		}

		sum += number
		pass = !pass
	}

	return sum%10 == 0
}
