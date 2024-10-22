package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(id int64, username, secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":       id,
			"username": username,
			"exp":      time.Now().Add(10 * time.Minute).Unix(),
		},
	)
	key := []byte(secretKey)
	tokenStr, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func ValidateToken(tokenStr, secretKey string) (int64, string, error) {
	key := []byte(secretKey)
	claims:=jwt.MapClaims{}

	token,err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return 0, "", err
	}
	if !token.Valid {
		return 0, "", errors.New("invalid token")
	}
	return int64(claims["id"].(float64)),claims["username"].(string), nil
}
