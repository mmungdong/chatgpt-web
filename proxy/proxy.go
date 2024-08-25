package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mmungdong/chatgpt-web/cmd/proxy/config"
	"github.com/mmungdong/chatgpt-web/pkg/proxy"
)

func main() {
	config.InitProxyConfig()
	config := config.GetConfig()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	proxy.InstallRouter(r)
	log.Printf("starting server on %s:%d", config.Http.Host, config.Http.Port)
	r.Run(fmt.Sprintf("%s:%v", config.Http.Host, config.Http.Port))
	// health check
	// r.GET("/health", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"status":  "200",
	// 		"message": "ok",
	// 	})
	// })

	// apis := r.Group("/v1")
	// apis.POST("/*relativePath", func(ctx *gin.Context) {
	// 	OpenAIBaseURL := "https://api.openai-sb.com/v1"
	// 	upstreamURL, err := url.Parse(strings.TrimSuffix(OpenAIBaseURL, "/v1"))
	// 	if err != nil {
	// 		log.Fatalf("[PROXY ERROR]: %s", err)
	// 	}
	// 	// 路由转发
	// 	upstream := httputil.NewSingleHostReverseProxy(upstreamURL)
	// 	ctx.Request.Host = upstreamURL.Host
	// 	upstream.ServeHTTP(ctx.Writer, ctx.Request)
	// })

}
