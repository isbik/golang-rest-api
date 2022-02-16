package middleware

import (
	"fmt"
	"net/http"

	"main/pkg/services"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "

		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA):]

		token, err := services.ValidateToken(tokenString)
		if token.Valid {
			if claims, ok := token.Claims.(*services.CustomClaims); ok {

				fmt.Println(claims.Id)
			} else {
				c.AbortWithStatusJSON(http.StatusInternalServerError, "Error paring token")
			}
		} else {
			fmt.Println(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, err)
		}
	}
}

// ValidateToken
