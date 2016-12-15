// DOCUMENTME
package router

import (
	"gopkg.in/redis.v5"
	"os"
	"echo/lib/concat"
)

// Open connection.
var Client = redis.NewClient(&redis.Options{
  Addr: concat.Concat(os.Getenv("REDISHOST"), ":", os.Getenv("PORT")),
  Password: "",
  DB: 0,
})

//look up a key in redis and return its value
func Lookup(hash string) (string) {
    value, _ := Client.Get(hash).Result()
    return value
}

  

