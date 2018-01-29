package model

// Message model
type Message struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
}
