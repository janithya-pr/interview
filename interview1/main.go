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

	// JWT Token Generation
	token, err := GenerateToken("TU-001", "Test User", "test@mail.com")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("JWT Token:", token)

	// JWT Token Validation
	claims, err := ValidateToken(token)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Name:", claims.Name)
	fmt.Println("Email:", claims.Email)
	fmt.Println("Public ID:", claims.Subject)
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