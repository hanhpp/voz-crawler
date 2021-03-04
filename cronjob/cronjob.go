package cronjob

import (
	"fmt"
	"github.com/fatih/color"
	"sync"
	"time"
	"voz/global"
	"voz/model"
)

var Threads = make(chan *model.Thread,100)

func RunCronjob() {
	go ThreadTicker()
	go CommentLoop()
	go UpdateLocalThread()
}
//go CrawlThreads(global.F33, "diem-bao")
func CrawlThreadsFromF17() {
	go CrawlThreads(global.F17)
	go CrawlThreads(global.F17_P2)
	go CrawlThreads(global.F17_P3)
	go CrawlThreads(global.F17_P4)
	go CrawlThreads(global.F17_P5)
	//go CrawlThreads(global.F17_P6, "chuyen-tro-linh-tinh")
	//go CrawlThreads(global.F17_P7, "chuyen-tro-linh-tinh")
	//go CrawlThreads(global.F17_P8, "chuyen-tro-linh-tinh")
	//go CrawlThreads(global.F17_P9, "chuyen-tro-linh-tinh")
	//go CrawlThreads(global.F17_P10, "chuyen-tro-linh-tinh")
}

func CrawlThreadFromUrl(urls []string) {
	var wg sync.WaitGroup
	for _,url := range urls {
		go CrawlThreads(url)
		wg.Add(1)
	}
	//Make sure we finish everything before returning
	wg.Wait()
	wg.Done()
}

func ThreadTicker() {
	//Run once every 5 seconds
	ticker := time.NewTicker(5000 * time.Millisecond)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
				CrawlThreadFromUrl(global.F17_Pages)
			}
		}
	}()
}

func CommentLoop() {
	for {
		select {
		case thread := <-Threads:
			go func() {
				color.Red("Received %s from Thread queue", thread.Link)
				CrawlComments(thread.Link, thread.ThreadId)
			}()
		}
	}
}
