package algorithms

// Imagine you have two arrays which have already been sorted.
// Write an algorithm to merge the two arrays into a single array,
// which should also be sorted.
// Input:[1, 3, 6, 8, 11] and [2, 3, 5, 8, 9, 10]
// Output:[1, 2, 3, 3, 5, 6, 8, 8, 9, 10, 11]

func Merge(list1 []int, list2 []int) []int {
	completedList := append(list1, list2...)
	return mergeSort(completedList)
}

func mergeSort(list []int) []int {
	if len(list) <= 1 {
		return list
	}
	mid := len(list) / 2
	left := mergeSort(list[:mid]) // <-- el mid en left es no inclusivo
	right := mergeSort(list[mid:])
	return mergeLeftRight(left, right)
}

func mergeLeftRight(left, right []int) []int {
	merged := []int{}
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			merged = append(merged, left[i])
			i++
		} else {
			merged = append(merged, right[j])
			j++
		}
	}

	merged = append(merged, left[i:]...)
	merged = append(merged, right[j:]...)
	return merged
}
