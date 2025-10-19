package main

import (
	"fmt"
	"strings"
)

// Soal 1:
// funssi untuk membalik setiap kata dalam kalimat
func reverseString(s string) string {
	str := []rune(s)
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		str[i], str[j] = str[j], str[i]
	}
	return string(str)
}

// fungsi untuk membalik setiap kata dalam kalimat
func reverseWords(sentence string) string {
	words := strings.Fields(sentence)
	for i := range words {
		words[i] = reverseString(words[i])
	}
	return strings.Join(words, " ")
}

func main() {
	fmt.Println(reverseWords("italem irad irigayaj"))
	fmt.Println(reverseWords("iadab itsap ulalreb"))
	fmt.Println(reverseWords("nalub kusutret gnalali"))
	fmt.Println(reverseWords("kasur rusak"))

}
