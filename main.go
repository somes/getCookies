package main

import (
	"getCookies/router"
	"getCookies/util"
	"github.com/gin-gonic/gin"
)

var conf = util.GetConfig()

var port = conf.SERVER.SERVER_PORT

func main() {
	gin.SetMode(gin.ReleaseMode)
	//gin.DefaultWriter = ioutil.Discard
	r := gin.Default()
	r = router.CollectRoute(r)
	if len(port) != 0 {
		panic(r.Run(":" + port))
	}
	panic(r.Run())
}
