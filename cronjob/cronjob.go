package cronjob

import (
	"github.com/fatih/color"
	"voz/global"
	"voz/model"
)

var Threads = make(chan *model.Thread,100)

func RunCronjob() {
	go CrawlThreadsFromF17()
	//go CrawlThreads(global.F33, "diem-bao")
	go CommentLoop()

}

func CrawlThreadsFromF17() {
	//for _,v := range global.F17_Pages {
	//	color.Red("%s",v)
	//}
	go CrawlThreads(global.F17, "chuyen-tro-linh-tinh")
	go CrawlThreads(global.F17_P2, "chuyen-tro-linh-tinh")
	go CrawlThreads(global.F17_P3, "chuyen-tro-linh-tinh")
	go CrawlThreads(global.F17_P4, "chuyen-tro-linh-tinh")
	go CrawlThreads(global.F17_P5, "chuyen-tro-linh-tinh")
	go CrawlThreads(global.F17_P6, "chuyen-tro-linh-tinh")
	go CrawlThreads(global.F17_P7, "chuyen-tro-linh-tinh")
}

func CommentLoop() {
	for {
		select {
		case thread := <-Threads:
			color.Red("Received %s from Thread queue", thread.Link)
			go CrawlComments(thread.Link, thread.ThreadId)
		}
	}
}
