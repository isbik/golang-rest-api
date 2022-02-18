package photo

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoadPhotos() gin.HandlerFunc {

	return func(c *gin.Context) {
		userId := c.GetString("userId")

		photos, err := LoadPhotosFromJsonPlaceholder()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, err)
			return
		}

		fmt.Println("call")

		InsertPhotos(userId, photos)

		c.JSON(http.StatusOK, "Ok")
	}
}

func DeletePhoto() gin.HandlerFunc {

	return func(c *gin.Context) {
		photoId := c.Param("photo_id")

		err := DeletePhotoById(photoId)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, err)
			return
		}

		c.JSON(http.StatusOK, "Ok")
	}
}

func GetPhoto() gin.HandlerFunc {

	return func(c *gin.Context) {
		photoId := c.Param("photo_id")

		var photo Photo

		err := FindPhotoById(photoId, &photo)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, err)
			return
		}

		c.JSON(http.StatusOK, photo)
	}
}
