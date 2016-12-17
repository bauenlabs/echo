// Main package initializes Echo.
package main

import (
	"echo/lib/server"
	"os"
)

func main() {
	server.Serve(os.Getenv("ECHO_SERVER_PORT"))
}
