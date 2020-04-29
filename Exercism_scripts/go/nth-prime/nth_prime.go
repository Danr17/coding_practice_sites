package prime

import "math"

//Nth list the nth prime number
func Nth(n int) (int, bool) {

	idx := 0
	i := 1

	if n == 1 {
		return 2, true
	}
	if n == 2 {
		return 3, true
	}
	if n == 0 {
		return 0, false
	}

	for {
		if IsPrimeSqrt(i) {
			idx++
		}
		if idx > n {
			break
		}
		i++
	}
	return i, true
}

//IsPrimeSqrt determine if the number is prime
func IsPrimeSqrt(value int) bool {
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(value)))); i++ {
		if value%i == 0 {
			return false
		}
	}
	return true
}
