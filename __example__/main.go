package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/demo1/:name", handlerDemo1)
	r.POST("/demo2/:name/:age", handlerDemo2)
	r.GET("/demo2/:name/:age", handlerDemo2)
	r.POST("/demo3/:name", handlerDemo3)
	r.GET("/demo3/:name", handlerDemo3)

	r.GET("/map", handlerBindMapAsJson)
	_ = r.Run(":9881")
}
