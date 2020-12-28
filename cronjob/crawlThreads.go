package cronjob

import (
	"errors"
	"github.com/fatih/color"
	"github.com/gocolly/colly"
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
	"voz/config"
	"voz/entity"
	"voz/global"
	"voz/model"
)

func CrawlThreads(url string, fileName string) {
	color.Green("cron job: Crawling thread from %s\nSaving into file /text/%s.txt", url, fileName)
	skipLogger := config.SkipLogger{}
	c := cron.New(
		cron.WithLocation(time.UTC),
		cron.WithChain(cron.SkipIfStillRunning(skipLogger)),
	)
	_, _ = c.AddFunc("@every 0h0m10s", func() { // 23h59m GMT +8
		config.GetLogger().Info("Running crawler...")
		VisitAndCollectThreadsFromURL(url, fileName)
	})
	c.Start()
}

func VisitAndCollectThreadsFromURL(URL string, fileName string) {
	c := colly.NewCollector()

	basePath := "./text"
	path := filepath.Join(basePath, fileName)
	f := createFile(path)
	defer f.Close()

	var titles []string
	c.OnHTML(global.ThreadTitle, func(e *colly.HTMLElement) {
		_, err := f.Write([]byte(e.Text))
		if err != nil {
			log.Fatal(err)
		}
		err = handleThreadContent(e, titles)
		logger := config.GetLogger()
		if err != nil {
			logger.Errorln(err)
		}
	})
	_ = c.Visit(URL)
}

func handleThreadContent(e *colly.HTMLElement, titles []string) error {
	logger := config.GetLogger()
	text := standardizeSpaces(e.Text)
	titles = append(titles, text)
	thread := &entity.Thread{}
	err := e.Unmarshal(thread)
	if err != nil {
		log.Fatal(err)
	}
	threadId := GetThreadID(thread.Link)
	link := global.VozBaseURL + thread.Link
	newThread := &model.Thread{
		Title:    thread.Title,
		Link:     link,
		ThreadId: threadId,
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
	} else {
		color.Red("Thread %d already exists!", threadId)
		color.Red("Link : %s", thread.Link)
	}
	return nil
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

//standardizeSpaces remove redundancies spaces
func standardizeSpaces(s string) string {
	//fields return splitted array of chars if function satisfy
	return strings.Join(strings.Fields(s), " ")
}

func createFile(name string) *os.File {
	f, err := os.Create(name + ".txt")

	if err != nil {
		log.Fatal(err)
	}

	return f
}
