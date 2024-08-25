package proxy

import (
	"log"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mmungdong/chatgpt-web/cmd/proxy/config"
)

type Proxy struct {
	upstreamURL *url.URL
	upstream    *httputil.ReverseProxy
}

func NewProxy() *Proxy {
	upstreamURL, err := url.Parse(strings.TrimSuffix(config.GetConfig().Chat.BaseUrl, "/v1"))
	if err != nil {
		log.Fatal(err)
	}
	upstream := httputil.NewSingleHostReverseProxy(upstreamURL)
	return &Proxy{
		upstreamURL: upstreamURL,
		upstream:    upstream,
	}
}

func (p *Proxy) ChatProxy(ctx *gin.Context) {
	ctx.Request.Host = p.upstreamURL.Host
	p.upstream.ServeHTTP(ctx.Writer, ctx.Request)
}
