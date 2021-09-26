// Package reverse is the solution to
// https://exercism.io/my/solutions/82806a3ca4bf4ae69a09f5a2b244954e
package reverse

// Reverse returns reverses the input.
func Reverse(input string) (result string) {
	tempInput := []rune(input)
	for i := len(tempInput) - 1; i >= 0; i-- {
		result += string(tempInput[i])
	}
	return
}
