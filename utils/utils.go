package utils

import (
	"fmt"
	"strings"
)

//"https://voz.vn/f/chuyen-tro-linh-tinh.17"
//https://voz.vn/f/chuyen-tro-linh-tinh.17/page-2
func AddPageSuffix(URL string, page uint64) string {
	result := ""
	if page <= 1 {
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
