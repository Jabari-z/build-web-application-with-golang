package router

import (
	"github.com/gin-gonic/gin"
	"go-gin-web/router/middleware"
	"go-gin-web/service"
	"net/http"

)

func InitRouter(g *gin.Engine)  {
	middlewares := []gin.HandlerFunc{}
	// Middlewares.
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(middlewares...)

	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	// The health check handlers
	router := g.Group("/user")
	{
		router.POST("/addUser", service.AddUser)					//添加用户
		router.POST("/selectUser", service.SelectUser)			//查询用户
	}
}