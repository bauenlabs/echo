// The cache package interacts with the Redis store containing cached items
package cache

import (
	"github.com/spaolacci/murmur3"
	"github.com/voiceis/echo/lib/concat"
	"github.com/voiceis/echo/lib/log"
	"gopkg.in/gin-gonic/gin.v1"
	"gopkg.in/redis.v5"
	"io/ioutil"
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
func Get(hash string) string {
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

// Takes a request object and generates a cache key from the request details.
func genCacheKey(c *gin.Context) string {
	var url string = concat.Concat(
		c.Request.Host,
		c.Request.URL.Path,
	)

	return strconv.Itoa(int(genHash(url)))
}

// Process request context objects, check for cache.
func Lookup(c *gin.Context) string {
	payload := Get(genCacheKey(c))
	return payload
}

// Takes a request object and generates a cache value for the specified key.
func Create(c *gin.Context) string {
	body, err := ioutil.ReadAll(c.Request.Body)

	// If there's an error with parsing the request body, fail.
	if err != nil {
		log.Fatal(err, 1)
	}

	return Set(genCacheKey(c), string(body))
}
