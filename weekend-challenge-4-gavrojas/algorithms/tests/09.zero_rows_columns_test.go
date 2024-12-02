package algorithms_test

import (
	"reflect"
	"testing"

	"trucode.com/challenge-w4/algorithms"
)

func TestZeroRowsColumns(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]int
		expected [][]int
	}{
		{
			name: "3x3 matrix with a zero",
			input: [][]int{
				{1, 1, 1},
				{1, 0, 1},
				{1, 1, 1},
			},
			expected: [][]int{
				{1, 0, 1},
				{0, 0, 0},
				{1, 0, 1},
			},
		},
		{
			name: "4x4 matrix with multiple zeros",
			input: [][]int{
				{1, 1, 1, 1},
				{1, 0, 1, 1},
				{1, 1, 0, 1},
				{1, 1, 1, 1},
			},
			expected: [][]int{
				{1, 0, 0, 1},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{1, 0, 0, 1},
			},
		},
		{
			name: "2x2 matrix with no zeros",
			input: [][]int{
				{1, 1},
				{1, 1},
			},
			expected: [][]int{
				{1, 1},
				{1, 1},
			},
		},
		{
			name: "example from comments",
			input: [][]int{
				{1,0,1,1,0},
				{0,1,1,1,0},
				{1,1,1,1,1},
				{1,0,1,1,1},
				{1,1,1,1,1},
			},
			expected: [][]int{
				{0,0,0,0,0},
				{0,0,0,0,0},
				{0,0,1,1,0},
				{0,0,0,0,0},
				{0,0,1,1,0},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := algorithms.ZeroRowsColumns(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ZeroRowsColumns(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}