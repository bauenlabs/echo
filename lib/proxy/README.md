# proxy
--
    import "github.com/voiceis/echo/lib/proxy"

Package proxy contains a Spawn method that takes a gin request and fetches the
response from actual host (proxy).

## Usage

#### func  Spawn

```go
func Spawn(c *gin.Context) (*http.Response, error)
```
Takes a gin request and fetches the request results from the proxy host.
