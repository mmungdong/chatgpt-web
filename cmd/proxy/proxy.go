package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mmungdong/chatgpt-web/cmd/proxy/options"
	"github.com/mmungdong/chatgpt-web/pkg/proxy"
)

func main() {
	options.InitProxyConfig()
	config := options.GetConfig()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	proxy.InstallRouter(r)
	log.Printf("starting server on %s:%d", config.Http.Host, config.Http.Port)
	r.Run(fmt.Sprintf("%s:%v", config.Http.Host, config.Http.Port))
}
