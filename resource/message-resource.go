package resource

import (
	"github.com/gin-gonic/gin"

	"../service"
)

type MessageResource struct {
}

func (messageResource MessageResource) InitRoutes(router *gin.Engine) {
	messageService := service.MessageService{}
	router.GET("/api/messages", func(context *gin.Context) {
		body := messageService.GetMessages()
		context.JSON(200, body)
	})
}
