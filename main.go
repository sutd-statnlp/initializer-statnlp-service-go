package main

import (
	"github.com/gin-gonic/gin"

	"./controller"
	"./resource"
)

func main() {
	router := setupRoutes()

	router.Run(":8080")
}

func setupRoutes() *gin.Engine {
	router := gin.Default()

	homeController := controller.HomeController{}
	homeController.InitRoutes(router)

	messageResource := resource.MessageResource{}
	messageResource.InitRoutes(router)

	return router
}
