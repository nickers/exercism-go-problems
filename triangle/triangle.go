package triangle

import "math"

const testVersion = 3

// KindFromSides return kind of triangles specified by sides of given length
func KindFromSides(a, b, c float64) Kind {
	lengthsArePositiveNumbers := a > 0 && b > 0 && c > 0 && !math.IsInf(a+b+c, 0)
	lengthEquationFulfilled := (a+b >= c && a+c >= b && b+c >= a)
	if !lengthEquationFulfilled || !lengthsArePositiveNumbers {
		return NaT
	}

	if a == b && b == c {
		return Equ
	}

	if a != b && b != c && a != c {
		return Sca
	}

	return Iso
}

// Kind specify kind of triangle
type Kind uint8

const (
	// NaT not a triangle
	NaT Kind = iota
	// Equ equilateral
	Equ
	// Iso isosceles
	Iso
	// Sca scalene
	Sca
)