// Package server sets up a cache server that handles requests and delegates
// them to a package router.
package server

import (
	"echo/lib/concat"
	"fmt"
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
)

// Sets up an http server that handles all requests.
func Serve(port string) {
	// Create a new gin router.
	router := gin.Default()

	// Respond to /* requests.
	router.GET("/*param", handleGET)
	router.POST("/*param", commission)
	router.PUT("/*param", commission)
	router.DELETE("/*param", commission)
	router.PATCH("/*param", commission)
	router.OPTIONS("/*param", commission)

	// Start the server on the specified port.
	router.Run(concat.Concat(":", port))
}

// handleGET handles POST requests and passes them off to the router.
func handleGET(c *gin.Context) {
	fmt.Println(c.Param("param"))
	fmt.Println(c.Request.URL.Query())
	c.String(http.StatusOK, "Hello world!")
}

func commission(c *gin.Context) {

}
