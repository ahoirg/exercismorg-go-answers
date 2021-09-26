package collatzconjecture

import "errors"

func CollatzConjecture(input int) (int, error) {
	if input < 1 {
		return 0, errors.New("value must be positive number")
	}

	counter := 0
	for input > 1 {
		if input%2 == 0 {
			input = input / 2
		} else {
			input = 3*input + 1
		}
		counter++
	}
	return counter, nil
}
