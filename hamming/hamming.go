package hamming

import (
	"fmt"
)

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("Can't determine Hamming Distance for sequences of different length (a:%d, b:%d)!", len(a), len(b))
	}

	difference := 0
	for indexA, value := range a {
		if value != rune(b[indexA]) {
			difference++
		}
	}

	return difference, nil
}
