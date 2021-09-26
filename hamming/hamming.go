// Package hamming is the solution to http://exercism.io/exercises/go/hamming
package hamming

import (
	"fmt"
	"unicode/utf8"
)

// Distance returns the hamming distance of two strings.
func Distance(a, b string) (counter int, err error) {
	lengthA := utf8.RuneCountInString(a)
	lengthB := utf8.RuneCountInString(b)

	if lengthA != lengthB {
		if lengthA < lengthB {
			err = fmt.Errorf("disallow second strand longer")
		} else {
			err = fmt.Errorf("disallow first strand longer")
		}

		return 0, err
	}

	for i := 0; i < lengthA; i++ {
		if a[i] != b[i] {
			counter++
		}
	}

	return counter, nil
}
