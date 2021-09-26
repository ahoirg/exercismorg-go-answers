// Package diffsquares is the solution to
// https://exercism.io/my/solutions/bbd5e152da3b4265a6d1807cd3d596d5
package diffsquares

// SquareOfSum returns the square of the sum of the numbers up to the given number.
func SquareOfSum(input int) int {
	// This approach benchmarks on my machine at
	// BenchmarkSquareOfSum-8   	1000000000	         0.5121 ns/op	       0 B/op	       0 allocs/op

	// link for formule, https://en.wikipedia.org/wiki/Summation
	sum := (input * (input + 1)) / 2
	return sum * sum

}

// SumOfSquares returns the sum of the squares of the numbers up to the given number.
func SumOfSquares(input int) int {
	// This approach benchmarks on my machine at
	// BenchmarkSumOfSquares-8   	1000000000	         0.4679 ns/op	       0 B/op	       0 allocs/op

	// link for formule, https://en.wikipedia.org/wiki/Summation
	return (input * (input + 1) * (2*input + 1)) / 6
}

// Difference returns difference of SquareOfSum and SumOfSquares.
func Difference(input int) int {
	// This approach benchmarks on my machine at
	// BenchmarkDifference-8   	1000000000	         0.4830 ns/op	       0 B/op	       0 allocs/op
	return SquareOfSum(input) - SumOfSquares(input)
}
