package algorithms_test

import (
	"testing"

	"trucode.com/challenge-w4/algorithms"
)

func TestDecode(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Single word - 'cycle'",
			input:    "cycle",
			expected: "l",
		},
		{
			name:     "Single word - 'mouse'",
			input:    "mouse",
			expected: " ",
		},
		{
			name:     "Full message",
			input:    "craft block argon meter bells brown croon droop",
			expected: "for loop",
		},
		{
			name:     "Word starting with 'a'",
			input:    "apple",
			expected: "p",
		},
		{
			name:     "Word starting with 'b'",
			input:    "beach",
			expected: "a",
		},
		{
			name:     "Word starting with 'd'",
			input:    "dream",
			expected: "m",
		},
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := algorithms.Decode(tt.input)
			if result != tt.expected {
				t.Errorf("Decode(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}