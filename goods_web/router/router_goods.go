package router

import (
	"github.com/gin-gonic/gin"
	"web_api/goods_web/api/goods"
	"web_api/goods_web/middlewares"
)

func InitGoodsRouter(Router *gin.RouterGroup) {
	//GoodsRouter := Router.Group("goods").Use(middlewares.JWTAuth()).Use(middlewares.Trace())
	GoodsRouter := Router.Group("goods").Use(middlewares.Trace())
	{
		GoodsRouter.GET("", goods.List) //商品列表
		//GoodsRouter.POST("", middlewares.JWTAuth(), middlewares.IsAdminAuth(), goods.New) //该接口需要管理员权限
		GoodsRouter.POST("", goods.New)       // 我们测试先将这2个接口权限关闭掉
		GoodsRouter.GET("/:id", goods.Detail) //获取商品的详情
		//GoodsRouter.DELETE("/:id",middlewares.JWTAuth(), middlewares.IsAdminAuth(), goods.Delete) //删除商品
		GoodsRouter.DELETE("/:id", goods.Delete)     // 我们测试先将这2个接口权限关闭掉
		GoodsRouter.GET("/:id/stocks", goods.Stocks) //获取商品的库存
		//GoodsRouter.PUT("/:id",middlewares.JWTAuth(), middlewares.IsAdminAuth(), goods.Update)
		GoodsRouter.PUT("/:id", goods.Update)
		//GoodsRouter.PATCH("/:id",middlewares.JWTAuth(), middlewares.IsAdminAuth(), goods.UpdateStatus)
		GoodsRouter.PATCH("/:id", goods.UpdateStatus)
	}
}
