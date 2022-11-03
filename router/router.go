package router

import (
	"getCookies/controller"
	"github.com/gin-gonic/gin"
)

var r = gin.Default()

func CollectRoute() *gin.Engine {
	r.GET("/ping", func(c *gin.Context) { c.JSON(200, gin.H{"msg": "pong"}) })
	r.GET("/api/cookies", controller.GetCookies)
	return r
}
