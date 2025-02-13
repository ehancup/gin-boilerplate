package example

import "github.com/gin-gonic/gin"

func InitRoutes(app *gin.RouterGroup, service *Service) {
	app.GET("/say-hello", service.SayHello)
	app.POST("/payload", service.GetPayload)
}