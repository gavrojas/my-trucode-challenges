package algorithms_test

import (
	"reflect"
	"testing"

	"trucode.com/challenge-w4/algorithms"
)

func TestFilterLowerThan(t *testing.T) {
	tests := []struct {
		name     string
		limit    int
		input    []int
		expected []int
	}{
		{
			name:     "Example from comment",
			limit:    5,
			input:    []int{10, 4, 5, 8, 2, 9},
			expected: []int{10, 5, 8, 9},
		},
		{
			name:     "All elements below limit",
			limit:    10,
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{},
		},
		{
			name:     "All elements above limit",
			limit:    0,
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Empty input array",
			limit:    5,
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Negative numbers",
			limit:    0,
			input:    []int{-3, -2, -1, 0, 1, 2, 3},
			expected: []int{0, 1, 2, 3},
		},
		{
			name:     "Limit is negative",
			limit:    -2,
			input:    []int{-3, -2, -1, 0, 1, 2, 3},
			expected: []int{-2, -1, 0, 1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := algorithms.FilterLowerThan(tt.limit, tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("FilterLowerThan(%d, %v) = %v, want %v", tt.limit, tt.input, result, tt.expected)
			}
		})
	}
}