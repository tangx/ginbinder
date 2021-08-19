package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tangx/ginbinder"
)

type paramsDemo2 struct {
	Uris struct {
		Name string `uri:"name"`
		Age  int    `uri:"age"`
	}
	Queries struct {
		Age   int   `query:"age,default=18"`
		Money int32 `query:"money" binding:"required"`
	}
	Cookies struct {
		Authorization string `cookie:"Authorization"`
	}
	Headers struct {
		UserAgent string `header:"User-Agent"`
	}
	BodyData struct {
		Replicas *int32 `json:"replicas" yaml:"replicas" xml:"replicas" form:"replicas"`
	} `body:"body"`
}

func handlerDemo2(c *gin.Context) {
	var err error
	params := &paramsDemo2{}

	err = ginbinder.ShouldBindRequest(c, params)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(200, params)
}
