// Package server sets up a cache server that handles requests and delegates
// them to a package router.
package server

import (
	"github.com/bauenlabs/echo/lib/cache"
	"github.com/bauenlabs/echo/lib/concat"
	"github.com/bauenlabs/echo/lib/proxy"
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
		ServerPort = port
	}

	// If a mode environment variable is specified, override default.
	if len(mode) > 0 {
		EchoMode = mode
	}

	// If the app is being profiled, always run gin in release mode.
	if EchoMode == "profile" {
		EchoMode = "release"
	}

	gin.SetMode(EchoMode)

	// Make sure error pages are cached and ready to be served.
	cache.Set("proxyError", "<!DOCTYPE html><html dir=ltr lang=en-US><head><title>:( Host not found</title><meta charset=utf-8><meta content=\"noindex, nofollow\"name=robots></head><body><header><h1 id=title>Unable to find host</h1></header><section id=descriptiong><p>This host was not found within Echo, so Echo cannot find a cached response. If you are the owner of this site, please contact an Echo administrator.</p><p>Echo's owners and mantainers appologize for this inconvenience :( We'll be looking into it shortly, and using this as an opportunity to improve our system.</p></section></body></html>")
	cache.Set("generalError", "<!DOCTYPE html><html dir=ltr lang=en-US><head><title>Not Found</title><meta charset=utf-8><meta content=\"noindex, nofollow\"name=robots></head><body><header><h1 id=title>Unable to find host</h1></header><section id=descriptiong><p>This host was not found within Echo, so Echo cannot find a cached response. If you are the owner of this site, please contact an Echo administrator.</p><p>Echo's owners and mantainers appologize for this inconvenience :( We'll be looking into it shortly, and using this as an opportunity to improve our system.</p></section></body></html>")
}

// Sets up an http server that handles all requests.
func Serve() {
	// Create a new gin router.
	r := gin.Default()
	g := r.Group("/")

	// Respond to all requests with cache first, then proxy, then failure.
	g.GET("/*param", cache.Middleware(), proxy.Middleware())
	g.POST("/*param", cache.Middleware(), proxy.Middleware())
	g.PUT("/*param", cache.Middleware(), proxy.Middleware())
	g.DELETE("/*param", cache.Middleware(), proxy.Middleware())
	g.PATCH("/*param", cache.Middleware(), proxy.Middleware())
	g.OPTIONS("/*param", cache.Middleware(), proxy.Middleware())

	// Start the server on the specified port.
	r.Run(concat.Concat(":", ServerPort))
}
