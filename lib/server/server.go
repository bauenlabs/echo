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

	router.Run(concat.Concat(":", port))
}

// HandleRequest handles all requests and hands them off to the router package.
func HandleRequest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hey there"))
}
