package utils

import (
	"fmt"
	"regexp"
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

func RemoveRedundantSpaces(str string) string {
	//https://stackoverflow.com/questions/37290693/how-to-remove-redundant-spaces-whitespace-from-a-string-in-golang
	re_leadclose_whtsp := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
	re_inside_whtsp := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
	final := re_leadclose_whtsp.ReplaceAllString(str, "")
	final = re_inside_whtsp.ReplaceAllString(final, " ")
	fmt.Println(final)
	return final
}

//StandardizeSpaces remove redundancies spaces
func StandardizeSpaces(s string) string {
	//fields return splitted array of chars if function satisfy
	return strings.Join(strings.Fields(s), " ")
}