// Main package initializes Echo.
package main

import (
	"echo/lib/server"
)

func main() {
	server.Serve("6000")
}
