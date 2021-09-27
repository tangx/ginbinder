package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tangx/ginbinder"
)

type paramsMap struct {
	Data map[string]interface{} `body:"body" mime:"json"`
}

func handlerBindMapAsJson(c *gin.Context) {
	parmas := &paramsMap{}

	err := ginbinder.ShouldBindRequest(c, parmas)
	if err != nil {
		panic(err)
	}

	c.JSON(200, parmas)
}
