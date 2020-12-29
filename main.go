package main

import (
	"github.com/fatih/color"
	"sync"
	"voz/cronjob"

	"voz/entity"
	"voz/global"
)

func main() {
	global.FetchEnvironmentVariables()
	entity.InitializeDatabaseConnection()
	entity.ProcessMigration()

	color.Blue("Starting voz crawler version 1.0")
	var wg sync.WaitGroup
	wg.Add(1)
	cronjob.RunCronjob()
	wg.Wait()
}
