// Package server sets up a cache server that handles requests and delegates
// them to a package router.
package server

import (
	"echo/lib/concat"
	"fmt"
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	"os"
)

var (
	ServerPort string = "80"
	ServerMode string = "release"
)

// Configure gin.
func init() {
	// Load port and mode env variables if they exist.
	port := os.Getenv("ECHO_SERVER_PORT")
	mode := os.Getenv("ECHO_SERVER_MODE")

	// If a port environment variable is specified, override default.
	if len(port) > 0 {
		ServerPort = port
	}

	// If a mode environment variable is specified, override default.
	if len(mode) > 0 {
		ServerMode = mode
	}

	gin.SetMode(ServerMode)
}

// Sets up an http server that handles all requests.
func Serve() {
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
	router.Run(concat.Concat(":", ServerPort))
}

// handleGET handles POST requests and passes them off to the router.
func handleGET(c *gin.Context) {
	fmt.Println(c.Param("param"))
	fmt.Println(c.Request.URL.Query())
	c.String(http.StatusOK, "Hello world!")
}

// Spans the commissioner and skips the cache entirely.
func commission(c *gin.Context) {

}
