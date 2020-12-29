package utils

import "strings"

//"https://voz.vn/f/chuyen-tro-linh-tinh.17"
//https://voz.vn/f/chuyen-tro-linh-tinh.17/page-2
func AddPageSuffix(URL string, page uint64) string {
	if page <= 1 {
		return URL
	}
	if strings.HasSuffix(URL, "/") {
		return URL + "page-" + string(page)
	} else {
		return URL + "/page-" + string(page)
	}
}
