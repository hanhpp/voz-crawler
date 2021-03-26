package routes

import (
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"voz/config"
	"voz/entity"
	"voz/model"
	"voz/utils"
)

func InitThreadRoutes(route *gin.Engine) {
	route.GET("/thread", Threads)
	route.GET("/thread/:id",ThreadByThreadId)
}

func Threads(c *gin.Context) {
	logger := config.GetLogger()
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
	return
}

func ThreadByThreadId(c *gin.Context) {
	logger := config.GetLogger()
	color.Red("Received threads request")
	thread := &model.Thread{}
	threadId := c.Param("id")
	db := entity.GetDBInstance().Debug()
	db = db.Where("thread_id = ?",threadId).Order("updated_at DESC").Limit(20)
	err := db.Find(&thread).Error
	if err != nil {
		logger.Errorln(err)
		utils.BadRequestWithMessage(c,"can't get threads from database")
		return
	}
	//color.Blue("%v",threads)
	color.Cyan("Found thread %+v",thread)
	c.JSON(200, thread)
	return
}