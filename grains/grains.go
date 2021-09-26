// Package grains is the solution to
// https://exercism.io/my/solutions/c32e18f7eada43ab9dfdb5681bfc7336
package grains

import (
	"fmt"
)

// Square returns the amount of wheat in a chess square.
func Square(n int) (result uint64, err error) {
	temp := 0
	switch {
	case n <= 0:
		return 0, fmt.Errorf("value can not be less than or equal 0")
	case n <= 8:
		result = 128
		temp = 8
	case n <= 16:
		result = 32768
		temp = 16
	case n <= 24:
		result = 8388608
		temp = 24
	case n <= 32:
		result = 2147483648
		temp = 32
	case n <= 40:
		result = 549755813888
		temp = 40
	case n <= 64:
		result = 9223372036854775808
		temp = 64
	default:
		return 0, fmt.Errorf("value can not be greater than 64")
	}

	for i := temp; i > n; i-- {
		result = result / 2
	}
	return result, nil
}

// Total returns the amount of wheat on the entire chessboard.
func Total() (result uint64) {
	for i := 1; i <= 64; i++ {
		temp, _ := Square(i)
		result += temp
	}
	return uint64(result)
}
