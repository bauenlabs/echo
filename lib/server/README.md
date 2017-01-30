# server
--
    import "github.com/voiceis/echo/lib/server"

Package server sets up a cache server that handles requests and delegates them
to a package router.

## Usage

```go
var (
	ServerPort string = "80"
	EchoMode   string = "release"
)
```

#### func  Serve

```go
func Serve()
```
Sets up an http server that handles all requests.
