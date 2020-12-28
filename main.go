package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/gocolly/colly"
	"log"
	"os"
	"path/filepath"
	"strings"
	"voz/entity"
	"voz/global"
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
		color.Red("%v", thread)
	})
	_ = c.Visit(URL)
}

//standardizeSpaces remove redundancies spaces
func standardizeSpaces(s string) string {
	//fields return splitted array of chars if function satisfy
	return strings.Join(strings.Fields(s), " ")
}
