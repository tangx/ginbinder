package ginbinder

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tangx/ginbinder/binding"
)

// BindRequest binds the passed struct pointer using binding.Request.
// It will abort the request with HTTP 400 if any error occurs.
func BindRequest(c *gin.Context, obj interface{}) error {
	if err := ShouldBindRequest(c, obj); err != nil {
		c.AbortWithError(http.StatusBadRequest, err).SetType(gin.ErrorTypeBind)
		return err
	}
	return nil
}

// BindOnlyRequest is BindRequest is without validate
func BindOnlyRequest(c *gin.Context, obj interface{}) error {
	if err := ShouldBindOnlyRequest(c, obj); err != nil {
		c.AbortWithError(http.StatusBadRequest, err).SetType(gin.ErrorTypeBind)
		return err
	}
	return nil
}

// ShouldBindRequest binds the passed struct pointer using the specified binding engine.
//
//	including
//	  `path`,
//	  `query`
//	  `header` and
//	  `body data` with tag `body:"body"`
//	and it's decoder is decided by header `Content-Type` value
//
//	type Params struct {
//		Name          string `path:"name"`
//		Age           int    `query:"age,default=18"`
//		Money         int32  `query:"money" binding:"required"`
//		Authorization string `cookie:"Authorization"`
//		UserAgent     string `header:"User-Agent"`
//		Data          struct {
//			Replicas *int32 `json:"replicas" yaml:"replicas" xml:"replicas" form:"replicas"`
//		} `body:"body"`
//	}
func ShouldBindRequest(c *gin.Context, obj interface{}) error {
	params := make(map[string][]string)
	for _, v := range c.Params {
		params[v.Key] = []string{v.Value}
	}

	return binding.Request.Bind(obj, c.Request, params)
}

// ShouldBindOnlyRequest is ShouldBindRequest is without validate
func ShouldBindOnlyRequest(c *gin.Context, obj interface{}) error {
	params := make(map[string][]string)
	for _, v := range c.Params {
		params[v.Key] = []string{v.Value}
	}

	return binding.Request.BindOnly(obj, c.Request, params)
}
