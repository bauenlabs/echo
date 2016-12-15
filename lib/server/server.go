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
	router := gin.Default()

	router.GET("/", handleRequest)
	router.POST("/", handleRequest)
	router.PUT("/", handleRequest)
	router.DELETE("/", handleRequest)
	router.PATCH("/", handleRequest)
	router.HEAD("/", handleRequest)
	router.OPTIONS("/", handleRequest)

	router.Run(concat.Concat(":", port))
}

// handleRequest handles all requests and hands them off to the router package.
func handleRequest(c *gin.Context) {
	c.String(http.StatusOK, "Hello world!")
}
