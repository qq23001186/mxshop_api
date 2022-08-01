package router

import (
	"github.com/gin-gonic/gin"
	"web_api/goods_web/api/goods"
)

func InitGoodsRouter(Router *gin.RouterGroup) {
	GoodsRouter := Router.Group("goods")
	{
		GoodsRouter.GET("", goods.List) //商品列表
	}
}
