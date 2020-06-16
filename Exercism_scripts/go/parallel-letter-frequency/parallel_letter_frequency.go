package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

//ConcurrentFrequency count the frequency of letters in texts using parallel computation
func ConcurrentFrequency(blob []string) FreqMap {
	results := make(chan FreqMap, 10)
	finalResult := FreqMap{}

	for _, text := range blob {
		go func(t string) {
			results <- Frequency(t)
		}(text)
	}

	for range blob {
		temp := <-results
		for index, value := range temp {
			finalResult[index] += value
		}

	}

	return finalResult
}
