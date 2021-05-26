package main

import (
	"github.com/gin-gonic/gin"
	"jxbdproject/controller"
	"jxbdproject/router"
)

func main() {
	r := gin.Default()
	go controller.TickerHandler() //这个定时器用来同步redis中的点赞数量
	r = router.InitRouter(r)
	r.Run(":8080")
}
