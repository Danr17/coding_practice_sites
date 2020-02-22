package listops

//IntList type define a slice of ints
type IntList []int

type unaryFunc func(int) int
type binFunc func(int, int) int
type predFunc func(int) bool

//Length meassure the length of the list
func (l IntList) Length() int {

	len := 0
	for range l {
		len++
	}

	return len
}

//Append extends the curent list with the one provided
func (l IntList) Append(appendThis IntList) IntList {
	var outputList IntList = make(IntList, l.Length()+appendThis.Length())

	copy(outputList, l)
	copy(outputList[l.Length():], appendThis)

	return outputList
}

//Concat given a series of lists, combine all items in all lists into one flattened list)
func (l IntList) Concat(args []IntList) IntList {
	for _, arg := range args {
		l = l.Append(arg)
	}
	return l
}

//Reverse given a list, return a list with all the original items, but in reversed order
func (l IntList) Reverse() IntList {
	for i, j := 0, l.Length()-1; i < j; i, j = i+1, j-1 {
		l[i], l[j] = l[j], l[i]
	}
	return l
}

//Foldr given a function, a list, and an initial accumulator, fold (reduce) each item into the accumulator from the right
func (l IntList) Foldr(fn binFunc, initial int) int {
	acc := initial
	for index := l.Length() - 1; index >= 0; index-- {
		acc = fn(l[index], acc)
	}
	return acc
}

//Foldl given a function, a list, and initial accumulator, fold (reduce) each item into the accumulator from the left
func (l IntList) Foldl(fn binFunc, initial int) int {
	acc := initial
	for _, item := range l {
		acc = fn(acc, item)
	}
	return acc
}

//Filter given a predicate and a list, return the list of all items for which predicate is True
func (l IntList) Filter(fn predFunc) IntList {
	outputList := IntList{}
	for _, item := range l {
		if fn(item) {
			outputList = outputList.Append(IntList{item})
		}
	}
	return outputList
}

//Map given a function and a list, return the list of the results of applying function(item) on all items
func (l IntList) Map(fn unaryFunc) IntList {
	for i, item := range l {
		l[i] = fn(item)
	}
	return l
}
