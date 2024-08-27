package proxy

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mmungdong/chatgpt-web/cmd/proxy/options"
)

// 鉴权中间件
func InstallAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 鉴权逻辑
		authorization := ctx.Request.Header.Get("Authorization")
		confAuthorization := fmt.Sprintf("Bearer %s", options.GetConfig().Http.AccessKey)
		if authorization != confAuthorization {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// 鉴权通过后设置请求头信息为真正的请求apikey
		ctx.Request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", options.GetConfig().Chat.APIKey))

		// 执行方法
		ctx.Next()
	}
}
