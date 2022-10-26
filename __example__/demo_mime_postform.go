package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tangx/ginbinder"
)

type paramsPostFrom struct {
	Auth struct {
		Name          string `uri:"name"`
		Authorization string `cookie:"Authorization"`
	}
	BodyData struct {
		Replicas *int32 `json:"replicas" yaml:"replicas" xml:"replicas" form:"replicas"`
		Content  string `json:"content" yaml:"content" xml:"content" form:"content"`
	} `body:"body" mime:"formPost"`
}

func handlePostForm(c *gin.Context) {
	var err error
	params := &paramsPostFrom{}

	err = ginbinder.ShouldBindRequest(c, params)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(200, params)
}
