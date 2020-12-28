package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"voz/config"
	"voz/entity"
	"voz/global"
	"voz/model"
)

func createFile(name string) *os.File {
	f, err := os.Create(name + ".txt")

	if err != nil {
		log.Fatal(err)
	}

	return f
}
func main() {
	global.FetchEnvironmentVariables()
	entity.InitializeDatabaseConnection()
	entity.ProcessMigration()

	visitAndCollectFromURL(global.F17, "title")
	fmt.Println("done")

}

func visitAndCollectFromURL(URL string, fileName string) {
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
		text := standardizeSpaces(e.Text)
		titles = append(titles, text)
		thread := &entity.Thread{}
		err = e.Unmarshal(thread)
		if err != nil {
			log.Fatal(err)
		}
		threadId := GetThreadID(thread.Link)
		newThread := &model.Thread{
			Title:    thread.Link,
			Link:     thread.Title,
			ThreadId: threadId,
		}
		//color.Red("%v", newThread)
		err = entity.GetDBInstance().Debug().Create(&newThread).Error
		logger := config.GetLogger()
		if err != nil {
			logger.Errorln(err)
			return
		}
	})
	_ = c.Visit(URL)
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
