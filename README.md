# Gin-gonic Boilerplate

this Gin Boilerplate is inspirated by NestJS.

this boilerplate include package :
- Gin gonic
- Gorm
- Swaggo
- Logger
- Validator

## How to use it
1. Run
```bash
$ go mod tidy
```
2. make a `.env` file (you can copy all necessary variable in `.env.example`)
3. set the value of `.env`   
    note : if you want to use database. you can see [gorm documentation](https://gorm.io/docs/connecting_to_the_database.html) and set the `DB_DSN` value in `.env`, if you dont use it, just leave it same as in the `.env.example`.
4. If you use database, you need to uncomment code in `src/index.go` line 43-53


## Basic use flow

1. make new folder in `src/app` with name of your service. for the example, i name it `example` folder.
2. Make 2 file inside the folder. `{folder}.service.go` and `{folder}.router.go`. fill `{folder}` with your folder name. for me its `example.service.go` and `example.router.go`
3. in `*.service.go` type the service type and your function.
```go
package example

import  "github.com/gin-gonic/gin"

type Service struct{}

func (Service) SayHello(c *gin.Context) {
	c.String(200, "hello world!")
}
```
4. then, go to `*.router.go` and make func with name `InitRoutes` with 2 paarmeter, `*gin.RouterGroup` and your service, and set your route
```go
package example

import "github.com/gin-gonic/gin"

func InitRoutes(app *gin.RouterGroup, service *Service) {
	app.GET("/say-hello", service.SayHello)
}
```
5. Go to `src/app.service.go` and make a func to return your current service.
```go
package src

import "gin-boilerplate/src/app/example"

func GetExampleServie() *example.Service{
	return &example.Service{}
}
```
6. Go to `src/app.router.go` and call the `InitRoutes` func you just made.
```go
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
```
7. Run the `main.go` and go to `/say-hello`, and you can see the route you just made.


## Contact
- [Email](mailto:rhanysuf@gmail.com)