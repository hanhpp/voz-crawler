package cronjob

import (
	"github.com/fatih/color"
	"github.com/gocolly/colly"
	"log"
	"path/filepath"
	"strings"
	"voz/config"
	"voz/entity"
	"voz/global"
)

func CrawlComments(url string, fileName string) {
	color.Green("cron job: Crawling thread from %s\nSaving into file /text/%s.txt", url, fileName)
	//skipLogger := config.SkipLogger{}
	//c := cron.New(
	//	cron.WithLocation(time.UTC),
	//	cron.WithChain(cron.SkipIfStillRunning(skipLogger)),
	//)
	//_, _ = c.AddFunc("@every 0h0m10s", func() { // 23h59m GMT +8
	//	config.GetLogger().Info("Running crawler...")
	VisitAndCollectCmtsFromURL(url, fileName)
	//})
	//c.Start()
}

func VisitAndCollectCmtsFromURL(URL string, fileName string) {
	c := colly.NewCollector()

	basePath := "./text"
	path := filepath.Join(basePath, fileName)
	f := createFile(path)
	defer f.Close()

	var titles []string
	c.OnHTML(global.CommentStruct, func(e *colly.HTMLElement) {
		//color.Red("%s", e.Text)
		_, err := f.Write([]byte(e.Text))
		if err != nil {
			log.Fatal(err)
		}
		//color.Cyan("%+v", e)
		err = handleCmtsContent(e, titles)
		logger := config.GetLogger()
		if err != nil {
			logger.Errorln(err)
		}
	})
	_ = c.Visit(URL)
}

func handleCmtsContent(e *colly.HTMLElement, titles []string) error {
	//logger := config.GetLogger()
	text := standardizeSpaces(e.Text)
	titles = append(titles, text)
	cmt := &entity.Comment{}
	err := e.Unmarshal(cmt)
	if err != nil {
		log.Fatal(err)
	}
	cmt.Desc = e.Attr(global.CommentNamespace)
	color.Yellow("%+v", cmt)
	ProcessDesc(cmt)
	//newThread := &model.Comment{
	//}
	//color.Red("%v", newThread)
	//localThread := &model.Comment{}
	//err = entity.GetDBInstance().Where("thread_id = ?", threadId).First(&localThread).Error
	//if errors.Is(err, gorm.ErrRecordNotFound) {
	//	//Only create when there is no such record with same threadID
	//	err = entity.GetDBInstance().Create(&newThread).Error
	//	if err != nil {
	//		logger.Errorln(err)
	//		return err
	//	}
	//	color.Cyan("Thread %d saved success %s\nTitle : %s", threadId, cmt.Link, cmt.Title)
	//} else {
	//	color.Red("Thread %d already exists!", threadId)
	//	color.Red("Link : %s", cmt.Link)
	//}
	return nil
}

func ProcessDesc(cmt *entity.Comment) {
	desc := cmt.Desc
	color.Red("Desc : %s", desc)
	res := strings.Split(desc, "·")
	color.Cyan("Res %+v", res)
}
