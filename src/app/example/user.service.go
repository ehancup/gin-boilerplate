package example

import (
	"gin-boilerplate/src/utils/handler"

	"github.com/gin-gonic/gin"
)

type Service struct{}

// SayHello godoc
//
//	@Summary		Say Hello
//	@Description	example. return "hello world!"
//	@Tags			example
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	string	"hello world!"
//	@Router			/say-hello [get]
func (Service) SayHello(c *gin.Context) {
	c.String(200, "hello world!")
}

type Payload struct {
	Email string `json:"email" validate:"required"`
	Name string `json:"name" validate:"min=4,required"`
}

// Payload godoc
//
//	@Summary		Accept Payload
//	@Description	example. getting value from json form
//	@Tags			example
//	@Accept			json
//	@Produce		json
//	@Param			payload	body		Payload	true	"payload"
//	@Success		200		{object}	Payload	
//	@Router			/payload [post]
func (Service) GetPayload(c *gin.Context) {

	// example of getting json body
	payload := handler.GetBody[Payload](c)

	c.JSON(200, payload)
}