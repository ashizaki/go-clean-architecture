package router

import (
	"github.com/ashizaki/go-clean-architecture/interface/controller"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(g *gin.RouterGroup, contoller controller.UserController) {
	g.POST("/users", func(c *gin.Context) { contoller.CreateUser(&context{base: c}) })
	g.GET("/users", func(c *gin.Context) { contoller.ListUsers(&context{base: c}) })
	g.GET("/users/:id", func(c *gin.Context) { contoller.GetUser(&context{base: c}) })
	g.PUT("/users/:id", func(c *gin.Context) { contoller.UpdateUser(&context{base: c}) })
	g.DELETE("/users/:id", func(c *gin.Context) { contoller.DeleteUser(&context{base: c}) })
}
