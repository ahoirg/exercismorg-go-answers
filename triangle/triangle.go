package triangle

import (
	"math"
	"sort"
)

type Kind int

const (
	NaT = iota // not a triangle
	Equ        // equilateral
	Iso        // isosceles
	Sca        // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	sides := [3]float64{a, b, c}
	sort.Float64s(sides[:])

	isNegativeOrNan := sides[0] <= 0 || math.IsNaN(sides[0])
	isInf := math.IsInf(sides[2], 0)
	isTriangle := sides[0]+sides[1] >= sides[2]

	if isNegativeOrNan || isInf || !isTriangle {
		return NaT
	}

	equilAB := a == b
	equilAC := a == c
	equilBC := b == c

	var k Kind = Sca
	if equilAB && equilAC {
		k = Equ
	} else if equilAB || equilAC || equilBC {
		k = Iso
	}

	return k
}
