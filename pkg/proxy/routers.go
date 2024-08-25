package proxy

import "github.com/gin-gonic/gin"

func InstallRouter(r *gin.Engine) {
	InstallProxyRouter(r)
}

func InstallProxyRouter(r *gin.Engine) {
	p := NewProxy()
	v1 := r.Group("/v1")
	v1.Any("/*relativePath", p.ChatProxy)
}
