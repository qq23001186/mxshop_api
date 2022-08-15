package router

import (
	"github.com/gin-gonic/gin"
	"web_api/goods_web/api/banners"
	"web_api/goods_web/middlewares"
)

func InitBannerRouter(Router *gin.RouterGroup) {
	BannerRouter := Router.Group("banners").Use(middlewares.Trace())
	{
		BannerRouter.GET("", banners.List) // 轮播图列表页
		//BannerRouter.DELETE("/:id", middlewares.JWTAuth(), middlewares.IsAdminAuth(), banners.Delete) // 删除轮播图
		//BannerRouter.POST("", middlewares.JWTAuth(), middlewares.IsAdminAuth(), banners.New)          //新建轮播图
		//BannerRouter.PUT("/:id", middlewares.JWTAuth(), middlewares.IsAdminAuth(), banners.Update)    //修改轮播图信息
		BannerRouter.DELETE("/:id", banners.Delete) // 删除轮播图
		BannerRouter.POST("", banners.New)          //新建轮播图
		BannerRouter.PUT("/:id", banners.Update)    //修改轮播图信息
	}
}
