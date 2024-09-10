package router

import (
	"github.com/gin-gonic/gin"
	_ "github.com/mizmorr/rest-example/docs"

	"github.com/mizmorr/rest-example/internal/controller"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(handler *gin.Engine, c *controller.UserController) {
	handler.Use(gin.Recovery())
	handler.Use(gin.Logger())
	handler.GET("/metrics", gin.WrapH(promhttp.Handler()))
	v1 := handler.Group("/v1")
	{
		user_routes := v1.Group("/user")
		user_routes.GET("/:id", c.Get)
		user_routes.POST("/create", c.Create)
	}
	handler.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
