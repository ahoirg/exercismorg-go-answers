// Package clock is the solution to
// https://exercism.org/tracks/go/exercises/clock
package clock

import (
	"fmt"
)

type Clock int

const (
	dayMinutes  = 1440
	hourMinutes = 60
)

// New creates a valid clock with required parameters.
func New(h, m int) Clock {
	return convert(h*hourMinutes + m)
}

// String converts Clock type variable to string.
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c/60, c%60)
}

// Subtract subtracts minutes from the current time.
func (c Clock) Subtract(input int) Clock {
	return convert(int(c) - input)
}

//Add adds minutes to the current time.
func (c Clock) Add(input int) Clock {
	return convert(int(c) + input)
}

// Convert is a private method where minute and hour conversions are made.
func convert(m int) Clock {
	if m < 0 {
		return Clock(dayMinutes + (m % dayMinutes))
	} else if m >= dayMinutes {
		return Clock(m % dayMinutes)
	}
	return Clock(m)
}
