package album

import "github.com/gin-gonic/gin"

func AlbumRoutes(router *gin.Engine) {

	albums := router.Group("/album")

	albums.DELETE("/:album_id", DeleteAlbum())
	albums.PATCH("/:album_id", UpdateAlbum())

}
