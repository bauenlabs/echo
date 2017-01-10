// Package proxy contains a Spawn method that takes a gin request and
// fetches the response from actual host (proxy).
package proxy

import (
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	"time"
)

var netClient = &http.Client{
	Timeout: time.Second * 10,
}

// Takes a gin request and fetches the request results from the proxy host.
func Spawn(c *gin.Context) *http.Response {
	var response *http.Response

	// Execute the request depending on the type of source request. This is more
	// performant than using reflection.. (I think).
	// @TODO: Benchmark switch vs reflection.
	switch c.Request.Method {
	case http.MethodGet:
		response, _ = netClient.Get("asdf")
	case http.MethodPost:
	case http.MethodDelete:
	case http.MethodPut:
	case http.MethodPatch:
	case http.MethodOptions:
	default:
		response, _ = netClient.Get("asdf")
	}

	return response
}
