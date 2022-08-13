package router

import (
	"github.com/gin-gonic/gin"
	"web_api/order_web/api/order"
	"web_api/order_web/api/pay"
	"web_api/order_web/middlewares"
)

func InitOrderRouter(Router *gin.RouterGroup) {
	OrderRouter := Router.Group("orders").Use(middlewares.SetUserId())
	{
		OrderRouter.GET("", order.List) // 订单列表
		//BannerRouter.GET("", middlewares.JWTAuth(), middlewares.IsAdminAuth(), order.List) // 订单列表
		OrderRouter.POST("", order.New)        // 新建订单
		OrderRouter.GET("/:id/", order.Detail) // 订单详情
	}
	PayRouter := Router.Group("pay")
	{
		PayRouter.POST("alipay/notify", pay.Notify)
	}
}
