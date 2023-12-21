package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tangx/ginbinder"
)

type paramsDemo1 struct {
	Name          string `path:"name"`
	Age           int    `query:"age,default=18"`
	Money         int32  `query:"money" binding:"required"`
	Authorization string `cookie:"Authorization"`
	UserAgent     string `header:"User-Agent"`
	Data          Data   `body:"body"` // Data **不能** 指针类型
}

type Data struct {
	Replicas *int32   `json:"replicas" yaml:"replicas" xml:"replicas" form-urlencoded:"replicas"`
	Students Students `json:"students"`
}

type Students struct {
	Zhangsan Student `json:"zhangsan"`
	Lisi     Student `json:"lisi`
}

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func handlerStructRequest(c *gin.Context) {
	var err error
	params := &paramsDemo1{}

	err = ginbinder.ShouldBindRequest(c, params)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	// cc := c.Copy()
	multipleReadBody(c)
	c.JSON(200, params)
}

func multipleReadBody(c *gin.Context) {

	b, err := io.ReadAll(c.Request.Body)
	if err != nil {
		panic(err)
	}
	defer c.Request.Body.Close()

	fmt.Printf("Multiple Read Body: %s\n", b)
}
