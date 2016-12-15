// DOCUMENTME
package router

import (
   "gopkg.in/redis.v5"
   "fmt"
)

//open connection 
var Client = redis.NewClient(&redis.Options{
  Addr:     "localhost:6379",
  Password: "", // no password set
  DB:       0,  // use default DB
})

//look up a key in redis and return its value
func lookup(hash string ) {
    value, err := Client.Get(hash).Result()
    fmt.Println(value, err)
    // Output: PONG <nil>
}

func main() {
  lookup("test")
}
