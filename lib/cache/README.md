# cache
--
    import "github.com/voiceis/echo/lib/cache"

The cache package interacts with the Redis store containing cached items

## Usage

```go
var (
	RedisPort string = "6380"
	RedisHost string = "localhost"
	Client    *redis.Client
)
```
Defualt Global Variables

#### func  Delete

```go
func Delete(hash string) int64
```
Delete a Key, return 1 for sucess and 0 for failure.

#### func  Lookup

```go
func Lookup(hash string) string
```
Look up a key in redis and return its value.

#### func  Process

```go
func Process(c *gin.Context)
```
Process request context objects, check for cache.

#### func  Set

```go
func Set(hash string, value string) string
```
Create new key/value or update existing one.
