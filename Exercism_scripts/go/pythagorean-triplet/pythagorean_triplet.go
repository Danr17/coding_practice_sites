package pythagorean

import "math"

type Triplet [3]int

//Range finds Pythagorean triplet in a range
func Range(min, max int) []Triplet {

	var returnTriplet []Triplet
	floatMin := float64(min)
	floatMax := float64(max)

	for x := floatMin; x < floatMax-1; x++ {
		for y := x + 1; y < floatMax; y++ {
			z := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
			if z != math.Trunc(z) {
				continue
			}
			if z <= floatMax && z > y {
				returnTriplet = append(returnTriplet, Triplet{int(x), int(y), int(z)})
				continue
			}
			break
		}
	}

	return returnTriplet
}

//Sum returns all triplets that sum to sum
func Sum(sum int) []Triplet {
	var sumTriplets []Triplet
	triplets := Range(1, sum)

	for _, triplet := range triplets {
		if triplet[0]+triplet[1]+triplet[2] == sum {
			sumTriplets = append(sumTriplets, triplet)
		}
	}

	return sumTriplets
}
