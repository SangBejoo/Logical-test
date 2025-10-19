package main

import "fmt"

// Soal 3: Fibonacci
func fibonacci(n int) []int {
	if n <= 0 {
		return []int{}
	}
	seq := make([]int, n)
	if n > 0 {
		seq[0] = 0
	}
	if n > 1 {
		seq[1] = 1
	}
	for i := 2; i < n; i++ {
		seq[i] = seq[i-1] + seq[i-2]
	}
	return seq
}

func main() {
	// Soal 3
	fmt.Println("\nSoal 3:")
	fmt.Println(fibonacci(10)) // first 10 numbers
	fmt.Println(fibonacci(5))  // first 5 numbers
}
