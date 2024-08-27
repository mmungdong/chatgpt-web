package proxy

import "github.com/gin-gonic/gin"

func InstallRouter(r *gin.Engine) {
	InstallProxyRouter(r)
}

func InstallProxyRouter(r *gin.Engine) {
	p := NewProxy()
	v1 := r.Group("/v1")

	// 安装中间件
	v1.Use(InstallAuthMiddleware())

	// 设置路由
	v1.Any("/*relativePath", p.ChatProxy)
}
