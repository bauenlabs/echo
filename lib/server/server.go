// Package server sets up a cache server that handles requests and delegates
// them to a package router.
package server

import (
	"fmt"
	"github.com/voiceis/echo/lib/concat"
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
)

// Sets up an http server that handles all requests.
func Serve(port string) {
	// Create a new gin router.
	router := gin.Default()

	// Respond to /* requests.
	router.GET("/*param", handleGET)
	router.POST("/*param", handlePOST)
	router.PUT("/*param", handlePUT)
	router.DELETE("/*param", handleDELETE)
	router.PATCH("/*param", handlePATCH)
	router.HEAD("/*param", handleHEAD)
	router.OPTIONS("/*param", handleOPTIONS)

	// Start the server on the specified port.
	router.Run(concat.Concat(":", port))
}

// handleGET handles POST requests and passes them off to the router.
func handleGET(c *gin.Context) {
	fmt.Println(c.Param("param"))
	fmt.Println(c.Request.URL.Query())
	c.String(http.StatusOK, "Hello world!")
}

// handlePOST handles POST requests and passes them off to the router.
func handlePOST(c *gin.Context) {

}

// handlePUT handles PUT requests and passes them off to the router.
func handlePUT(c *gin.Context) {

}

// handleDELETE handles DELETE requests and passes them off to the router.
func handleDELETE(c *gin.Context) {

}

// handlePATCH handles PATCH requests and passes them off to the router.
func handlePATCH(c *gin.Context) {

}

// handleHEAD handles HEAD requests and passes them off to the router.
func handleHEAD(c *gin.Context) {

}

// handleOPTIONS handles OPTIONS requests and passes them off to the router.
func handleOPTIONS(c *gin.Context) {

}
