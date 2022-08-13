package initialize

import (
	"github.com/gin-gonic/gin"

	"web_api/userop_web/middlewares"
	"web_api/userop_web/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()

	//配置跨域
	Router.Use(middlewares.Cors())
	ApiGroup := Router.Group("/up/v1")

	router.InitAddressRouter(ApiGroup)
	router.InitMessageRouter(ApiGroup)
	router.InitUserFavRouter(ApiGroup)

	return Router
}
