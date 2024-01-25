package main

import "fmt"

func reverseCharacters(inputStr string) string {
	runes := []rune(inputStr)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func main() {
	inputStr := "Golang is fun :v"

	result := reverseCharacters(inputStr)

	fmt.Println(result)
}
