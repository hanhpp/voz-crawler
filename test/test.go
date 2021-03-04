package main

import (
	"fmt"
	"regexp"
)

func main() {
	input := "   Text   More here     "
	re_leadclose_whtsp := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
	re_inside_whtsp := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
	final := re_leadclose_whtsp.ReplaceAllString(input, "")
	final = re_inside_whtsp.ReplaceAllString(final, " ")
	fmt.Println(final)
	//global.FetchEnvironmentVariables()
	//entity.InitializeDatabaseConnection()
	//entity.ProcessMigration()
	//url := "https://voz.vn/t/dan-ong-thi-cai-gi-quan-trong-nhat.244714"
	//cronjob.CrawlComments(url,244714)
}
