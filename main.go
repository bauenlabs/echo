// Main package initializes Echo.
package main

import (
	"github.com/stackimpact/stackimpact-go"
	"github.com/voiceis/echo/lib/server"
)

// Executes the Echo server.
func main() {
	// Set up profiling.
	agent := stackimpact.NewAgent()
	agent.Start(stackimpact.Options{
		AgentKey: "ea7aa282cb7ef7eba851796d7553a1bd7652bace",
		AppName:  "Echo",
	})

	server.Serve()
}
