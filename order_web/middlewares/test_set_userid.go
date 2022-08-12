package middlewares

import (
	"github.com/gin-gonic/gin"
)

func SetUserId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var userId uint = 1
		ctx.Set("userId", userId)
		ctx.Next()
	}
}
