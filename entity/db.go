package entity

import (
	"fmt"
	"github.com/fatih/color"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
	"voz/config"
	"voz/global"
)

var dBInstance *gorm.DB

//GetLogger get our config instance
func GetDBInstance() *gorm.DB {
	if dBInstance == nil {
		InitializeDatabaseConnection()
	}

	return dBInstance
}

func InitializeDatabaseConnection() {
	databaseName := fmt.Sprintf("Postgres")
	// Connect to Database
	var err error

	logLevel := logger.Error
	//logLevel := logger.Silent
	dBInstance, err = gorm.Open(postgres.Open(global.Config.PostgresConnectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	if err == nil {
		db, _ := dBInstance.DB()
		db.SetMaxIdleConns(5)
		db.SetMaxOpenConns(10)
		db.SetConnMaxLifetime(5 * time.Minute)
		color.Green("Yay! " + databaseName + " Database Connected!")
	} else {
		fmt.Println(databaseName, "Connection ERROR!", err)
	}
}

func ProcessMigration() {
	DBAutoMigration()
}

func DBAutoMigration() {
	logger := config.GetLogger()
	err := GetDBInstance().Debug().AutoMigrate(
		&Thread{},
		&Comment{},
	)
	if err != nil {
		logger.Errorln(err)
		return
	}
}
