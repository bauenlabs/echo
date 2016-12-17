# cache
--
    import "github.com/voiceis/echo/lib/cache"

Package router accepts a string and returns a cached value.

## Usage

```go
var Client = redis.NewClient(&redis.Options{
	Addr: concat.Concat(
		os.Getenv("ECHO_REDIS_HOST"),
		":",
		os.Getenv("ECHO_REDIS_PORT")),
	Password: "",
	DB:       0,
})
```
Open connection.

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
