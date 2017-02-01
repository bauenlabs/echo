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
	"strings"
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

	// Parse out accepted content type. Selecting the first content should be
	// good enough, applications should always list the accepted content types
	// in order of acceptability.
	contentTypes := c.Request.Header.Get("Accept")
	multipleIndex := strings.Index(contentTypes, ",")
	contentType := "text/html"

	// If there is only one content type, do not try to parse out the first one.
	if multipleIndex == -1 {
		contentType = contentTypes
	} else {
		contentType = contentTypes[:strings.Index(contentTypes, ",")]
	}

	// If there is a payload, respond. Otherwise, let this method return false.
	if len(payload) > 0 {
		log.Info("Responding with Cache.")
		c.Data(http.StatusOK, contentType, payload)
		return true
	}

	return false
}

// Fetches a value for the given request from the proxy origin, and responds to
// the client. Returns true if a response was sent, or false if response failed.
func respondWithProxy(c *gin.Context) bool {
	// Fetch a response object.
	response, err := proxy.Spawn(c)

	// If the response is an error, send a proxy error, and return true to the
	// spawner since the proxy handled the response with a more secific failure.
	if err != nil {
		log.Error(err)
		c.Data(500, "text/html", []byte(cache.Get("proxyError")))
		return true
	}

	// Read the response body, and grab the content type.
	body, err := ioutil.ReadAll(response.Body)
	contentType := response.Header.Get("Content-Type")

	// If the body cannot be parsed, return false. This whole thing was a bust.
	if err != nil {
		return false
	}

	// Respond with the correct status code, content type, and body.
	c.Data(response.StatusCode, contentType, []byte(string(body)))

	// If Echo is in test mode, and this request should be cached, go ahead and
	// create a cache object for this request.
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
		// Respond from cache. If that failes, fall back to proxy. If the proxy
		// fails, respond with failure.
		if !respondWithCache(c) {
			if !respondWithProxy(c) {
				respondWithFailure(c)
			}
		}
	} else {
		// If this cannot be cached, attempt to respond with a proxy. If the proxy
		// fails, respond with failure.
		if !respondWithProxy(c) {
			respondWithFailure(c)
		}
	}
}
