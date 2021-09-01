package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tangx/ginbinder"
)

type paramsDemo3 struct {
	Auth struct {
		Name          string `uri:"name"`
		Authorization string `cookie:"Authorization"`
	}
	BodyData struct {
		Replicas *int32 `json:"replicas" yaml:"replicas" xml:"replicas" form:"replicas"`
	} `body:"body" mime:"json"`
}

func handlerDemo3(c *gin.Context) {
	var err error
	params := &paramsDemo3{}

	err = ginbinder.ShouldBindRequest(c, params)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(200, params)
}
