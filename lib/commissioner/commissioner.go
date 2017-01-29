// Package comissioner contains a Spawn method that takes a gin request and
// fetches the response from the cache or the actual host (proxy).
package commissioner

import (
	"github.com/voiceis/echo/lib/cache"
	"github.com/voiceis/echo/lib/log"
	"github.com/voiceis/echo/lib/proxy"
	"gopkg.in/gin-gonic/gin.v1"
	"io/ioutil"
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
	contentType := c.Request.Header.Get("Content-Type")

	if len(payload) > 0 {
		c.Data(http.StatusOK, contentType, payload)
		return true
	}

	log.Info("Responding with Cache.")

	return false
}

// Fetches a value for the given request from the proxy origin, and responds to
// the client. Returns true if a response was sent, or false if response failed.
func respondWithProxy(c *gin.Context) bool {
	// Fetch a response object.
	response, err := proxy.Spawn(c)

	// If the response is an error, send that up the chain.
	if err != nil {
		log.Error(err)
		c.Data(500, "text/html", []byte(cache.Get("proxyError")))
		return true
	}

	body, _ := ioutil.ReadAll(response.Body)
	contentType := response.Header.Get("Content-Type")

	// Respond to the client.
	c.Data(response.StatusCode, contentType, []byte(string(body)))

	// If Echo is in release mode, cache
	if EchoMode == "test" && canBeCached(c) && response.StatusCode == http.StatusOK {
		log.Info("Inserting item into cache.")
		cache.Create(c, string(body))
	}

	log.Info("Responding with Proxy")

	return true
}

// Responds to a request with an Echo failure, which occurs only if all other
// attempts to fetch a response fail.
func respondWithFailure(c *gin.Context) {
	c.Data(404, "text/html", []byte(cache.Get("generalError")))
}

// Inspects a context object, and returns a bool indicating  whether or not a
// cache object could or should exist for the request response.
func canBeCached(c *gin.Context) bool {
	if c.Request.Method == http.MethodGet || c.Request.Method == "" {
		return true
	}

	return false
}

// Takes a gin request and delegates the request to the cache or proxy depending
// on the request type, and whether or not the response is in the cache.
func Spawn(c *gin.Context) {
	// Only check the cache if this request should be cached.
	if canBeCached(c) {
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
