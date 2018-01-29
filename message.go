package main

import (
	"statnlp-initializer-service-go/app"
	"strconv"

	"github.com/goadesign/goa"
)

// MessageController implements the message resource.
type MessageController struct {
	*goa.Controller
}

// NewMessageController creates a message controller.
func NewMessageController(service *goa.Service) *MessageController {
	return &MessageController{Controller: service.NewController("MessageController")}
}

// List runs the list action.
func (c *MessageController) List(ctx *app.ListMessageContext) error {
	// MessageController_List: start_implement

	// Put your logic here

	res := loadDefaultData()

	return ctx.OK(res)
	// MessageController_List: end_implement
}

func loadDefaultData() app.MessageCollection {
	var slice app.MessageCollection
	for i := 0; i < 8; i++ {

		title := "title" + strconv.Itoa(i)
		content := "content" + strconv.Itoa(i)
		message := app.Message{Title: &title, Content: &content}
		slice = append(slice, &message)
	}
	return slice
}
