package routers

import (
	"adminVideos/routers/api"
	"adminVideos/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	//静态文件服务 提供图片路径
	r.Static("/images","./public/upload/images")
	r.Static("/static/images","./public/images")
	r.Static("/videos","./public/upload/videos")
	r.Static("/index.html","./public/dist")
	
	//登录模块
	api1 := r.Group("/api/v1")
	{
		api1.POST("/login",v1.Login)
		api1.POST("/registered",v1.Registered)
	}

	//用户管理模块
	api2 := r.Group("/api/v2")
	api2.Use(api.GetAuth)
	{

	}




return r
}
