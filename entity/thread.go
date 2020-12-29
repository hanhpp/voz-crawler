package entity

//type Thread struct {
//	Title    string `selector:"a"`
//	Link     string `selector:"a[href]" attr:"href"`
//	PageJump string `selector:".structItem-pageJump"`
//}

type Thread struct {
	Title    string `selector:".structItem-title > a"`
	Link     string `selector:".structItem-title > a[href]" attr:"href"`
	PageJump []string `selector:".structItem-pageJump > a"`
}
