package router

import (
	"github.com/gin-gonic/gin"
	"web_api/userop_web/api/message"
	"web_api/userop_web/middlewares"
)

func InitMessageRouter(Router *gin.RouterGroup) {
	//MessageRouter := Router.Group("message").Use(middlewares.JWTAuth())
	MessageRouter := Router.Group("message").Use(middlewares.SetUserId())
	{
		MessageRouter.GET("", message.List) // 轮播图列表页
		MessageRouter.POST("", message.New) //新建轮播图
	}
}
