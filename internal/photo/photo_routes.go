package photo

import (
	"main/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func PhotoRoutes(router *gin.Engine) {

	albums := router.Group("/photos")
	albums.GET("/:photo_id", GetPhoto())

	protected := albums.Group("/", middleware.AuthMiddleware())
	protected.POST("/load", LoadPhotos())
	protected.DELETE("/:photo_id", DeletePhoto())

}
