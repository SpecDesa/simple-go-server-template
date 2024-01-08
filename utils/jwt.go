package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "averysecretkey"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (int64, error) {
	parsedToken, error := jwt.Parse(token, func(token *jwt.Token) (any, error) {

		// Check that it is signed with your signing method above.
		_, ok := token.Method.(*jwt.SigningMethodHMAC) // Would return the value (dont care) and bool

		if !ok {
			return nil, errors.New("Unexpected signing method")
		}

		return []byte(secretKey), nil // Need to byte slice it because when signing needed to do that.
	})

	if error != nil {
		return 0, errors.New("Could not parse token.")
	}

	tokenIsValid := parsedToken.Valid

	if !tokenIsValid {
		return 0, errors.New("Invalid token.")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims) // Check if it is our claims type set in generate token

	if !ok {
		return 0, errors.New("Invalid token claims.")
	}

	// email := claims["email"].(string)
	userId := int64(claims["userId"].(float64))

	return userId, nil
}
