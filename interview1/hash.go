package main

import (
	"fmt"
	
	"golang.org/x/crypto/bcrypt"
)

// Password hashing and validation using bcrypt
func PasswordHash(password, loginPassword string) {
	// HASH PASSWORD
	hashedPassword, _ := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	fmt.Printf("HASH PASSWORD: %v\n", string(hashedPassword))

	// VALIDATE PASSWORD
	err := bcrypt.CompareHashAndPassword(
		hashedPassword,
		[]byte(loginPassword),
	)

	if err == nil {
		fmt.Println("loging..")
	} else {
		fmt.Println(err)
	}
}