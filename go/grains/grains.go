package grains

import (
	"errors"
	"math"
)

//Square calculate how many grains are on a given square
func Square(input int) (uint64, error) {

	if input < 1 || input > 64 {
		return 0, errors.New("The value is out of range, 1 to 64 are allowed")
	}

	return uint64(math.Pow(2, float64(input-1))), nil
}

//Total calculate the total number of grains on the chessboard
func Total() uint64 {
	var total uint64
	n := 64
	for n > 0 {
		val, err := Square(n)
		if err != nil {
			return 0
		}
		total += val
		n--
	}
	return total
}
