package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tangx/ginbinder"
)

type paramsDemo1 struct {
	Name          string `uri:"name"`
	Age           int    `query:"age,default=18"`
	Money         int32  `query:"money" binding:"required"`
	Authorization string `cookie:"Authorization"`
	UserAgent     string `header:"User-Agent"`
	Data          struct {
		Replicas *int32 `json:"replicas" yaml:"replicas" xml:"replicas" form:"replicas"`
	} `body:"body"`
}

func handlerDemo1(c *gin.Context) {
	var err error
	params := &paramsDemo1{}

	err = ginbinder.ShouldBindRequest(c, params)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(200, params)
}
