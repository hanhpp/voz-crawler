package main

import (
	"github.com/fatih/color"
	"github.com/gocolly/colly"
	"voz/global"
)

func main() {
	url := "https://voz.vn/f/chuyen-tro-linh-tinh.17/"
	VisitAndCollectThreadsFromURL(url)
}

func test() {
	//input := "   Text   More here     "
	//re_leadclose_whtsp := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
	//re_inside_whtsp := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
	//final := re_leadclose_whtsp.ReplaceAllString(input, "")
	//final = re_inside_whtsp.ReplaceAllString(final, " ")
	//fmt.Println(final)
	//global.FetchEnvironmentVariables()
	//entity.InitializeDatabaseConnection()
	//entity.ProcessMigration()
	//url := "https://voz.vn/t/dan-ong-thi-cai-gi-quan-trong-nhat.244714"
	//cronjob.CrawlComments(url,244714)
}
func VisitAndCollectThreadsFromURL(URL string) {
	c := colly.NewCollector()
	//var titles []string
	c.OnHTML(global.ThreadStruct, func(e *colly.HTMLElement) {
		//err := handleThreadContent(e, titles, URL)
		//logger := config.GetLogger()
		//if err != nil {
		//	logger.Errorln(err)
		//}
		color.Red(e.Text)
	})
	_ = c.Visit(URL)
}
