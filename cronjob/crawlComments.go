package cronjob

import (
	"errors"
	"github.com/fatih/color"
	"github.com/gocolly/colly"
	"gorm.io/gorm"
	"log"
	"regexp"
	"strconv"
	"strings"
	"voz/config"
	"voz/entity"
	"voz/global"
	"voz/model"
	"voz/utils"
)

func CrawlComments(url string, threadID uint64) {
	color.Yellow("cron job: Crawling comments from %s", url)
	//Let's crawl from multiple pages
	for i := uint64(2); i < 5; i++ {
		newURL := utils.AddPageSuffix(url, i)
		VisitAndCollectCommentFromURL(newURL, threadID, i)
	}
}

func VisitAndCollectCommentFromURL(URL string, threadID uint64, page uint64) {
	c := colly.NewCollector()

	var titles []string
	c.OnHTML(global.CommentStruct, func(e *colly.HTMLElement) {
		err := handleCommentContent(e, titles, threadID, page)
		logger := config.GetLogger()
		if err != nil {
			logger.Errorln(err)
		}
	})
	_ = c.Visit(URL)
}

func handleCommentContent(e *colly.HTMLElement, titles []string, threadID uint64, page uint64) error {
	logger := config.GetLogger()
	text := standardizeSpaces(e.Text)
	titles = append(titles, text)
	localCmt := ProcessDesc(e, threadID, page)

	existedCmt := &model.Comment{}
	err := entity.GetDBInstance().Where("comment_id = ?", localCmt.CommentId).First(&existedCmt).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		//Only create when there is no such record with same threadID
		err = entity.GetDBInstance().Create(&localCmt).Error
		if err != nil {
			logger.Errorln(err)
			return err
		}
		color.Green("[%d] Comment %d by user %s saved success!", localCmt.ThreadId, localCmt.CommentId, localCmt.UserName)
		logger.WithField("text",localCmt.Text).WithField("threadID",localCmt.ThreadId).
			WithField("commentId",localCmt.CommentId).WithField("username",localCmt.UserName).
			Info("Comment saved success!")
		//color.Blue("Content [%s]", localCmt.Text)
	} else {
		//color.Red("Comment %d by user %s already exists!", localCmt.CommentId, localCmt.UserName)
		//color.Red("Content: \n%s",localCmt.Text)
		//color.Red("Link : %s", cmt.Link)
	}
	return nil
}

func GetCmtId(cmtId string) string {
	r := regexp.MustCompile(global.CommentIDRegex)
	res := r.FindAllString(cmtId, -1)
	if len(res) == 1 {
		return res[0]
	} else {
		color.Red("Too many matched %+v", res)
		return ""
	}
}

func ProcessDesc(e *colly.HTMLElement, threadId uint64, page uint64) *model.Comment {
	cmt := &entity.Comment{}
	logger := config.GetLogger()
	err := e.Unmarshal(cmt)
	if err != nil {
		log.Fatal(err)
	}
	cmt.Desc = e.Attr(global.CommentNamespace)
	cmt.CommentId = e.Attr(global.CommentId)
	cmtIdStr := GetCmtId(cmt.CommentId)
	cmtId, err := strconv.ParseUint(cmtIdStr, 10, 64)
	if err != nil {
		logger.Errorln(err)
		return nil
	}
	desc := cmt.Desc
	res := strings.Split(desc, "·")
	for i, v := range res {
		res[i] = strings.TrimSpace(v)
	}
	if len(res) == 2 {
		cmt.Name = res[0]
		cmt.TimePosted = res[1]
	} else {
		color.Red("Res len is not 2 but %d", len(res))
	}


	localCmt := &model.Comment{
		ThreadId:   threadId,
		UserName:   cmt.Name,
		Text:       cmt.Text,
		TimePosted: cmt.TimePosted,
		CommentId:  cmtId,
	}
	if page <2 {
		localCmt.Page = 1
	} else {
		localCmt.Page = page
	}
	return localCmt
}
