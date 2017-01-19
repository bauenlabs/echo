// Package comissioner contains a Spawn method that takes a gin request and
// fetches the response from the cache or the actual host (proxy).
package commissioner

import (
	"github.com/voiceis/echo/lib/cache"
	"github.com/voiceis/echo/lib/proxy"
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	"os"
)

var (
	EchoMode string = "release"
)

func init() {
	mode := os.Getenv("ECHO_MODE")

	// If a mode environment variable is specified, override default.
	if len(mode) > 0 {
		EchoMode = mode
	}
}

// Fetches a value for the given request from the cache, and responds to the
// client. Returns true if a response was sent, or false if response failed.
func respondWithCache(c *gin.Context) bool {
	payload := []byte(cache.Lookup(c))

	if len(payload) > 0 {
		c.Data(http.StatusOK, "text/html", payload)
		return true
	}

	return false
}

// Fetches a value for the given request from the proxy origin, and responds to
// the client. Returns true if a response was sent, or false if response failed.
func respondWithProxy(c *gin.Context) bool {
	//response := proxy.Spawn(c)

	return true
}

// Responds to a request with an Echo failure, which occurs only if all other
// attempts to fetch a response fail.
func respondWithFailure(c *gin.Context) {
	c.Data(http.StatusOK, "text/html", []byte("HAHAAAAAAAAAAA COCK SUCKER"))
}

// Takes a gin request and delegates the request to the cache or proxy depending
// on the request type, and whether or not the response is in the cache.
func Spawn(c *gin.Context) {
	// If this is a GET method, look in the cache before delegating to the proxy.
	if c.Request.Method == http.MethodGet || c.Request.Method == "" {
		// Respond from cache. If that failes, fall back to proxy.
		if !respondWithCache(c) {
			if !respondWithProxy(c) {
				respondWithFailure(c)
			}
		}
	} else {
		if !respondWithProxy(c) {
			respondWithFailure(c)
		}
	}
}
