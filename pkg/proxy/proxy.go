package proxy

import (
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mmungdong/chatgpt-web/cmd/proxy/options"
)

type Proxy struct {
	upstreamURL *url.URL
	upstream    *httputil.ReverseProxy
}

func NewProxy() *Proxy {
	upstreamURL, err := url.Parse(strings.TrimSuffix(options.GetConfig().Chat.BaseUrl, "/v1"))
	if err != nil {
		log.Fatal(err)
	}
	upstream := httputil.NewSingleHostReverseProxy(upstreamURL)

	// 自定义配置，以便控制代理请求的各类超时行为
	upstream.Transport = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 10 * time.Second,
		}).DialContext,
		TLSHandshakeTimeout:   10 * time.Second,
		ResponseHeaderTimeout: 10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	return &Proxy{
		upstreamURL: upstreamURL,
		upstream:    upstream,
	}
}

// 反向代理
func (p *Proxy) ChatProxy(ctx *gin.Context) {
	ctx.Request.Host = p.upstreamURL.Host
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered from panic: %v", r)
			ctx.AbortWithStatus(http.StatusInternalServerError)
		}
	}()
	p.upstream.ServeHTTP(ctx.Writer, ctx.Request)
}
