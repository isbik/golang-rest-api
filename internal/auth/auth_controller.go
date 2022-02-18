package auth

import (
	"main/internal/user"
	"main/pkg/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUser user.User

		if err := c.BindJSON(&loginUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		result, err := user.CountUsersByEmail(loginUser.Email)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if result == 0 {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Email already taken"})
			return
		}

		user.FindUserByEmail(loginUser.Email, &loginUser)

		token := services.GenerateToken(services.JWTUser{Id: loginUser.ID.Hex()})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		response := make(map[string]interface{})

		response["token"] = token
		response["user"] = loginUser

		c.JSON(http.StatusOK, response)
	}
}

func Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		var newUser user.User

		if err := c.BindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		err := validator.New().Struct(newUser)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		result, err := user.CountUsersByEmail(newUser.Email)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if result > 0 {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Email already taken"})
			return
		}

		user.CreateUser(&newUser)

		token := services.GenerateToken(services.JWTUser{Id: newUser.ID.Hex()})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		response := make(map[string]interface{})

		response["token"] = token
		response["user"] = newUser

		c.JSON(http.StatusOK, response)
	}
}
