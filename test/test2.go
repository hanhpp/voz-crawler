package main

import (
	"fmt"
	"time"
	"voz/cronjob"

	"voz/entity"
	"voz/global"
)

func main() {
	global.FetchEnvironmentVariables()
	entity.InitializeDatabaseConnection()
	entity.ProcessMigration()
	CrawlThreads()
	//Hold it here
	for {
		select {}
	}
}

func CrawlThreads() {
	ticker := time.NewTicker(5000 * time.Millisecond)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
				cronjob.CrawlThreadFromUrl(global.F17_Pages)
			}
		}
	}()
}

func CrawlComments() {
	ticker := time.NewTicker(3000 * time.Millisecond)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
				//cronjob.CrawlComments(global.F17_Pages)
			}
		}
	}()
}