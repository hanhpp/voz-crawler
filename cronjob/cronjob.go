package cronjob

import (
	"github.com/fatih/color"
	"voz/global"
	"voz/model"
)

var Threads = make(chan *model.Thread)

func RunCronjob() {
	go CrawlThreads(global.F17, "thread")
	for {
		select {
		case thread := <-Threads:
			color.Red("Received %s from Thread queue", thread.Link)
			go CrawlComments(thread.Link, thread.Title, thread.ThreadId)
		}
	}
}
