package proxy

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Proxy struct {
	baseUrl         string
	enableBasicAuth bool
	userName        string
	password        string
}

func NewProxy(viper *viper.Viper) *Proxy {
	return &Proxy{
		baseUrl:         viper.GetString("PROXY_BASE_URL"),
		enableBasicAuth: viper.GetBool("PROXY_BASIC_AUTH_ENABLED"),
		userName:        viper.GetString("PROXY_BASIC_AUTH_USER_NAME"),
		password:        viper.GetString("PROXY_BASIC_AUTH_PASSWORD"),
	}
}

func (p *Proxy) Handle(c *gin.Context) {
	remote, err := url.Parse(p.baseUrl)
	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.Director = func(req *http.Request) {
		req.Header = c.Request.Header

		if p.enableBasicAuth {
			req.Header.Add("Authorization", "Basic "+basicAuth(p.userName, p.password))
		}

		req.Host = remote.Host
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
		req.URL.Path = fmt.Sprintf("/transmission%s", c.Param("proxyPath"))
	}

	proxy.ServeHTTP(c.Writer, c.Request)
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
