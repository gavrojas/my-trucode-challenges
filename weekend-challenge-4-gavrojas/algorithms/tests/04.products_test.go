package algorithms_test

import (
	"reflect"
	"testing"

	"trucode.com/challenge-w4/algorithms"
)

func TestProducts(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Basic example from comment",
			input:    []int{1, 3, 9, 4},
			expected: []int{108, 36, 12, 27},
		},
		{
			name:     "Array with zeros",
			input:    []int{0, 1, 2, 3},
			expected: []int{6, 0, 0, 0},
		},
		{
			name:     "Array with negative numbers",
			input:    []int{-1, 2, -3, 4},
			expected: []int{-24, 12, -8, 6},
		},
		{
			name:     "Two element array",
			input:    []int{3, 4},
			expected: []int{4, 3},
		},
		{
			name:     "Empty array",
			input:    []int{},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := algorithms.Products(tt.input...)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Products(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}