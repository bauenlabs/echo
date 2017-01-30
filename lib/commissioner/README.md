# commissioner
--
    import "github.com/voiceis/echo/lib/commissioner"

Package comissioner contains a Spawn method that takes a gin request and fetches
the response from the cache or the actual host (proxy).

## Usage

```go
var (
	EchoMode string = "release"
)
```

#### func  Spawn

```go
func Spawn(c *gin.Context)
```
Takes a gin request and delegates the request to the cache or proxy depending on
the request type, and whether or not the response is in the cache.
