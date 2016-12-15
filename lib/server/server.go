// Package server sets up a cache server that handles requests and delegates
// them to a package router.
package server

import (
	"echo/lib/concat"
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
)

// Sets up an http server that handles all requests.
func Serve(port string) {
	// Create a new gin router.
	router := gin.Default()

	// Respond to / requests.
	router.GET("/", handleRequest)
	router.POST("/", handleRequest)
	router.PUT("/", handleRequest)
	router.DELETE("/", handleRequest)
	router.PATCH("/", handleRequest)
	router.HEAD("/", handleRequest)
	router.OPTIONS("/", handleRequest)

	// Respond to /* requests.
	router.GET("/:path", handleRequest)
	router.POST("/:path", handleRequest)
	router.PUT("/:path", handleRequest)
	router.DELETE("/:path", handleRequest)
	router.PATCH("/:path", handleRequest)
	router.HEAD("/:path", handleRequest)
	router.OPTIONS("/:path", handleRequest)

	// Start the server on the specified port.
	router.Run(concat.Concat(":", port))
}

// handleRequest handles all requests and hands them off to the router package.
func handleRequest(c *gin.Context) {
	c.String(http.StatusOK, "Hello world!")
}
