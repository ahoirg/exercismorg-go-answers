// Package lasagna is the solution to
// https://exercism.org/tracks/go/exercises/lasagna
package lasagna

const OvenTime = 40

// RemainingOvenTime returns the remaining minutes for the lasagna to be cooked.
func RemainingOvenTime(input int) int {
	return OvenTime - input
}

// PreparationTime returns how many minutes you spent preparing the lasagna
func PreparationTime(input int) int {
	return input * 2
}

// ElapsedTime returns the elapsed working time in minutes
func ElapsedTime(layers, min int) int {
	return PreparationTime(layers) + min
}
