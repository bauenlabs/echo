// Takes a host, an IP address, and inserts it into redis.
package main

import (
	"fmt"
	"github.com/voiceis/echo/lib/host"
	"os"
)

func main() {
	url := os.Args[1]
	ip := os.Args[2]
	host.Create(url, ip)
	fmt.Println("Inserted host.")
}
