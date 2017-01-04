// Package server sets up a cache server that handles requests and delegates
// them to a package router.
package server

import (
	"fmt"
	"github.com/voiceis/echo/lib/cache"
	"github.com/voiceis/echo/lib/commissioner"
	"github.com/voiceis/echo/lib/concat"
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	"os"
)

var (
	ServerPort string = "80"
	ServerMode string = "release"
)

// Configures gin from environment variables.
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
	router.GET("/*param", commissioner.Spawn)
	router.POST("/*param", commissioner.Spawn)
	router.PUT("/*param", commissioner.Spawn)
	router.DELETE("/*param", commissioner.Spawn)
	router.PATCH("/*param", commissioner.Spawn)
	router.OPTIONS("/*param", commissioner.Spawn)

	// Start the server on the specified port.
	router.Run(concat.Concat(":", ServerPort))
}

// handleGET handles POST requests and passes them off to the router.
func handleGET(c *gin.Context) {
	fmt.Println(c.Param("param"))
	fmt.Println(c.Request.URL.Query())
	payload := []byte(cache.Process(c))

	c.Data(http.StatusOK, "text/html", payload)
}
