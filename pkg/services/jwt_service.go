package services

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

type JWTUser struct {
	Id string `json:"id"`
}

type CustomClaims struct {
	jwt.StandardClaims
}

var secret = []byte(viper.GetString("jwt.secret"))

func GenerateToken(user JWTUser) (string, error) {
	claims := &CustomClaims{
		StandardClaims: jwt.StandardClaims{
			Id:        user.Id,
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return t, nil
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return secret, nil
	})

}
