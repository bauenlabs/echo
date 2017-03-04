# proxy
--
    import "github.com/voiceis/echo/lib/proxy"

Package proxy contains a Spawn method that takes a gin request and fetches the
response from actual host (proxy).

## Usage

```go
var EchoCache string = "true"
```

#### func  Middleware

```go
func Middleware() gin.HandlerFunc
```
Gin middleware wrapper for Spawn method.

#### func  Spawn

```go
func Spawn(c *gin.Context)
```
Takes a gin request and fetches the request results from the proxy host.
