package config

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
)

var loggerInstance *logrus.Logger

//initLogger create our local config instance
func initLogger() {
	logger := logrus.New()

	//Check if logs folder exists, create if not
	logsPath := filepath.Join(".", "logs")
	err := os.MkdirAll(logsPath, 0777)
	if err != nil {
		logger.Error(err)
	}

	//Currently only log out to one file
	filePath := filepath.Join("logs", GetFileName())
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logger.Out = file
	} else {
		logger.Infof("Failed to log to file with err : %v, using default stderr\n", err)
	}

	if loggerInstance != nil {
		color.Red("Logger initiated to file %s!", filePath)
	}
	//Show line number and function name
	logger.SetReportCaller(true)
	loggerInstance = logger
}

//GetLogger get our config instance
func GetLogger() *logrus.Logger {
	if loggerInstance == nil {
		initLogger()
	}

	return loggerInstance
}

//GetFileName ex : 01-02-2006.log
func GetFileName() string {
	return fmt.Sprintf("%s.log", time.Now().Format("01-02-2006"))
}

type SkipLogger struct{}

func (l SkipLogger) Info(msg string, keysAndValues ...interface{}) {
	//Currently only printing out "skip", uncomment if you want to monitor skip state
	//fmt.Println(msg)
}

func (l SkipLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	GetLogger().Errorln(err, msg)
}
