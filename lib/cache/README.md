# cache
--
    import "github.com/voiceis/echo/lib/cache"

The cache package interacts with the Redis store containing cached items

## Usage

```go
var (
	RedisPort     string = "6379"
	RedisHost     string = "localhost"
	RedisPassword string = ""
	RedisDB       int    = 0
	Client        *redis.Client
)
```
Default Global Variables.

#### func  Create

```go
func Create(c *gin.Context, body string) string
```
Takes a request object and a body, generates a cache key, and inserts into the
cache store.

#### func  Delete

```go
func Delete(hash string) int64
```
Delete a Key, return 1 for sucess and 0 for failure.

#### func  Get

```go
func Get(hash string) string
```
Look up a key in redis and return its value.

#### func  Lookup

```go
func Lookup(c *gin.Context) string
```
Process request context objects, check for cache.

#### func  Set

```go
func Set(hash string, value string) string
```
Create new key/value or update existing one.
