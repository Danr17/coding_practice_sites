package letter

import "sync"

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
	var wg sync.WaitGroup
	var lock sync.Mutex
	results := make(chan rune)
	finalResult := FreqMap{}

	wg.Add(len(blob))

	for _, text := range blob {
		go func(t string) {
			defer wg.Done()
			for _, r := range t {
				results <- r
			}
		}(text)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		lock.Lock()
		if _, ok := finalResult[result]; !ok {
			finalResult[result] = 1
		} else {
			finalResult[result]++
		}
		lock.Unlock()
	}

	return finalResult
}
