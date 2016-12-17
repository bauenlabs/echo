// Package router accepts a string and returns a cached value.
package cache

import (
	"echo/lib/concat"
	"gopkg.in/redis.v5"
	"os"
)

// Open connection.
var Client = redis.NewClient(&redis.Options{
	Addr: concat.Concat(
		os.Getenv("ECHO_REDIS_HOST"),
		":",
		os.Getenv("ECHO_REDIS_PORT")),
	Password: "",
	DB:       0,
})

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
