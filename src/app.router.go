package src

import (
	"gin-boilerplate/src/app/example"

	"github.com/gin-gonic/gin"
)

func InitRoutes(app *gin.Engine) {

	// change if you want to set base path
	route := app.Group("/") 

	example.InitRoutes(route, GetExampleServie())
}