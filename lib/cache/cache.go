// The cache package interacts with the Redis store containing cached items
package cache

import (
	"github.com/spaolacci/murmur3"
	"github.com/voiceis/echo/lib/concat"
	"gopkg.in/gin-gonic/gin.v1"
	"gopkg.in/redis.v5"
	"os"
	"strconv"
)

// Defualt Global Variables
var (
	RedisPort string = "6379"
	RedisHost string = "localhost"
	Client    *redis.Client
)

// The init function sets global variables and defines a Redis client
func init() {
	port := os.Getenv("ECHO_REDIS_PORT")
	host := os.Getenv("ECHO_REDIS_HOST")

	if len(port) > 0 {
		RedisPort = port
	}

	if len(host) > 0 {
		RedisHost = host
	}

	Client = redis.NewClient(&redis.Options{
		Addr:     concat.Concat(RedisHost, ":", RedisPort),
		Password: "",
		DB:       0,
	})
}

// Look up a key in redis and return its value.
func Lookup(hash string) string {
	value, _ := Client.Get(hash).Result()
	return value
}

// Delete a Key, return 1 for sucess and 0 for failure.
func Delete(hash string) int64 {
	success, _ := Client.Del(hash).Result()
	return success
}

// Create new key/value or update existing one.
func Set(hash string, value string) string {
	status, _ := Client.Set(hash, value, 0).Result()
	return status
}

// Generates murmur3 hash of the url, passed to the func as a string
func genHash(urlString string) uint64 {
	data := []byte(urlString)
	return murmur3.Sum64(data)
}

// Process request context objects, check for cache.
func Process(c *gin.Context) string {
	var url string = concat.Concat(
		c.Request.Host,
		c.Request.URL.Path,
	)
	var hash string = strconv.Itoa(int(genHash(url)))
	payload := Lookup(hash)
	return payload
}
