package routes

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes() {
	route := gin.Default()
	route.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	InitThreadRoutes(route)
	InitCommentRoutes(route)
	_ = route.Run("localhost:9500") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
