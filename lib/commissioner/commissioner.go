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
	fmt.Println(c.Param("param"))
	fmt.Println(c.Request.URL.Query())
	payload := []byte(cache.Process(c))

	c.Data(http.StatusOK, "text/html", payload)

}
