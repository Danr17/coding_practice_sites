package strain

//Ints type
type Ints []int

//Lists type
type Lists [][]int

//Strings type
type Strings []string

//Keep returns a new collection containing those elements where the predicate is true
func (source Ints) Keep(f func(int) bool) Ints {
	var outInts []int
	for _, i := range source {
		if f(i) {
			outInts = append(outInts, i)
		}
	}
	return outInts
}

//Discard returns a new collection containing those elements where the predicate is false
func (source Ints) Discard(f func(int) bool) Ints {
	var outInts []int
	for _, i := range source {
		if !f(i) {
			outInts = append(outInts, i)
		}
	}
	return outInts
}

//Keep returns a new collection containing those elements where the predicate is true
func (source Lists) Keep(f func([]int) bool) Lists {
	var outLists [][]int
	for _, i := range source {
		if f(i) {
			outLists = append(outLists, i)
		}
	}
	return outLists
}

//Keep returns a new collection containing those elements where the predicate is true
func (source Strings) Keep(f func(string) bool) Strings {
	var outStrings []string
	for _, i := range source {
		if f(i) {
			outStrings = append(outStrings, i)
		}
	}
	return outStrings
}
