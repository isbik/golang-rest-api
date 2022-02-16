package auth

import (
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {

	router.POST("/auth/register", Register())
	// router.POST("/auth/login", Login())
}
