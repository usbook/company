package routers

import (
	"Teach/controllers"
	_ "Teach/docs"
	"Teach/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	route := gin.New()
	route.Use(gin.Logger())
	route.Use(gin.Recovery())
	//设置跨域
	route.Use(middlewares.Cors())
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiv1 := route.Group("/api/v1")
	{
		apiv1.POST("/add_tag", controllers.AddTag)
		apiv1.POST("/edit_tag", controllers.EditTag)
	}

	return route
}
