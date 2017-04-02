// Package proxy contains a Spawn method that takes a gin request and
// fetches the response from actual host (proxy).
package proxy

import (
	"bytes"
	"github.com/bauenlabs/echo/lib/cache"
	"github.com/bauenlabs/echo/lib/concat"
	"github.com/bauenlabs/echo/lib/host"
	log "github.com/bauenlabs/scribe"
	"gopkg.in/gin-gonic/gin.v1"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

var EchoCache string = "true"

// Initialize this package.
func init() {
	cache := os.Getenv("ECHO_CACHE")

	// If the cache env variable has been set, overwrite the EchoCache var.
	if len(cache) > 0 {
		EchoCache = cache
	}
}

// Custom transport struct for the proxy.
type transport struct {
	http.RoundTripper
	originalRequest *http.Request
}

// Performs the round trip request, and optionally caches response.
func (t *transport) RoundTrip(request *http.Request) (response *http.Response, err error) {
	request.Header.Set("Host", t.originalRequest.Host)
	request.Header.Set("Accept-Encoding", "")
	request.Host = t.originalRequest.Host

	// Spawn the round trip. If an error is return, return no response and an err.
	response, err = t.RoundTripper.RoundTrip(request)

	// If this is not supposed to be cached, return right away.
	if EchoCache != "true" || !cache.ShouldBeCached(request) || response.StatusCode != 200 {
		return response, err
	}

	// Read response body, and then re-set it to it's original state. This must
	// be done prior to attempting to cache the response, and prior to this
	// transport's returning of the response. This is admitidly slow, but in order
	// to cache the body, the full body value must be received from the network.
	// Once the body has been read once, it must be re-set to it's initial value
	// to be readable again. This is slow, but it is nessecary, and all subsequent
	// reads of the response body will be much more quick, as it's not receiving
	// the bytes from the network, it's fetching them from memory.
	body, parseErr := ioutil.ReadAll(response.Body)
	response.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	if parseErr != nil {
		log.Error(parseErr)
	}

	// Spawn an asyncronous subroutine to insert this item into the cache.
	go func() {
		log.Info("Inserting item into cache.")
		cache.Create(t.originalRequest, string(body))
	}()

	return response, err
}

// Fetches the origin url that the proxy should be passed to.
func fetchOriginUrl(c *gin.Context) (*url.URL, error) {
	urlString := ""

	// Fetch the origin's IP address.
	originIp, err := host.Lookup(c.Request.Host)

	// If there was no problem finding the host, construct url.
	if err == nil {
		urlString = concat.Concat(
			"http://",
			originIp,
		)
	}

	// Parse url string into a url.URL object.
	urlObj, err := url.Parse(urlString)

	return urlObj, err
}

// Takes a gin request and fetches the request results from the proxy host.
func Spawn(c *gin.Context) {
	originUrl, err := fetchOriginUrl(c)

	// If the origin url failed to be constructed, move to the next middleware.
	if err != nil {
		log.Error(err)
		c.Data(418, "text/html", []byte(cache.Get("proxyError")))
		return
	}

	// Form a proxy to the origin url.
	proxy := httputil.NewSingleHostReverseProxy(originUrl)

	// Replace the proxy transport with Echo's custom transport.
	proxy.Transport = &transport{http.DefaultTransport, c.Request}

	// Write the proxy's response to the request response writer.
	log.Info("Delegating request to proxy")
	proxy.ServeHTTP(c.Writer, c.Request)
}

// Gin middleware wrapper for Spawn method.
func Middleware() gin.HandlerFunc {
	return Spawn
}
