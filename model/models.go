package model

type Thread struct {
	Title    string `json:"title,omitempty"`
	Link     string `json:"link,omitempty"`
	ThreadId uint64 `json:"thread_id,omitempty" gorm:"primary_key;unique;not null;index"`
	ParentURL string `json:"parent_url,omitempty"`
	BaseModel
}

type DeletedThread struct {
	Thread
}

type Comment struct {
	ThreadId   uint64 `json:"thread_id,omitempty"`
	CommentId  uint64 `json:"post_id,omitempty" gorm:"primary_key;unique;not null;index"`
	Text       string `json:"text,omitempty"`
	UserName   string `json:"user_name,omitempty"`
	TimePosted string `json:"time_posted,omitempty"`
	BaseModel
}
