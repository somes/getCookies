package controller

import (
	"github.com/gin-gonic/gin"
)

func GetCookies(ctx *gin.Context) {
	err := ctx.Request.Header.Write(ctx.Writer)
	if err != nil {
		ctx.Writer.WriteString("Error: " + err.Error())
		return
	}
	return
}
