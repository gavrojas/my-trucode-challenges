package algorithms

// Imagine you have an array of numbers. Write an algorithm to remove all numbers
// less than five from the array.
// Input: '[10,4,5,8,2,9]'
// Output: '[10,5,8,9]'

func FilterLowerThan(limit int, input []int) []int {
	newArray := []int{}
	for i := 0; i < len(input); i++ {
		if input[i] >= limit {
			newArray = append(newArray, input[i])
		}
	}
	return newArray
}
