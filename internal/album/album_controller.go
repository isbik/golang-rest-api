package album

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteAlbum() gin.HandlerFunc {

	return func(c *gin.Context) {
		albumId := c.Param("album_id")

		err := DeleteAlbumById(albumId)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, err)
			return
		}

		c.JSON(http.StatusOK, "Ok")
	}
}

func UpdateAlbum() gin.HandlerFunc {

	return func(c *gin.Context) {
		albumId := c.Param("album_id")

		var album Album

		err := c.ShouldBindJSON(&album)
		if err == nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, err)
			return
		}

		err = UpdateAlbumById(albumId, &album)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, album)
	}
}
