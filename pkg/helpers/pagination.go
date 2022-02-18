package helpers

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func Pagination(c *gin.Context) (int64, int64) {
	if c.Query("page") != "" && c.Query("limit") != "" {
		page, _ := strconv.ParseInt(c.Query("page"), 10, 32)
		limit, _ := strconv.ParseInt(c.Query("limit"), 10, 32)
		if page == 1 {
			return page, limit
		}
		return page, limit
	}
	return 0, 10
}
