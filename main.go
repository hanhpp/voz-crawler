package main

import (
	"fmt"
	"sync"
	"voz/cronjob"

	"voz/entity"
	"voz/global"
)

func main() {
	global.FetchEnvironmentVariables()
	entity.InitializeDatabaseConnection()
	entity.ProcessMigration()

	var wg sync.WaitGroup
	wg.Add(1)
	cronjob.RunCronjob()
	wg.Wait()
	//VisitAndCollectThreadsFromURL(global.F17, "title")
	fmt.Println("done")
}
