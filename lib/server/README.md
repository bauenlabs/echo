# server
--
    import "github.com/voiceis/echo/lib/server"

Package server sets up a cache server that handles requests and delegates them
to a package router.

## Usage

#### func  Serve

```go
func Serve(port string)
```
Sets up an http server that handles all requests.
