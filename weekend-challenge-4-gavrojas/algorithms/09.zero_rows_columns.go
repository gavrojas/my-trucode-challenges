package algorithms

// Write an algorithm which searches through a 2D array, and whenever it finds a zero should
// set the entire row and column to zero
// Input:
// [[1,0,1,1,0],
// [0,1,1,1,0],
// [1,1,1,1,1],
// [1,0,1,1,1],
// [1,1,1,1,1]];
// Output:
// [[0,0,0,0,0],
// [0,0,0,0,0],
// [0,0,1,1,0],
// [0,0,0,0,0],
// [0,0,1,1,0]];

func ZeroRowsColumns(array2D [][]int) [][]int {
	copyArray2D := make([][]int, len(array2D))

	// Hacer una copia exacta del array2D
	for i := range array2D {
		copyArray2D[i] = make([]int, len(array2D[i]))
		copy(copyArray2D[i], array2D[i])
	}

	for i := 0; i < len(array2D); i++ {
		for j := 0; j < len(array2D); j++ {
			if copyArray2D[i][j] == 0 {
				array2D = zeroConversion(array2D, i, j)
			}
		}
	}
	return array2D
}

func zeroConversion(array [][]int, a, b int) [][]int {
	for i := 0; i < len(array); i++ {
		array[i][b] = 0
	}
	for j := 0; j < len(array); j++ {
		array[a][j] = 0
	}
	return array
}
