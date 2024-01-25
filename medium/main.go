package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isPalingdrome(inputStr string) bool {
	var cleanedStr strings.Builder

	for _, char := range inputStr {
		if unicode.IsLetter(char) || unicode.IsNumber(char) {
			cleanedStr.WriteRune(unicode.ToLower(char))
		}
	}

	cleanedInput := cleanedStr.String()

	for i := 0; i < len(cleanedInput)/2; i++ {
		if cleanedInput[i] != cleanedInput[len(cleanedInput)-1-i] {
			return false
		}
	}

	return true
}

func main() {
	inputStr := "A man, a plan, a canal: Panama"

	result := isPalingdrome(inputStr)

	fmt.Println(result)
}
