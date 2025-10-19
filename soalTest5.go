package main

import (
	"fmt"
	"strconv"
)

// Soal 5: Count numbers in string array
func countNumbers(arr []string) int {
	count := 0
	for _, s := range arr {
		if _, err := strconv.Atoi(s); err == nil {
			count++
		}
	}
	return count
}

func countString(arr []string) int {
	count := 0
	for _, s := range arr {
		if _, err := strconv.Atoi(s); err != nil {
			count++
		}
	}
	return count
}

func main() {
	// Soal 5
	fmt.Println("\nSoal 5:")
	fmt.Println("contoh Input : [\"2\", \"h\", \"6\", \"u\", \"y\", \"t\", \"7\", \"j\", \"y\", \"h\", \"8\"]", "Output: ", countNumbers([]string{"2", "h", "6", "u", "y", "t", "7", "j", "y", "h", "8"}))
	fmt.Println("contoh Input : [\"b\", \"7\", \"h\", \"6\", \"h\", \"k\", \"i\", \"5\", \"8\", \"7\", \"8\"]", "Output: ", countNumbers([]string{"b", "7", "h", "6", "h", "k", "i", "5", "8", "7", "8"}))
	fmt.Println("contoh Input : [\"7\", \"b\", \"8\", \"5\", \"6\", \"9\", \"n\", \"f\", \"y\", \"6\", \"9\"]", "Output: ", countNumbers([]string{"7", "b", "8", "5", "6", "9", "n", "f", "y", "6", "9"}))
	fmt.Println("contoh Input : [\"u\", \"h\", \"b\", \"n\", \"7\", \"6\", \"5\", \"1\", \"g\", \"7\", \"9\"]", "Output: ", countNumbers([]string{"u", "h", "b", "n", "7", "6", "5", "1", "g", "7", "9"}))

	fmt.Println("jika itu string yang dihitung:")
	fmt.Println("contoh Input : [\"2\", \"h\", \"6\", \"u\", \"y\", \"t\", \"7\", \"j\", \"y\", \"h\", \"8\"]", "Output: ", countString([]string{"2", "h", "6", "u", "y", "t", "7", "j", "y", "h", "8"}))
}
