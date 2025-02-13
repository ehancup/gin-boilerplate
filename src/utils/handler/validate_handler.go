package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)


func translateError(err error, trans ut.Translator) (errs []string) {
	if err == nil {
		return nil
	}
	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		translatedErr := fmt.Errorf("%s", e.Translate(trans)).Error()
		errs = append(errs, translatedErr)
	}
	return errs
}

func GetBody[T any](ctx *gin.Context) T {
	var b T

	if err := ctx.ShouldBindJSON(&b); err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"message": err.Error(),
			"success": false,
		})
	}
	
	v := validator.New()
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	_ = en_translations.RegisterDefaultTranslations(v, trans)

	if err := v.Struct(b); err != nil {
		errors := translateError(err, trans)
		ctx.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"message": "invalid validation!",
			"success": false,
			"data":    errors,
		})
	}

	return b
}
