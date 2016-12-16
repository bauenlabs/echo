// Package router accepts a string and returns a cached value.
package router

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
