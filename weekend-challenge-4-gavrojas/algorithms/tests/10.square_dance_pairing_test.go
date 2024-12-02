package algorithms_test

import (
	"reflect"
	"testing"

	"trucode.com/challenge-w4/algorithms"
)

func TestSquareDancePairing(t *testing.T) {
	tests := []struct {
		name     string
		input    []algorithms.Person
		expected [][]algorithms.Person
	}{
		{
			name: "Example from comments",
			input: []algorithms.Person{
				algorithms.John, algorithms.Jane, algorithms.Serena, algorithms.Luna, algorithms.Darien, algorithms.Artemis, algorithms.Rei, algorithms.Nicolas,
			},
			expected: [][]algorithms.Person{
				{algorithms.John, algorithms.Jane},
				{algorithms.Serena, algorithms.Darien},
				{algorithms.Luna, algorithms.Artemis},
				{algorithms.Rei, algorithms.Nicolas},
			},
		},
		{
			name: "Uneven number of men and women",
			input: []algorithms.Person{
				algorithms.John, algorithms.Jane, algorithms.Serena, algorithms.Luna, algorithms.Darien, algorithms.Artemis, algorithms.Rei,
			},
			expected: [][]algorithms.Person{
				{algorithms.John, algorithms.Jane},
				{algorithms.Serena, algorithms.Darien},
				{algorithms.Luna, algorithms.Artemis},
			},
		},
		{
			name: "All men, then all women",
			input: []algorithms.Person{
				algorithms.John, algorithms.Darien, algorithms.Artemis, algorithms.Nicolas, algorithms.Jane, algorithms.Serena, algorithms.Luna, algorithms.Rei,
			},
			expected: [][]algorithms.Person{
				{algorithms.John, algorithms.Jane},
				{algorithms.Darien, algorithms.Serena},
				{algorithms.Artemis, algorithms.Luna},
				{algorithms.Nicolas, algorithms.Rei},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := algorithms.SquareDancePairing(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("SquareDancePairing(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}