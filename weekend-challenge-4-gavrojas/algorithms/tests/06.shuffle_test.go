package algorithms_test

import (
	"reflect"
	"sort"
	"testing"

	"trucode.com/challenge-w4/algorithms"
)

func TestShuffle(t *testing.T) {
	tests := []struct {
		name  string
		input []int
	}{
		{
			name:  "Example from comment",
			input: []int{1, 2, 3, 4, 5},
		},
		{
			name:  "Empty slice",
			input: []int{},
		},
		{
			name:  "Single element",
			input: []int{1},
		},
		{
			name:  "Two elements",
			input: []int{1, 2},
		},
		{
			name:  "Larger slice",
			input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name:  "Slice with duplicates",
			input: []int{1, 2, 2, 3, 3, 3, 4, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := make([]int, len(tt.input))
			copy(original, tt.input)

			result := algorithms.Shuffle(tt.input)

			// Check if the result has the same length as the input
			if len(result) != len(tt.input) {
				t.Errorf("Shuffle(%v) returned slice of length %d, want %d", tt.input, len(result), len(tt.input))
			}

			// Check if the result contains the same elements as the input
			sort.Ints(result)
			sort.Ints(original)
			if !reflect.DeepEqual(result, original) {
				t.Errorf("Shuffle(%v) returned slice with different elements: %v", tt.input, result)
			}

			// Check if the result is different from the input (for slices with more than one element)
			if len(tt.input) > 1 && reflect.DeepEqual(result, tt.input) {
				t.Errorf("Shuffle(%v) returned the same slice, expected it to be shuffled", tt.input)
			}
		})
	}
}