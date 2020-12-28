package entity

type Thread struct {
	Title string `selector:"a"`
	Link  string `selector:"a[href]" attr:"href"`
}
