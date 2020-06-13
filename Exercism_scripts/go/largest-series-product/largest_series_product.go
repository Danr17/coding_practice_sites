package lsproduct

import (
	"errors"
	"strconv"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	var big int64 = 0
	if span > len(digits) || span < 0 {
		return 0, errors.New("wrong inputs")
	}
	for n := 0; n < len(digits)-span+1; n++ {
		var temp int64 = 1
		for i := 0; i < span; i++ {
			t, err := strconv.Atoi(string(digits[n+i]))
			if err != nil {
				return 0, err
			}
			temp *= int64(t)
		}
		if big < temp {
			big = temp
		}

	}
	return big, nil
}
