package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HashPassword()   {}
func VerifyPassword() {}
func GetUser() gin.HandlerFunc {

	return func(c *gin.Context) {
		userId := c.Param("user_id")

		user, err := FindUserById(userId)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, err)
		}

		c.JSON(http.StatusOK, user)
	}
}

func GetUsers() gin.HandlerFunc {

	return func(c *gin.Context) {
		users, err := FindAllUsers()

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, users)
	}
}

func GetCurrentUser() gin.HandlerFunc {

	return func(c *gin.Context) {
		users, err := FindUserById("a3")

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, users)
	}
}
