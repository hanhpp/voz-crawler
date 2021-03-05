package cronjob

import (
	"errors"
	"github.com/fatih/color"
	"github.com/gocolly/colly"
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
	"log"
	"os"
	"regexp"
	"strconv"
	"voz/config"
	"voz/entity"
	"voz/global"
	"voz/model"
	"voz/utils"
)

func CrawlThreads(url string) {
	color.Green("Crawling thread from [%s]", url)
	skipLogger := config.SkipLogger{}
	c := cron.New(
		cron.WithChain(cron.SkipIfStillRunning(skipLogger)),
	)
	_, _ = c.AddFunc(utils.GetCronString(global.CrawlInterval), func() {
		VisitAndCollectThreadsFromURL(url)
	})
	c.Start()
}

func VisitAndCollectThreadsFromURL(URL string) {
	c := colly.NewCollector()
	var titles []string
	c.OnHTML(global.ThreadStruct, func(e *colly.HTMLElement) {
		err := handleThreadContent(e, titles, URL)
		logger := config.GetLogger()
		if err != nil {
			logger.Errorln(err)
		}
	})
	_ = c.Visit(URL)
}

func handleThreadContent(e *colly.HTMLElement, titles []string, parentURL string) error {
	logger := config.GetLogger()
	text := utils.StandardizeSpaces(e.Text)
	titles = append(titles, text)
	//color.Blue("%+v", e)
	thread := &entity.Thread{}
	err := e.Unmarshal(thread)
	if err != nil {
		log.Fatal(err)
	}
	threadId := GetThreadID(thread.Link)
	lastPage := GetLastPage(thread)
	link := global.VozBaseURL + thread.Link
	newThread := &model.Thread{
		Title:     thread.Title,
		Link:      link,
		ThreadId:  threadId,
		ParentURL: parentURL,
		LastPage:  lastPage,
	}
	//color.Red("%v", newThread)
	localThread := &model.Thread{}
	err = entity.GetDBInstance().Where("thread_id = ?", threadId).First(&localThread).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		//Only create when there is no such record with same threadID
		err = entity.GetDBInstance().Create(&newThread).Error
		if err != nil {
			logger.Errorln(err)
			return err
		}
		color.Cyan("Thread %d saved success %s\nTitle : %s", threadId, thread.Link, thread.Title)

		//Push it to our link queue
		color.Red("Pushing to Thread queue\n%+v", newThread)
		Threads <- newThread
	} else {
		//logger.WithField(" threadId", threadId).WithField("thread.Link", thread.Link).Info("Thread already exists!")
		//If thread already exists, check if last page is higher than database last page
		color.Green("[%d]local last page vs remote : %d ? %d\n%s",newThread.ThreadId,localThread.LastPage,newThread.LastPage,newThread.Link)
		if newThread.LastPage > localThread.LastPage {
			color.Red("Local thread's comments outdated, proceed to update comments... (%d < %d)",localThread.LastPage,newThread.LastPage)
			//LastPage is higher, proceed to push thread to queue
			Threads <- newThread
		}
	}
	return nil
}

func GetLastPage(thread *entity.Thread) uint64 {
	logger := config.GetLogger()
	if thread == nil {
		return 1
	}
	pages := thread.PageJump
	l := len(pages)
	if l < 2 {
		return 1
	}
	val, err := strconv.ParseUint(pages[l-1], 10, 64)
	if err != nil {
		logger.Errorln(err)
		return 1
	}
	return val
}
func GetThreadID(link string) uint64 {
	logger := config.GetLogger()
	r := regexp.MustCompile(global.ThreadIDRegex)
	result := r.FindAllString(link, -1)
	if len(result) > 0 {
		str := result[0][1:]
		l := len(str)
		//regex result format .123456/ : so we remove . and /
		res, err := strconv.ParseUint(str[:l-1], 10, 64)
		if err != nil {
			logger.Errorln(err)
			return 0
		}
		return res
	}
	return 0
}


func createFile(name string) *os.File {
	f, err := os.Create(name + ".txt")

	if err != nil {
		log.Fatal(err)
	}

	return f
}
