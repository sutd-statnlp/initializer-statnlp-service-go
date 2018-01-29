package service

import (
	"strconv"

	"../model"
)

// MessageService .
type MessageService struct {
}

// GetMessages  .
func (messageService MessageService) GetMessages() []model.Message {
	var slice []model.Message
	for i := 0; i < 8; i++ {
		title := "title" + strconv.Itoa(i)
		content := "content" + strconv.Itoa(i)
		message := model.Message{Title: title, Content: content}
		slice = append(slice, message)
	}
	return slice
}
