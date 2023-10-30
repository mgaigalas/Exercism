// Package triangle is a small library containing tools for determining if type of triangle.
package triangle

import "slices"

type Kind int

const (
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

// KindFromSides returns type of triangle based on provided data about its sides
func KindFromSides(a, b, c float64) Kind {
	switch {
	case 2*slices.Max([]float64{a, b, c}) >= a+b+c:
		return NaT
	case a == b && b == c:
		return Equ
	case a == b || b == c || a == c:
		return Iso
	default:
		return Sca
	}
}
