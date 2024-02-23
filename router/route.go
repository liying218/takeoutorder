package router

import "github.com/gin-gonic/gin"

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"isok": true,
		})
	})
	return r
}
