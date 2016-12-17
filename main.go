// Main package initializes Echo.
package main

import (
	"echo/lib/log"
	"echo/lib/server"
)

func main() {
	server.Serve("8000")
}
