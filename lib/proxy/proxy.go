// Package proxy contains a Spawn method that takes a gin request and
// fetches the response from actual host (proxy).
package proxy

import (
	"github.com/voiceis/echo/lib/cache"
	"github.com/voiceis/echo/lib/concat"
	"github.com/voiceis/echo/lib/host"
	"github.com/voiceis/echo/lib/log"
	"gopkg.in/gin-gonic/gin.v1"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

var EchoMode string = "release"

// Initialize this package.
func init() {
	mode := os.Getenv("ECHO_MODE")

	// If a mode environment variable is specified, override default.
	if len(mode) > 0 {
		EchoMode = mode
	}
}

// Custom transport struct for the proxy.
type transport struct {
	http.RoundTripper
	cacheHost string
}

// Performs the round trip request, and optionally caches response.
func (t *transport) RoundTrip(request *http.Request) (response *http.Response, err error) {
	// Spawn the round trip. If an error is return, return no response and an err.
	response, err = t.RoundTripper.RoundTrip(request)

	// If this is not supposed to be cached, return right away.
	if EchoMode != "test" || !cache.ShouldBeCached(request) || response.StatusCode != http.StatusOK {
		return response, err
	}

	// Spawn a routine to insert this item into the cache.
	go func() {
		// Read the response body, and exit if there's an error.
		body, err := ioutil.ReadAll(response.Body)

		if err != nil {
			log.Error(err)
		}

		// Create a cache object for this request, and return.
		log.Info("Inserting item into cache.")
		request.Host = t.cacheHost
		cache.Create(request, string(body))
	}()

	return response, err
}

// Fetches the origin url that the proxy should be passed to.
func fetchUrl(c *gin.Context) (*url.URL, error) {
	urlString := ""

	// Fetch the origin's IP address.
	originIp, err := host.Lookup(c.Request.Host)

	// If there was no problem finding the host, construct url.
	if err == nil {
		urlString = concat.Concat(
			"http://",
			originIp,
			c.Request.URL.Path,
		)
	}

	// Parse url string into a url.URL object.
	urlObj, err := url.Parse(urlString)

	return urlObj, err
}

// Takes a gin request and fetches the request results from the proxy host.
func Spawn(c *gin.Context) {
	originUrl, err := fetchUrl(c)

	// If the origin url failed to be constructed, move to the next middleware.
	if err != nil {
		log.Error(err)
		c.Data(418, "text/html", []byte(cache.Get("proxyError")))
		return
	}

	// Form a proxy to the origin url.
	proxy := httputil.NewSingleHostReverseProxy(originUrl)

	// Replace the proxy transport with Echo's custom transport.
	proxy.Transport = &transport{http.DefaultTransport, c.Request.Host}

	c.Request.Header.Set("Accept-Encoding", "")
	c.Request.Host = originUrl.Host

	// Write the proxy's response to the request response writer.
	log.Info("Delegating request to proxy")
	proxy.ServeHTTP(c.Writer, c.Request)
}

// Gin middleware wrapper for Spawn method.
func Middleware() gin.HandlerFunc {
	return Spawn
}
