package middlewares

import (
	"ExchangeApp/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 检查请求头是否包含 Authorization 字段
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(401, gin.H{"error": "Authorization header is required"})
			ctx.Abort()
			return
		}
		// 解析 JWT 令牌
		username, err := utils.ParseJWT(authHeader)
		if err != nil {
			ctx.JSON(401, gin.H{"error": "Invalid token"})
			ctx.Abort()
			return
		}
		// 将用户名设置到请求上下文中
		ctx.Set("username", username)
		ctx.Next()
	}
}
