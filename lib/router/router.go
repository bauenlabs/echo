// DOCUMENTME
package router

import (
   "gopkg.in/redis.v5"
)

//open connection 
var Client = redis.NewClient(&redis.Options{
  Addr:     "localhost:6379",
  Password: "", // no password set
  DB:       0,  // use default DB
})

//look up a key in redis and return its value
func lookup(hash string) (string) {
    value, _ := Client.Get(hash).Result()
    return value
}

