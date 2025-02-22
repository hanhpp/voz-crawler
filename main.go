package main

import (
	"github.com/fatih/color"
	"sync"
	"voz/cronjob"

	"voz/entity"
	"voz/global"
)

const Version string = "1.1"
func main() {
	global.FetchEnvironmentVariables()
	entity.InitializeDatabaseConnection()
	entity.ProcessMigration()

	color.Blue("Starting voz crawler version %s",Version)
	var wg sync.WaitGroup
	wg.Add(1)
	cronjob.RunCronjob()
	wg.Wait()
}
