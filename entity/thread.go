package entity

import (
	"voz/config"
	"voz/model"
)

//type Thread struct {
//	Title    string `selector:"a"`
//	Link     string `selector:"a[href]" attr:"href"`
//	PageJump string `selector:".structItem-pageJump"`
//}

type Thread struct {
	Title    string `selector:".structItem-title > a"`
	Link     string `selector:".structItem-title > a[href]" attr:"href"`
	PageJump []string `selector:".structItem-pageJump > a"`
}

func Save(thr model.Thread) error {
	logger := config.GetLogger()
	err := GetDBInstance().Save(thr).Error
	if err != nil {
		logger.Errorln(err)
		return err
	}
	return nil
}