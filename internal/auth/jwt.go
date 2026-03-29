package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateToken creates JWT for user
func GenerateToken(userID int, secret []byte) (string, error) {

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secret) // returns the signed token string along with any errors that may occur during the signing process
}

// ValidateToken validates the token and returns the user id if valid
func ValidateToken(tokenString string, secret []byte) (int, error) {

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil || !token.Valid {
		return 0, err
	}

	claims := token.Claims.(jwt.MapClaims)

	userID := int(claims["user_id"].(float64))

	return userID, nil
}
