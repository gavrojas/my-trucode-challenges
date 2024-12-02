package algorithms_test

import (
	"testing"

	"trucode.com/challenge-w4/algorithms"
)

func TestStrRotation(t *testing.T) {
	tests := []struct {
		name     string
		str1     string
		str2     string
		expected bool
	}{
		{
			name:     "Example 1 from comment",
			str1:     "amazon",
			str2:     "azonma",
			expected: false,
		},
		{
			name:     "Example 2 from comment",
			str1:     "amazon",
			str2:     "azonam",
			expected: true,
		},
		{
			name:     "Same string",
			str1:     "hello",
			str2:     "hello",
			expected: true,
		},
		{
			name:     "Different length strings",
			str1:     "rotation",
			str2:     "rotations",
			expected: false,
		},
		{
			name:     "Empty strings",
			str1:     "",
			str2:     "",
			expected: true,
		},
		{
			name:     "Single character strings",
			str1:     "a",
			str2:     "a",
			expected: true,
		},
		{
			name:     "Rotation with repeated characters",
			str1:     "waterbottle",
			str2:     "erbottlewat",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := algorithms.StrRotation(tt.str1, tt.str2)
			if result != tt.expected {
				t.Errorf("StrRotation(%q, %q) = %v, want %v", tt.str1, tt.str2, result, tt.expected)
			}
		})
	}
}