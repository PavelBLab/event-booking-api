package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "superSecretKey@345"

func GenerateJWT(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func VerifyJWT(tokenString string) (int64, error) {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, jwt.ErrInvalidKey
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, err
	}

	tokenIsValid := parsedToken.Valid

	if !tokenIsValid {
		return 0, jwt.ErrSignatureInvalid
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok || claims == nil {
		return 0, jwt.ErrInvalidKey
	}

	//email := claims["email"].(string)
	userId := int64(claims["userId"].(float64))

	return userId, nil
}
