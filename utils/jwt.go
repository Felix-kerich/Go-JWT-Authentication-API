package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWT secret key
var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

// Claims struct for JWT
type Claims struct {
	Email string `json:"email"`
	Type  string `json:"type"`
	jwt.RegisteredClaims
}

// GenerateJWT creates a new JWT token
func GenerateJWT(email, tokenType string, duration time.Duration) string {
	expirationTime := time.Now().Add(duration)
	claims := &Claims{
		Email: email,
		Type:  tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		log.Println("Error generating token:", err)
		return ""
	}
	return tokenString
}

// ValidateJWT checks if a JWT is valid and returns the email
func ValidateJWT(tokenString, expectedType string) (string, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		return "", fmt.Errorf("invalid token")
	}
	if claims.Type != expectedType {
		return "", fmt.Errorf("token type mismatch")
	}
	return claims.Email, nil
}
