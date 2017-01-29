// Package proxy contains a Spawn method that takes a gin request and
// fetches the response from actual host (proxy).
package proxy

import (
	"github.com/voiceis/echo/lib/concat"
	"github.com/voiceis/echo/lib/host"
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	"time"
)

var netClient = &http.Client{
	Timeout: time.Second * 10,
}

func fetchUrl(c *gin.Context) (string, error) {
	url := ""

	// Fetch the origin's IP address.
	originIp, err := host.Lookup(c.Request.Host)

	// If there was no problem finding the host, construct url.
	if err == nil {
		url = concat.Concat(
			"http://",
			originIp,
			c.Request.URL.Path,
		)
	}

	return url, err
}

// Takes a gin request and fetches the request results from the proxy host.
func Spawn(c *gin.Context) (*http.Response, error) {
	originUrl, err := fetchUrl(c)

	// If there was an issue creating the originUrl, exit.
	if err != nil {
		return nil, err
	}

	// Create a request.
	request, _ := http.NewRequest(c.Request.Method, originUrl, nil)

	// Set up request object.
	request.Host = c.Request.Host
	request.Header = c.Request.Header
	request.Proto = c.Request.Proto
	request.ProtoMajor = c.Request.ProtoMajor
	request.ProtoMinor = c.Request.ProtoMinor

	// Add a host header.
	request.Header.Set("Host", c.Request.Host)

	// Perform the request and return it.
	response, err := netClient.Do(request)
	return response, err
}
