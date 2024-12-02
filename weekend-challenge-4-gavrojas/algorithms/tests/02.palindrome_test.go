package algorithms_test

import (
	"testing"

	"trucode.com/challenge-w4/algorithms"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Simple non-palindrome",
			input:    "amazon",
			expected: false,
		},
		{
			name:     "Simple palindrome",
			input:    "dad",
			expected: true,
		},
		{
			name:     "Longer palindrome",
			input:    "racecar",
			expected: true,
		},
		{
			name:     "Palindrome with spaces",
			input:    "a man a plan a canal panama",
			expected: true,
		},
		{
			name:     "Non-palindrome with spaces",
			input:    "hello world",
			expected: false,
		},
		{
			name:     "Palindrome with mixed case",
			input:    "Able was I ere I saw Elba",
			expected: true,
		},
		{
			name:     "Empty string",
			input:    "",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := algorithms.IsPalindrome(tt.input)
			if result != tt.expected {
				t.Errorf("IsPalindrome(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}