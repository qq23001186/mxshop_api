package router

import (
	"github.com/gin-gonic/gin"
	"web_api/userop_web/api/user_fav"
	"web_api/userop_web/middlewares"
)

func InitUserFavRouter(Router *gin.RouterGroup) {
	//UserFavRouter := Router.Group("userfavs").Use(middlewares.JWTAuth())
	UserFavRouter := Router.Group("userfavs").Use(middlewares.SetUserId())
	{
		UserFavRouter.DELETE("/:id", user_fav.Delete) // 删除收藏记录
		UserFavRouter.GET("/:id", user_fav.Detail)    // 获取收藏记录
		UserFavRouter.POST("", user_fav.New)          //新建收藏记录
		UserFavRouter.GET("", user_fav.List)          //获取当前用户的收藏
	}
}
