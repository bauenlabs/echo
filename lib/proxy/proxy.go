// Package proxy contains a Spawn method that takes a gin request and
// fetches the response from actual host (proxy).
package proxy

import (
	"bytes"
	"github.com/voiceis/echo/lib/concat"
	"github.com/voiceis/echo/lib/host"
	"github.com/voiceis/echo/lib/log"
	"gopkg.in/gin-gonic/gin.v1"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
)

// Custom transport struct for the proxy.
type transport struct {
	http.RoundTripper
}

// Performs the round trip request, and optionally caches response.
func (t *transport) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	resp, err = t.RoundTripper.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = resp.Body.Close()
	if err != nil {
		return nil, err
	}

	b = bytes.Replace(b, []byte("server"), []byte("schmerver"), -1)
	body := ioutil.NopCloser(bytes.NewReader(b))
	resp.Body = body
	resp.ContentLength = int64(len(b))
	resp.Header.Set("Content-Length", strconv.Itoa(len(b)))
	return resp, nil
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
		c.Next()
		return
	}

	c.Request.Header.Set("Accept-Encoding", "")
	c.Request.Host = originUrl.Host

	// Form a proxy to the origin url.
	proxy := httputil.NewSingleHostReverseProxy(originUrl)

	// Replace the proxy transport with Echo's custom transport.
	//proxy.Transport = &transport{http.DefaultTransport}

	// Write the proxy's response to the request response writer.
	log.Info("Delegating request to proxy")
	proxy.ServeHTTP(c.Writer, c.Request)
}

// Gin middleware wrapper for Spawn method.
func Middleware() gin.HandlerFunc {
	return Spawn
}
