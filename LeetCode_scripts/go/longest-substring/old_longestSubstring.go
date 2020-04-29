package lengthstring

/*
func lengthOfLongestSubstring(s string) int {

	var largestnum int

	for i := 0; i < len(s); i++ {
		num := checkFirstLength(s[i:])
		if num > largestnum {
			largestnum = num
		}
	}

	return largestnum
}

func checkFirstLength(s string) int {
	listSt := []string{s[0:1]}
	for i := 1; i < len(s); i++ {
		if stringInSlice(s[i:i+1], listSt) {
			break
		}
		listSt = append(listSt, s[i:i+1])
	}
	return len(listSt)
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
*/
