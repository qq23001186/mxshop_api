package router

import (
	"github.com/gin-gonic/gin"
	"web_api/userop_web/api/address"
	"web_api/userop_web/middlewares"
)

func InitAddressRouter(Router *gin.RouterGroup) {
	//AddressRouter := Router.Group("address").Use(middlewares.JWTAuth())
	AddressRouter := Router.Group("address").Use(middlewares.SetUserId())
	{
		AddressRouter.GET("", address.List)
		AddressRouter.DELETE("/:id", address.Delete)
		AddressRouter.POST("", address.New)
		AddressRouter.PUT("/:id", address.Update)
	}
}
