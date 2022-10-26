package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/struct/request/:name", handlerStructRequest)
	r.POST("/struct/nested/:name/:age", handlerStructNested)
	r.GET("/struct/nested/:name/:age", handlerStructNested)
	r.POST("/mime/json/:name", handlerJson)
	r.GET("/mime/json/:name", handlerJson)
	r.POST("/mime/postform/:name", handlePostForm)
	r.GET("/mime/postform/:name", handlePostForm)

	r.GET("/map", handlerBindMapAsJson)
	_ = r.Run(":9881")
}
