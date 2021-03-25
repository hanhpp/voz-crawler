package main

import (
	"voz/entity"
	"voz/global"
	"voz/routes"
)

func main() {
	global.FetchEnvironmentVariables()
	entity.InitializeDatabaseConnection()
	entity.ProcessMigration()
	routes.InitRoutes()
}