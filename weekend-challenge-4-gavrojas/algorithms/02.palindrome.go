package algorithms

import (
	"strings"
)

// A palindrome is a word, phrase, or number that is spelled the same forward and backward. For example, “dad” is a palindrome; “A man, a plan, a canal: Panama” is a palindrome if you take out the spaces and ignore the punctuation; and 1,001 is a numeric palindrome. We can use a stack to determine whether or not a given string is a palindrome.
// Write a function that takes a string of letters and returns true or false to determine whether it is palindromic.
// Input: amazon
// Output: false
//
// Input: dad
// Output: true

func IsPalindrome(text string) bool {
	text = strings.ToLower(strings.ReplaceAll(text, " ", "")) //<-- Quitar espacios y convertir todo a minúsculas
	var cont int                                              // <-- para contar aciertos

	for i := 0; i < len(text); i++ {
		if text[i] == text[len(text)-1-i] {
			cont++
		}
	}
	if cont == len(text) {
		return true
	} else {
		return false
	}
}
