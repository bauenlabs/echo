// Package proxy contains a Spawn method that takes a gin request and
// fetches the response from actual host (proxy).
package proxy

import (
	"fmt"
	"github.com/voiceis/echo/lib/concat"
	"github.com/voiceis/echo/lib/host"
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	"strings"
	"time"
)

var netClient = &http.Client{
	Timeout: time.Second * 10,
}

func fetchUrl(c *gin.Context) string {
	// Fetch the origin's IP address.
	originIp := host.Lookup(c.Request.Host)
	var url string = concat.Concat(
		"http://",
		originIp,
		c.Request.URL.Path,
	)

	return url
}

// Takes a gin request and fetches the request results from the proxy host.
func Spawn(c *gin.Context) (*http.Response, error) {
	originUrl := fetchUrl(c)

	// Create a request.
	request, _ := http.NewRequest(c.Request.Method, originUrl, nil)

	// Make sure the proxy request has all the correct headers.
	for k, v := range c.Request.Header {
		fmt.Println(k)
		request.Header.Set(k, strings.Join(v, ""))
	}

	// Add a host header.
	request.Header.Set("Host", c.Request.Host)

	// Perform the request and return it.
	response, err := netClient.Do(request)
	return response, err
}
