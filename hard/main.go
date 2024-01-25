package main

import (
	"fmt"
	"reflect"
	"strings"
)

func areAnagrams(str1, str2 string) bool {
	str1Lower := strings.ToLower(str1)
	str2Lower := strings.ToLower(str2)

	return reflect.DeepEqual(countChars(str1Lower), countChars(str2Lower))
}

func countChars(s string) map[rune]int {
	counts := make(map[rune]int)

	for _, char := range s {
		counts[char]++
	}

	return counts
}

func main() {
	word1 := "Race"
	word2 := "Care"

	result := areAnagrams(word1, word2)

	fmt.Println(result)
}
