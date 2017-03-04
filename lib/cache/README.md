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
func Create(r *http.Request, body string) string
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

#### func  IsCacheableContentType

```go
func IsCacheableContentType(r *http.Request) bool
```
Takes a content type string, and returns true if it can/should be cached, false
if it should not be cached.

#### func  Lookup

```go
func Lookup(r *http.Request) string
```
Process request context objects, check for cache.

#### func  Middleware

```go
func Middleware() gin.HandlerFunc
```
Gin middleware for caching mechanism. Looks for cached values for the current
request, and respond with cached values if they exist. If this is not a
cache-able request, or there is no cached values, this middleware will just skip
to the next middleware.

#### func  Set

```go
func Set(hash string, value string) string
```
Create new key/value or update existing one.

#### func  ShouldBeCached

```go
func ShouldBeCached(r *http.Request) bool
```
Inspects a context object, and returns a bool indicating whether or not a cache
object could or should exist for the request response.
