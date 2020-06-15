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
	results := make(chan FreqMap)
	finalResult := FreqMap{}

	for _, text := range blob {
		go func(t string) {
			result := FreqMap{}
			for _, r := range t {
				if _, ok := result[r]; !ok {
					result[r] = 1
				} else {
					result[r]++
				}
			}
			results <- result
		}(text)
	}

	for i := 0; i < len(blob); i++ {
		temp := <-results
		for index, value := range temp {

			if _, ok := finalResult[index]; !ok {
				finalResult[index] = value
			} else {
				finalResult[index] += value
			}
		}

	}

	close(results)
	return finalResult
}
