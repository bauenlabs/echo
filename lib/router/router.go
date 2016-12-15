// DOCUMENTME
package main

import (
	"gopkg.in/redis.v5"
	"fmt"
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
func lookup(hash string) (string) {
    value, _ := Client.Get(hash).Result()
    return value
}

func main() {
  taco := lookup("test")
  fmt.Println(taco)
}
  

