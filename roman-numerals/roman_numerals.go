// Package romannumerals is the solution to
// https://exercism.io/my/solutions/a04e37bfc19d441c9bf71ec9bf7c2fe3
package romannumerals

import (
	"fmt"
	"sort"
	"strings"
)

var romanNumerals = map[int]string{
	1000: "M",
	900:  "CM",
	500:  "D",
	400:  "CD",
	100:  "C",
	90:   "XC",
	50:   "L",
	40:   "XL",
	10:   "X",
	9:    "IX",
	5:    "V",
	4:    "IV",
	1:    "I",
}

// ToRomanNumeral converts integers to romanNumeral
func ToRomanNumeral(value int) (result string, err error) {
	var numbers []int

	for k := range romanNumerals {
		numbers = append(numbers, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))

	if value > 0 && value <= 3000 {
		for _, number := range numbers {
			dividend := value / number
			if dividend != 0 {
				result = result + strings.Repeat(romanNumerals[number], dividend)
				value = value - (dividend * number)
				if value == 0 {
					break
				}
			}

		}
	} else {
		return result, fmt.Errorf("value must greater than 0")
	}
	return result, nil
}
