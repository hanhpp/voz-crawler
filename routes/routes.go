package routes

import (
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"voz/config"
	"voz/entity"
	"voz/model"
	"voz/utils"
)

func InitRoutes() {
	logger := config.GetLogger()
	route := gin.Default()
	route.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	route.GET("/threads", func(c *gin.Context) {
		color.Red("Received threads request")
		threads := []model.Thread{}
		db := entity.GetDBInstance().Debug()
		db = db.Order("updated_at DESC").Limit(20)
		err := db.Find(&threads).Error
		if err != nil {
			logger.Errorln(err)
			utils.BadRequestWithMessage(c,"can't get threads from database")
			return
		}
		//color.Blue("%v",threads)
		color.Cyan("Found %d threads",len(threads))
		c.JSON(200, threads)
	})
	route.Run("localhost:9500") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
