package main

import "fmt"

// Soal 4: Best stock profit - return buy price for max profit
func bestStockProfit(prices []int) int {
	if len(prices) < 2 {
		return 0 // Tidak bisa beli dan jual jika kurang dari 2 harga
	}
	minPrice := prices[0] // Harga terendah saat ini
	maxProfit := 0        // Keuntungan maksimal saat ini
	buyPrice := prices[0] // Harga beli untuk keuntungan maksimal
	for _, price := range prices[1:] {
		// Hitung keuntungan jika jual di harga ini
		currentProfit := price - minPrice
		if currentProfit > maxProfit {
			maxProfit = currentProfit
			buyPrice = minPrice // Update harga beli
		}
		// Update harga terendah jika harga ini lebih rendah
		if price < minPrice {
			minPrice = price
		}
	}
	return buyPrice
}

func main() {
	// Soal 4
	fmt.Println("\nSoal 4:")
	fmt.Println("contoh Input: [10, 9, 6, 5, 15]", "Output: ", bestStockProfit([]int{10, 9, 6, 5, 15}))
	fmt.Println("input : [7, 8, 3, 10, 8]", "Output: ", bestStockProfit([]int{7, 8, 3, 10, 8}))
	fmt.Println("input : [5, 12, 11, 12, 10]", "Output: ", bestStockProfit([]int{5, 12, 11, 12, 10}))
	fmt.Println("input : [7, 18, 27, 10, 29]", "Output: ", bestStockProfit([]int{7, 18, 27, 10, 29}))
	fmt.Println("input : [20, 17, 15, 14, 10]", "Output: ", bestStockProfit([]int{20, 17, 15, 14, 10}))
	fmt.Println("input : [20, 17, 15, 14, 10]", "Output: ", bestStockProfit([]int{20, 17, 15, 14, 10}))

}
