package algorithms

// Given two strings, str1 and str2, write a program that checks if str2 is a rotation of str1.
// Input: amazon, azonma
// Output: False
//
// Input: amazon, azonam
// Output: true

func StrRotation(str1 string, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	if str1 == str2 {
		return true
	}

	for i := len(str2) - 1; i > 0; i-- {
		right := str2[i:]
		left := str2[:i]
		chain := right + left
		if chain == str1 {
			return true
		}
	}
	return false
}
