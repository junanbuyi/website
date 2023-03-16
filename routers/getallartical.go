package routers

import (
	"test/db"

	"github.com/gin-gonic/gin"
)

func Getallartical(c *gin.Context) {
	articles, err := db.GetAllArticles()
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(200, articles)
	}
}
