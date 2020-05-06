package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type context struct {
	base *gin.Context
}

func (c *context) Query(s string) string {
	return c.base.Query(s)
}

func (c *context) Param(s string) string {
	return c.base.Param(s)
}

func (c *context) Bind(i interface{}) error {
	return c.base.Bind(i)
}

func (c *context) Status(i int) {
	c.base.Status(i)
}

func (c *context) JSON(i int, i2 interface{}) {
	c.base.JSON(i, i2)
}

func (c *context) Request() *http.Request {
	return c.base.Request
}
