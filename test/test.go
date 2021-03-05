package main

import (
	"github.com/fatih/color"
	"github.com/gocolly/colly"
	"strings"
	"voz/global"
)

func main() {
	//url := "https://voz.vn/f/chuyen-tro-linh-tinh.17/"
	//VisitAndCollectThreadsFromURL(url)
	testDel()
}

func testDel(){
	res := strings.Contains("{ \"lightbox_close\": \"Close\", \"lightbox_next\": \"Next\", \"lightbox_previous\": \"Previous\", \"lightbox_error\": \"The requested content cannot be loaded. Please try again later.\", \"lightbox_start_slideshow\": \"Start slideshow\", \"lightbox_stop_slideshow\": \"Stop slideshow\", \"lightbox_full_screen\": \"Full screen\", \"lightbox_thumbnails\": \"Thumbnails\", \"lightbox_download\": \"Download\", \"lightbox_share\": \"Share\", \"lightbox_zoom\": \"Zoom\", \"lightbox_new_window\": \"New window\", \"lightbox_toggle_sidebar\": \"Toggle sidebar\" } Bọn này chắc là tổ chức abcxyz vấn đề không gì đang nói nhưng tổ chức này đang bôi nhọ bộ mặt nước chúng ta cụ thể hơn là các nước Cộng Sản hoặc đối đầu chính trị trong có các anh Ấn độ trong khi tối ngày các anh toàn gửi hình cho Việt Nam chúng ta. Tôi biết là bọn này đang nói tới \"tự do dân chủ\"\nCụ thể hơn là các loser 3/ đang cố công kích Việt Nam anh em hãy vào report page và web twit vì đã kích chính trị cho chúng nó biết Vn có tự do dùng facebook instagram twiter bọn này láo hết sức", "The requested content cannot be loaded")
	color.Red("Res : %v",res)
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
