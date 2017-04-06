// Package host fetches the IP addresses of hosts with a given url.
package host

import (
	"errors"
	"github.com/bauenlabs/rivet"
	"gopkg.in/redis.v5"
	"os"
)

// Default Global Variables.
var (
	RedisPort     string = "6379"
	RedisHost     string = "localhost"
	RedisPassword string = ""
	RedisDB       int    = 1
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
		Addr:     rivet.Concat(RedisHost, ":", RedisPort),
		Password: RedisPassword,
		DB:       RedisDB,
	})
}

// Fetch a host IP for a url.
func Lookup(url string) (string, error) {
	value, err := Client.Get(url).Result()

	if err != nil || len(value) <= 0 {
		err = errors.New(rivet.Concat("Host ", url, " not found."))
	}

	return value, err
}

// Create a host in the cache.
func Create(url string, ip string) string {
	status, _ := Client.Set(url, ip, 0).Result()
	return status
}
