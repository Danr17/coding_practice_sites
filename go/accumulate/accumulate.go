package accumulate

func Accumulate(listOfStrings []string, f func(string) string) []string {
	newList := make([]string, len(listOfStrings))
	for i, s := range listOfStrings {
		newList[i] = f(s)

	}
	return newList
}
