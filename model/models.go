package model

import "gorm.io/gorm"

type Thread struct {
	Title    string `json:"title,omitempty"`
	Link     string `json:"link,omitempty"`
	ThreadId uint64 `json:"thread_id,omitempty" gorm:"primary_key;unique;not null;index"`
	BaseModel
}

type Comment struct {
	ThreadId uint64 `json:"thread_id,omitempty"`
	UserName string `json:"user_name,omitempty"`
	Text     string `json:"text,omitempty"`
	gorm.Model
}
