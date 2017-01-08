// Package comissioner takes a gin request and fetches the response from the
// cache or the actual host (proxy).
package commissioner

import (
	"fmt"
	"github.com/voiceis/echo/lib/cache"
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
)

// Takes a gin request and delegates the request to the cache or proxy depending
// on the request type, and whether or not the response is in the cache.
func Spawn(c *gin.Context) {
	//fmt.Println(c.Param("param"))
	//fmt.Println(c.Request.URL.Query())

	// Initialize an empty payload.
	payload := []byte("")

	if c.Request.Method == http.MethodGet || c.Request.Method == "" {
		payload = []byte(cache.Lookup(c))

		// If the payload coming from the cache is empty, spawn a proxy, fetch the
		// desired value for the cache key, cache the value, and re-set payload.
		if len(payload) == 0 {
			fmt.Println("asdf")

		}
	} else {
		// Spawn proxy, relay value.
	}

	c.Data(http.StatusOK, "text/html", payload)
}
