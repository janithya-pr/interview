package main

import (
	"fmt"
	"strings"
)

func main() {
	// Count Vowels
	total := CountVowels("Hello, world!")
	fmt.Println("Total vowels:", total)

	// Password Hashing
	PasswordHash("abc123", "xyz")
}

// Number of vowels in a given text
func CountVowels(text string) int {
	count := 0

	for _, char := range strings.ToLower(text) {
		switch char {
		case 'a', 'e', 'i', 'o', 'u':
			count++
		}
	}
	return count
}