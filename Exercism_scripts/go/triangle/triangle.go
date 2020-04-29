//package triangle identify the type of traingles
package triangle

import (
	"math"
)

type Kind int

const (
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

// KindFromSides identify the traingle type
func KindFromSides(a, b, c float64) Kind {

	if a <= 0 || b <= 0 || c <= 0 || math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) || math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		return NaT
	}
	if (a > b+c) || (b > a+c) || (c > a+b) {
		return NaT
	}
	if (a == b) && (b == c) {
		return Equ
	}
	if (a == b) || (b == c) || (a == c) {
		return Iso
	}
	return Sca
}
