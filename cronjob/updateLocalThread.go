package cronjob

import (
	"github.com/fatih/color"
	"voz/config"
	"voz/entity"
	"voz/model"
)

func UpdateLocalThread() {
	logger := config.GetLogger()
	localThreads := []model.Thread{}
	err := entity.GetDBInstance().Model(&model.Thread{}).Find(&localThreads).Error
	if err != nil {
		logger.Errorln(err)
		return
	}
	l := len(localThreads)
	if l < 2 {
		return
	}
	color.Red("Found %d threads",l)
	color.Cyan("First thread : \n%+v",localThreads[0])
	color.Cyan("Last thread : \n%+v",localThreads[l-1])
	for _,v := range localThreads {
		//Push into it again to update
		go func() {
			Threads <- &v
		}()
	}
}
