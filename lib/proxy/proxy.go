// Package proxy contains a Spawn method that takes a gin request and
// fetches the response from actual host (proxy).
package proxy

import (
	"fmt"
	"github.com/voiceis/echo/lib/concat"
	"github.com/voiceis/echo/lib/host"
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	"time"
)

var netClient = &http.Client{
	Timeout: time.Second * 10,
}

func fetchUrl(c *gin.Context) string {
	// Fetch the origin's IP address.
	originIp := host.Lookup(c.Request.Host)
	var url string = concat.Concat(
		originIp,
		c.Request.URL.Path,
	)

	fmt.Println(url)

	return url
}

// Takes a gin request and fetches the request results from the proxy host.
func Spawn(c *gin.Context) (*http.Response, error) {
	var response *http.Response
	var err error
	originUrl := fetchUrl(c)

	// Execute the request depending on the type of source request. This is more
	// performant than using reflection.. (I think).
	// @TODO: Benchmark switch vs reflection.
	switch c.Request.Method {
	case http.MethodGet:
		response, err = netClient.Get(originUrl)
	case http.MethodPost:
	case http.MethodDelete:
	case http.MethodPut:
	case http.MethodPatch:
	case http.MethodOptions:
	default:
		response, _ = netClient.Get(originUrl)
	}

	return response, err
}
