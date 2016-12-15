// DOCUMENTME
package main

import (
   "gopkg.in/redis.v5"
   "fmt"
)

//preform a look up in redis
func lookup(hash string ) {
    client := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })

    value, err := client.Get(hash).Result()
    fmt.Println(value, err)
    // Output: PONG <nil>
}

func main() {
  lookup("test")
}
