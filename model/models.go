package model

type Thread struct {
	Title string `json:"title,omitempty"`
	Link  string `json:"link,omitempty"`
}

type Comment struct {
	ThreadId uint64 `json:"thread_id,omitempty"`
	UserName string `json:"user_name,omitempty"`
	Text     string `json:"text,omitempty"`
}