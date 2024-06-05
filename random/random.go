package random

import (
	"math/rand"
)

func CreateRandom(n int) []int {

	numbers := make([]int, n)

	for i := range numbers {
		numbers[i] = rand.Intn(100)
	}

	return numbers
}
