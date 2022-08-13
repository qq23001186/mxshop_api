package middlewares

import (
	"github.com/gin-gonic/gin"
	"web_api/order_web/models"
)

func SetUserId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var userId uint = 1
		ctx.Set("claims", &models.CustomClaims{
			AuthorityId: 1,
		})
		ctx.Set("userId", userId)
		ctx.Next()
	}
}
