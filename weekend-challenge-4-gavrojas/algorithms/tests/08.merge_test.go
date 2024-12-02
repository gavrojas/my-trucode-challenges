package algorithms_test

import (
	"reflect"
	"testing"

	"trucode.com/challenge-w4/algorithms"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name     string
		arr1     []int
		arr2     []int
		expected []int
	}{
		{
			name:     "Basic merge",
			arr1:     []int{1, 3, 5},
			arr2:     []int{2, 4, 6},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "Merge with duplicates",
			arr1:     []int{1, 2, 3, 4},
			arr2:     []int{2, 3, 4, 5},
			expected: []int{1, 2, 2, 3, 3, 4, 4, 5},
		},
		{
			name:     "Merge with empty array",
			arr1:     []int{1, 2, 3},
			arr2:     []int{},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Merge two empty arrays",
			arr1:     []int{},
			arr2:     []int{},
			expected: []int{},
		},
		{
			name:     "Merge with negative numbers",
			arr1:     []int{-3, -1, 0, 2},
			arr2:     []int{-2, 1, 3},
			expected: []int{-3, -2, -1, 0, 1, 2, 3},
		},
		{
			name:     "Merge with one array fully contained in the other",
			arr1:     []int{1, 2, 3, 4, 5},
			arr2:     []int{3},
			expected: []int{1, 2, 3, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := algorithms.Merge(tt.arr1, tt.arr2)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Merge(%v, %v) = %v, want %v", tt.arr1, tt.arr2, result, tt.expected)
			}
		})
	}
}