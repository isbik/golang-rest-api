package services

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTUser struct {
	Id string `json:"id"`
}

type CustomClaims struct {
	Id string `json:"id"`
	jwt.StandardClaims
}

var secret = []byte("my_secret_key")

func GenerateToken(user JWTUser) string {
	fmt.Println("user", user)

	fmt.Println(user.Id)

	claims := &CustomClaims{
		Id: user.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString(secret)
	if err != nil {
		panic(err)
	}

	return t
}

func ParseToken(tokenString string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})

}
