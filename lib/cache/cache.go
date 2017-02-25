// The cache package interacts with the Redis store containing cached items
package cache

import (
	"github.com/spaolacci/murmur3"
	"github.com/voiceis/echo/lib/concat"
	"github.com/voiceis/echo/lib/log"
	"gopkg.in/gin-gonic/gin.v1"
	"gopkg.in/redis.v5"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// Default Global Variables.
var (
	RedisPort     string = "6379"
	RedisHost     string = "localhost"
	RedisPassword string = ""
	RedisDB       int    = 0
	Client        *redis.Client
)

// The init function sets global variables and defines a Redis client.
func init() {
	port := os.Getenv("ECHO_REDIS_PORT")
	host := os.Getenv("ECHO_REDIS_HOST")
	password := os.Getenv("ECHO_REDIS_PASSWORD")

	if len(port) > 0 {
		RedisPort = port
	}

	if len(host) > 0 {
		RedisHost = host
	}

	if len(password) > 0 {
		RedisPassword = password
	}

	Client = redis.NewClient(&redis.Options{
		Addr:     concat.Concat(RedisHost, ":", RedisPort),
		Password: RedisPassword,
		DB:       RedisDB,
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

// Generates murmur3 hash of the url, passed to the func as a string.
func genHash(urlString string) uint64 {
	data := []byte(urlString)
	return murmur3.Sum64(data)
}

// Takes a request object and generates a cache key from the request details.
func genCacheKey(r *http.Request) string {
	var url string = concat.Concat(
		r.Host,
		r.URL.Path,
	)

	return strconv.Itoa(int(genHash(url)))
}

// Process request context objects, check for cache.
func Lookup(r *http.Request) string {
	payload := Get(genCacheKey(r))
	return payload
}

// Takes a request object and a body, generates a cache key, and inserts into
// the cache store.
func Create(r *http.Request, body string) string {
	return Set(genCacheKey(r), body)
}

// Inspects a context object, and returns a bool indicating  whether or not a
// cache object could or should exist for the request response.
func ShouldBeCached(r *http.Request) bool {
	if (r.Method == http.MethodGet || r.Method == "") && IsCacheableContentType(r) {
		return true
	}

	return false
}

// Takes a content type string, and returns true if it can/should be cached,
// false if it should not be cached.
func IsCacheableContentType(r *http.Request) bool {
	contentType := acceptToContentTypeHeader(r)
	switch contentType {
	case
		"*/*",
		"text/html",
		"application/html",
		"text/css",
		"application/css",
		"text/javascript",
		"application/javascript",
		"text/json",
		"application/json":
		return true
	}

	return false
}

// Parses Accept header and figures out the primary accepted content type.
func acceptToContentTypeHeader(r *http.Request) string {
	// Parse out accepted content type. Selecting the first content should be
	// good enough, applications should always list the accepted content types
	// in order of acceptability.
	contentTypes := r.Header.Get("Accept")
	multipleIndex := strings.Index(contentTypes, ",")
	contentType := "text/html"

	// If there is only one content type, do not try to parse out the first one.
	if multipleIndex == -1 {
		contentType = contentTypes
	} else {
		contentType = contentTypes[:strings.Index(contentTypes, ",")]
	}

	return contentType
}

// Gin middleware for caching mechanism. Looks for cached values for the current
// request, and respond with cached values if they exist. If this is not a
// cache-able request, or there is no cached values, this middleware will just
// skip to the next middleware.
func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// First off, check for cache-ability. If this request should never be
		// cached, don't even start talking to the cache store.
		if !ShouldBeCached(c.Request) {
			return
		}

		// Perform a cache lookup, and parse out the response content type.
		payload := []byte(Lookup(c.Request))
		contentType := acceptToContentTypeHeader(c.Request)

		// If there is a payload, respond. Otherwise, move to the next middleware.
		if len(payload) > 0 {
			log.Info("Responding with Cache.")
			c.Data(http.StatusOK, contentType, payload)
			c.Abort()
			return
		}
	}
}
