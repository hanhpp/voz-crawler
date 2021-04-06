package routes

import (
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"voz/config"
	"voz/entity"
	"voz/model"
	"voz/utils"
)

func InitCommentRoutes(route *gin.Engine) {
	route.GET("/comments", Comments)
	route.GET("/comments/:id", CommentById)
	route.GET("/thread-comments/:id", CommentsByThreadId)
}

func Comments(c *gin.Context) {
	logger := config.GetLogger()
	color.Red("Received comment request")
	comments := []model.Comment{}
	db := entity.GetDBInstance().Debug()
	db = db.Order("updated_at DESC").Limit(50)
	err := db.Find(&comments).Error
	if err != nil {
		logger.Errorln(err)
		utils.BadRequestWithMessage(c,"can't get comments from database")
		return
	}
	//color.Blue("%v",comments)
	color.Cyan("Found %d comments",len(comments))
	c.JSON(200, comments)
	return
}

func CommentById(c *gin.Context) {
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

func CommentsByThreadId(c *gin.Context) {
	logger := config.GetLogger()
	color.Red("Received threads request")
	comments := []model.Comment{}
	threadId := c.Param("id")
	db := entity.GetDBInstance().Debug()
	db = db.Where("thread_id = ?",threadId).Order("updated_at DESC").Limit(50)
	err := db.Find(&comments).Error
	if err != nil {
		logger.Errorln(err)
		utils.BadRequestWithMessage(c,"can't get threads from database")
		return
	}
	//color.Blue("%v",threads)
	color.Cyan("Found comments \n%+v",comments)
	c.JSON(200, comments)
	return
}