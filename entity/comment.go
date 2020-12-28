package entity

type Comment struct {
	ThreadId uint64 `json:"thread_id,omitempty"`
	UserName string `json:"user_name,omitempty"`
	Text     string `json:"text,omitempty"`
}
