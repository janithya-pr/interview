package main

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWT Claims structure
type JWTClaims struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

var SecretKey = []byte("super-secret-key")

// Generate JWT tokens using HS256
func GenerateToken(publicID, name, email string) (string, error) {
	now := time.Now()

	claims := JWTClaims{
		Name:  name,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   publicID,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(15 * time.Second)),
		},
	}
	
	// token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, claims)
	// return token.SignedString(PrivateKey)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(SecretKey)
}

// Validate JWT tokens
func ValidateToken(tokenString string) (*JWTClaims, error) {
	claims := &JWTClaims{}

	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(token *jwt.Token) (interface{}, error) {

			// Verify signing method
			if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
				return nil, jwt.ErrSignatureInvalid
			}
			return SecretKey, nil
		},
	)

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}

// func GenerateAccessToken(userID string, privateKey ed25519.PrivateKey) (string, error) {}
// func GenerateRefreshToken(userID string, privateKey ed25519.PrivateKey) (string, error) {}
// func ValidateToken(tokenString string, privateKey ed25519.PublicKey) (*JWTClaims, error) {}