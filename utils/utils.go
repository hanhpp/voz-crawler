package utils

import (
	"fmt"
	"strings"
)

//"https://voz.vn/f/chuyen-tro-linh-tinh.17"
//https://voz.vn/f/chuyen-tro-linh-tinh.17/page-2
func AddPageSuffix(URL string, page uint64) string {
	result := ""
	if page < 2 {
		result = URL
		return result
	}
	if strings.HasSuffix(URL, "/") {
		result = URL + "page-" + fmt.Sprintf("%d",page)
	} else {
		result = URL + "/page-" + fmt.Sprintf("%d",page)
	}
	return result
}

func GetCronString(time uint64) string {
	res :=fmt.Sprintf("@every 0h0m%ds",time)
	//color.Red("CronStr : %s",res)
	return res
}