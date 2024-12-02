package main

import (
	"fmt"
	"time"

	"trucode.com/challenge-w4/algorithms"
)

func main() {
	// ?Implementación 01.books
	start := time.Now()
	algorithms.RunSort()
	duration := time.Since(start)
	fmt.Println("Time: ", duration)
	// ?Implementación 02-10 algorithms
	fmt.Println("\n2. Palindrome -> ", algorithms.IsPalindrome("sugus"))
	fmt.Println("\n3. Decode chain -> ", algorithms.Decode("bits case debts cast air bond age"))
	fmt.Println("\n4. Product of numbers -> ", algorithms.Products(1, 2, 3))
	fmt.Println("\n5. Filter numbers lower than -> ", algorithms.FilterLowerThan(5, []int{1, 2, 5, 6}))
	fmt.Println("\n6. Shuffle arrays -> ", algorithms.Shuffle([]int{1, 2, 3, 4, 5}))
	fmt.Println("\n7. String Rotation -> ", algorithms.StrRotation("amazon", "zonama"))
	fmt.Println("\n8. Merge arrays-> ", algorithms.Merge([]int{1, 5, 8}, []int{2, 3, 4, 9}))
	fmt.Println("\n9. Zero Rows Columns -> ", algorithms.ZeroRowsColumns([][]int{
		{1, 0, 1, 1, 0},
		{0, 1, 1, 1, 0},
		{1, 1, 1, 1, 1},
		{1, 0, 1, 1, 1},
		{1, 1, 1, 1, 1},
	}))
	man := algorithms.Person{Name: "Man", Gender: algorithms.MAN}
	woman := algorithms.Person{Name: "Woman", Gender: algorithms.WOMAN}
	fmt.Println("\n10. Dance pairing -> ", algorithms.SquareDancePairing([]algorithms.Person{man, woman}))
}
