// Package raindrops is the solution to http://exercism.io/exercises/go/raindrops
package raindrops

import (
	"strconv"
	"strings"
)

const (
	pling = "Pling"
	plang = "Plang"
	plong = "Plong"
)

// Convert checks the factor of the integer
// according to 3,5,7 and returns the result as string.
func Convert(input int) (result string) {
	var builder strings.Builder

	if input%3 == 0 {
		builder.WriteString(pling)
	}

	if input%5 == 0 {
		builder.WriteString(plang)
	}

	if input%7 == 0 {
		builder.WriteString(plong)
	}

	result = builder.String()
	if result == "" {
		result = strconv.FormatInt(int64(input), 10)
	}

	return
}
