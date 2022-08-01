package initialize

import (
	"github.com/gin-gonic/gin"

	"web_api/goods_web/middlewares"
	"web_api/goods_web/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()

	//配置跨域
	Router.Use(middlewares.Cors())
	ApiGroup := Router.Group("/g/v1")

	router.InitGoodsRouter(ApiGroup)

	return Router
}
