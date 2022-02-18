package user

import (
	"main/internal/photo"
	"main/pkg/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HashPassword()   {}
func VerifyPassword() {}
func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("user_id")

		var user User
		err := FindUserById(userId, &user)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, err)
			return
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
		userId := c.GetString("userId")
		var user User

		err := FindUserById(userId, &user)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, user)
	}
}

func GetUserPhotos() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.GetString("userId")

		page, limit := helpers.Pagination(c)

		var photos []photo.Photo

		err := photo.FindUserPhotos(userId, page, limit, &photos)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusOK, err)
			return
		}

		c.JSON(http.StatusOK, photos)
	}
}
