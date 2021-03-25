package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
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
	//fmt.Println(final)
	return final
}

//StandardizeSpaces remove redundancies spaces
func StandardizeSpaces(s string) string {
	//fields return splitted array of chars if function satisfy
	return strings.Join(strings.Fields(s), " ")
}


func BadRequest(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
}

func BadRequestWithMessage(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest, gin.H{"message": msg})
}

func InternalServerError(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{"message": "internal server error"})
}

func InternalServerErrorMsg(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest, gin.H{"message": msg})
}
func Ok(c *gin.Context, msg gin.H) {
	c.JSON(http.StatusOK, msg)
}

func OkMsg(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{"message": msg})
}

func UnauthorizedMsg(c *gin.Context, msg string) {
	c.JSON(http.StatusUnauthorized, gin.H{"message": msg})
}

func BadReqWithDetail(c *gin.Context,detail string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status_code": http.StatusBadRequest,
		"error_code":  400,
		"detail":      detail,
	})
}