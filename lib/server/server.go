// Package server sets up a cache server that handles requests and delegates
// them to a package router.
package server

import (
	"github.com/voiceis/echo/lib/commissioner"
	"github.com/voiceis/echo/lib/concat"
	"gopkg.in/gin-gonic/gin.v1"
	"os"
)

var (
	ServerPort string = "80"
	EchoMode   string = "release"
)

// Configures gin from environment variables.
func init() {
	// Load port and mode env variables if they exist.
	port := os.Getenv("ECHO_SERVER_PORT")
	mode := os.Getenv("ECHO_MODE")

	// If a port environment variable is specified, override default.
	if len(port) > 0 {
		EchoMode = port
	}

	// If a mode environment variable is specified, override default.
	if len(mode) > 0 {
		EchoMode = mode
	}

	gin.SetMode(EchoMode)
}

// Sets up an http server that handles all requests.
func Serve() {
	// Create a new gin router.
	router := gin.Default()

	// Respond to all requests.
	router.GET("/*param", commissioner.Spawn)
	router.POST("/*param", commissioner.Spawn)
	router.PUT("/*param", commissioner.Spawn)
	router.DELETE("/*param", commissioner.Spawn)
	router.PATCH("/*param", commissioner.Spawn)
	router.OPTIONS("/*param", commissioner.Spawn)

	// Start the server on the specified port.
	router.Run(concat.Concat(":", ServerPort))
}
