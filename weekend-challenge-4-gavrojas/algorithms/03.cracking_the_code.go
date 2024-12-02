package algorithms

import (
	"fmt"
	"strings"
)

// A code has been invented which replaces each character in a sentence with a
// five letter word. The first letter of each encoded word determines which of the remaining
// four characters contains the decoded character according to this table:
// First letter	Character number
// a	            2
// b	            3
// c	            4
// d	            5
// So for example, the encoded word 'cycle' would be decoded to the character 'l'.
// This is because the first letter is a 'c', so you look for the fourth character, which is 'l'.
// If the first letter of the encoded word isn't 'a', 'b', 'c', or 'd'
// (for example 'mouse') this should be decoded to a space.
// Write a function called decode which takes an encoded word as an argument,
// and returns the correct decoded character.
// Use your function to decode the following message:
// 'craft block argon meter bells brown croon droop'.
// Input: 'craft block argon meter bells brown croon droop'
// Output: 'for loop'

type letter int

const (
	NOLETTER letter = iota
	a
	b
	c
	d
)

func Decode(encoded string) string {
	encodedWords := strings.Split(encoded, " ") //<-- cadena de texto a arreglo
	var chainText string

	for i := 0; i < len(encodedWords); i++ {
		if len(encodedWords[i]) == 5 {
			switch encodedWords[i][0] {
			case 'a':
				chainText = fmt.Sprintf("%s%c", chainText, encodedWords[i][a])
			case 'b':
				chainText = fmt.Sprintf("%s%c", chainText, encodedWords[i][b])
			case 'c':
				chainText = fmt.Sprintf("%s%c", chainText, encodedWords[i][c])
			case 'd':
				chainText = fmt.Sprintf("%s%c", chainText, encodedWords[i][d])
			default:
				chainText = fmt.Sprintf("%s ", chainText)
			}
		}
	}

	return chainText
}
