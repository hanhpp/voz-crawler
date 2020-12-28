package entity

type Comment struct {
	//ThreadId uint64 `selector:""`
	//UserName string `selector:""`
	Desc       string `selector:"data-lb-caption-desc"`
	Text       string `selector:".bbWrapper"`
	Name       string
	TimePosted string
	CommentId  string
	ThreadId   string
	//Title string `selector:"a"`
	//Link  string `selector:"a[href]" attr:"href"`
}
