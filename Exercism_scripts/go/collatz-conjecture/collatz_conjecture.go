package collatzconjecture

import "fmt"
import "errors"

func CollatzConjecture(input int) (int, error) {
	var steps int

	if input <= 0 {
		return 0, errors.New("The number has to be >= 0")
	}

	for input != 1 {
		fmt.Println(input)
		if input%2 == 0 {
			input = input / 2
			steps++
		} else {
			input = (input * 3) + 1
			steps++
		}
	}
	return steps, nil
}
