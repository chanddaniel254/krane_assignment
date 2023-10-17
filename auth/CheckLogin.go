package auth

import (
	"errors"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt"
)

func CheckLogin(tokenString string) (string, error) {
	secret := os.Getenv("SECRET")

	if tokenString == "" {
		return "", errors.New("Token not found")
	}
	token, err := verifyToken(tokenString, secret)
	if err != nil || !token.Valid {
		return "", errors.New("token verification failed")
	}

	// Access claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("failed to access claim")
	}

	// Example: Access a specific claim and convert it to an integer
	userId, ok := claims["id"].(string)
	if !ok {
		return "", errors.New("failed to access the 'id' claim as an integer")
	}

	return userId, nil

}

func verifyToken(tokenString, secretKey string) (*jwt.Token, error) {
	// Parse the token with the provided secret key
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if token.Method.Alg() != "HS256" {
			return nil, fmt.Errorf("Invalid signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
