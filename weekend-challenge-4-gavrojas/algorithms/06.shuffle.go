package algorithms

import (
	"math/rand"
)

// Write an algorithm to shuffle an array into a random order
// in-place (i.e. without creating a new array).
// The important thing here is to shuffle the array positions,
// do not worry if the result is a pseudo-random shuffle, avoid the use of rand.Shuffle
// Input:[1, 3, 4, 9]
// Output:[9, 3, 1, 4]

func Shuffle(input []int) []int {
	shuffled := make([]int, len(input))
	copy(shuffled, input)

	if len(input) > 1 {
		for i := len(input) - 1; i > 0; i-- {
			j := rand.Intn(i + 1)
			input[i], input[j] = input[j], input[i]
		}

		if isEqual(shuffled, input) {
			return Shuffle(input)
		}
	}
	return input
}

func isEqual(slice1, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}
