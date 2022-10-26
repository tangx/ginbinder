package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tangx/ginbinder"
)

type paramsJson struct {
	Auth struct {
		Name          string `uri:"name"`
		Authorization string `cookie:"Authorization"`
	}
	BodyData struct {
		Replicas *int32 `json:"replicas" yaml:"replicas" xml:"replicas" form:"replicas"`
	} `body:"body" mime:"json"`
}

func handlerJson(c *gin.Context) {
	var err error
	params := &paramsJson{}

	err = ginbinder.ShouldBindRequest(c, params)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(200, params)
}
