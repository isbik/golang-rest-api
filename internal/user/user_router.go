package user

import (
	"main/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {

	users := router.Group("/users", middleware.AuthMiddleware())

	users.GET("/", GetUsers())
	users.GET("/me", GetCurrentUser())
	users.GET("/:user_id", GetUser())
	users.GET("/:user_id/photos", GetUserPhotos())
}
