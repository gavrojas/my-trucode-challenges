package algorithms

import "fmt"

//Imagine that I gave you twenty books to sort in alphabetical order. How
//would you go about it? Can you express this as an algorithm?
func RunSort() {
	books := []string{
		"Cien años de soledad",
		"Don Quijote de la Mancha",
		"Orgullo y prejuicio",
		"1984",
		"Matar a un ruiseñor",
		"El gran Gatsby",
		"Crimen y castigo",
		"El principito",
		"La sombra del viento",
		"Los pilares de la tierra",
		"El nombre de la rosa",
		"Cumbres borrascosas",
		"El retrato de Dorian Gray",
		"Fahrenheit 451",
		"El viejo y el mar",
		"Ulises",
		"La casa de los espíritus",
		"Siddhartha",
		"La carretera",
		"El túnel",
		"La tregua",
	}
	fmt.Println("1. Sorting Books -> ", sortBooks(books))
}

//* I would use merge Arghorithm because it has a complexity constant of O(n log n), in the best or worst case.

func sortBooks(books []string) []string {
	if len(books) <= 1 {
		return books
	}
	mid := len(books) / 2
	left := sortBooks(books[:mid])
	right := sortBooks(books[mid:])
	return merge(left, right)
}

func merge(left, right []string) []string {
	merged := []string{}
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
