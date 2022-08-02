package router

import (
	"github.com/gin-gonic/gin"
	"web_api/goods_web/api/goods"
)

func InitGoodsRouter(Router *gin.RouterGroup) {
	GoodsRouter := Router.Group("goods")
	{
		GoodsRouter.GET("", goods.List) //商品列表
		//GoodsRouter.POST("", middlewares.JWTAuth(), middlewares.IsAdminAuth(), goods.New) //该接口需要管理员权限
		GoodsRouter.POST("", goods.New) // 我们测试先将这2个接口权限关闭掉
	}
}
