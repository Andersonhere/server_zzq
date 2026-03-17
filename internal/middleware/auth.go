package middleware

import (
	"net/http"
	"strings"

	"server_zzq/internal/utils"

	"github.com/gin-gonic/gin"
)

// Auth JWT 鉴权中间件
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, utils.Unauthorized("未登录或 token 已失效"))
			c.Abort()
			return
		}

		// 提取 token
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, utils.Unauthorized("token 格式错误"))
			c.Abort()
			return
		}

		tokenString := parts[1]
		claims, err := utils.ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, utils.Unauthorized("token 无效或已过期"))
			c.Abort()
			return
		}

		// 将商家 ID 存入上下文
		c.Set("shopID", claims.ShopID)
		c.Next()
	}
}
