package cronjob

import (
	"fmt"
	"github.com/fatih/color"
	"sync"
	"time"
	"voz/config"
	"voz/entity"
	"voz/global"
	"voz/model"
)

var Threads = make(chan *model.Thread,100)

func RunCronjob() {
	go ThreadTicker()
	go CommentLoop()
	//go FindDeletedThreads()
}
//go CrawlThreads(global.F33, "diem-bao")

func FindDeletedThreads() {
	logger := config.GetLogger()
	for {
			cmts := []model.Comment{}
			count := int64(0)
			err := entity.GetDBInstance().Debug().Where("text like '%lightbox_error%'").Find(&cmts).Count(&count).Error
			if err != nil {
				logger.Errorln(err)
				return
			}
			color.Red("Found %d deleted threads",count)
			thrIds := []uint64{}
			for _,cmt := range cmts {
				thrIds = append(thrIds,cmt.ThreadId)
			}
			threads := []model.Thread{}
			err = entity.GetDBInstance().Debug().Where("thread_id in (?)",thrIds).Find(&threads).Error
			if err != nil {
				logger.Errorln(err)
				return
			}
			for _,thr := range threads {
				delThr := createDeletedThread(thr)
				err = entity.GetDBInstance().Debug().Save(&delThr).Error
				if err != nil {
					logger.Errorln(err)
					return
				}
			}
			time.Sleep(30*time.Second)
	}
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
				CrawlThreadFromUrl(global.F33_Pages)
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
				CrawlComments(thread)
			}()
		}
	}
}

func UpdateLocalThread() {
	logger := config.GetLogger()
	localThreads := []model.Thread{}
	err := entity.GetDBInstance().Model(&model.Thread{}).Find(&localThreads).Error
	if err != nil {
		logger.Errorln(err)
		return
	}
	l := len(localThreads)
	if l < 2 {
		return
	}
	color.Red("Found %d threads",l)
	color.Cyan("First thread : \n%+v",localThreads[0])
	color.Cyan("Last thread : \n%+v",localThreads[l-1])
	for _,v := range localThreads {
		//Push into it again to update
		go func() {
			Threads <- &v
		}()
	}
}
