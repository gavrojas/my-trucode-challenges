package algorithms

// Given an array of numbers, write an algorithm to find out the products of every number,
// except from the one at the current index.
// Input:[1, 3, 9, 4]
// Output:[108, 36, 12, 27]

func Products(nums ...int) []int {
	result := []int{}
	for i := 0; i < len(nums); i++ {
		product := 1
		for j := 0; j < len(nums); j++ {
			if j != i {
				product = product * nums[j]
			}
		}
		result = append(result, product)
	}
	return result
}
