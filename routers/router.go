package routers

import (
	"Teach/controllers"
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
		//获取标签列表
		apiv1.GET("/edit_tags", controllers.EditTag)
	}

	return route
}
