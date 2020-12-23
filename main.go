package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/gocolly/colly"
	"log"
	"os"
	"strings"
)

const (
	f17         = "https://voz.vn/f/chuyen-tro-linh-tinh.17/"
	threadLink  = ".structItem-minor .structItem-parts"
	threadTitle = ".structItem-title"
)

//filterSpace := "([A-Za-z0-9àáãạảăắằẳẵặâấầẩẫậèéẹẻẽêềếểễệđìíĩỉịòóõọỏôốồổỗộơớờởỡợùúũụủưứừửữựỳỵỷỹýÀÁÃẠẢĂẮẰẲẴẶÂẤẦẨẪẬÈÉẸẺẼÊỀẾỂỄỆĐÌÍĨỈỊÒÓÕỌỎÔỐỒỔỖỘƠỚỜỞỠỢÙÚŨỤỦƯỨỪỬỮỰỲỴỶỸÝ.,])\\w+"
//for i, t := range titles {
//	match, err := regexp.MatchString(filterSpace, t)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(match)
//}

type Thread struct {
	Title string `selector:"a"`
	Link  string `selector:"a[href]" attr:"href"`
}

func createFile(name string) *os.File {
	f, err := os.Create(name + ".txt")

	if err != nil {
		log.Fatal(err)
	}

	return f
}
func main() {
	c := colly.NewCollector()

	title := createFile("title")
	defer title.Close()

	var titles []string
	c.OnHTML(threadTitle, func(e *colly.HTMLElement) {
		_, err := title.Write([]byte(e.Text))
		if err != nil {
			log.Fatal(err)
		}
		text := standardizeSpaces(e.Text)
		titles = append(titles, text)
		thread := &Thread{}
		err = e.Unmarshal(thread)
		if err != nil {
			log.Fatal(err)
		}
		color.Red("%v", thread)
	})
	_ = c.Visit(f17)
	fmt.Println("done")

}

//fields return splitted array of chars if function satisfy
func standardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
