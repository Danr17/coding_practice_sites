package darts

import "math"

//Score for Dart game
func Score(x float64, y float64) int {
	//https://en.wikipedia.org/wiki/Pythagorean_theorem
	z := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))

	switch {
	case z <= 1:
		return 10
	case z <= 5:
		return 5
	case z <= 10:
		return 1
	default:
		return 0
	}

}
