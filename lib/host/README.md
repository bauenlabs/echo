# host
--
    import "github.com/bauenlabs/echo/lib/host"

Package host fetches the IP addresses of hosts with a given url.

## Usage

```go
var (
	RedisPort     string = "6379"
	RedisHost     string = "localhost"
	RedisPassword string = ""
	RedisDB       int    = 1
	Client        *redis.Client
)
```
Default Global Variables.

#### func  Create

```go
func Create(url string, ip string) string
```
Create a host in the cache.

#### func  Lookup

```go
func Lookup(url string) (string, error)
```
Fetch a host IP for a url.
