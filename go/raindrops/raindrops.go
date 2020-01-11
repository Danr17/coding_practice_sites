package raindrops

import "strconv"

//Convert check modulo and add Pling, Plnag , Plong if the case
func Convert(input int) string {
	var output string
	if input%3 == 0 {
		output += "Pling"
	}
	if input%5 == 0 {
		output += "Plang"
	}
	if input%7 == 0 {
		output += "Plong"
	}

	if output == "" {
		return strconv.Itoa(input)
	}
	return output

}
