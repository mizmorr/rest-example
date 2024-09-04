package server

import (
	"github.com/gin-gonic/gin"
	"github.com/mizmorr/rest-example/internal/controller"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewRouter(handler *gin.Engine, c *controller.UserController) {
	handler.Use(gin.Recovery())
	handler.Use(gin.Logger())
	handler.GET("/metrics", gin.WrapH(promhttp.Handler()))
	v1 := handler.Group("/v1")
	{
		user_routes := v1.Group("/user")
		user_routes.GET("/:id", c.Get)
	}
}
