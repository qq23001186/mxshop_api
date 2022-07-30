package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"web_api/user_web/api"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	//UserRouter := Router.Group("user").Use(middlewares.JWTAuth()) //给整个user都添加登录验证
	zap.S().Info("配置用户相关的url")
	{
		//UserRouter.GET("list", middlewares.JWTAuth(), middlewares.IsAdminAuth(), api.GetUserList)
		UserRouter.GET("list", api.GetUserList) // 为了测试方便，去掉jwt和管理员的中间件
		UserRouter.POST("pwd_login", api.PassWordLogin)
		UserRouter.POST("register", api.Register)
	}
}
