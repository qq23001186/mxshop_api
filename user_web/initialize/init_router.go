package initialize

import (
	"github.com/gin-gonic/gin"

	"web_api/user_web/middlewares"
	"web_api/user_web/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()

	//配置跨域
	Router.Use(middlewares.Cors())
	ApiGroup := Router.Group("/u/v1")

	router.InitUserRouter(ApiGroup)
	router.InitBaseRouter(ApiGroup)
	return Router
}
