package util

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetPaging(c *gin.Context, offset, limit *int32) {
	limitQuery := c.Query("limit")
	if len(limitQuery) == 0 {
		limitQuery = "10"
	}

	_limit, err := strconv.Atoi(limitQuery)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid limit query parameter"})
		return
	}

	pageQuery := c.Query("page")
	if len(pageQuery) == 0 {
		pageQuery = "1"
	}

	p, err := strconv.Atoi(pageQuery)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid page query parameter"})
		return
	}

	*limit = int32(_limit)
	*offset = CalculateOffset(p, _limit)
}
