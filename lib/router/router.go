// DOCUMENTME
package main

import (
   "gopkg.in/redis.v5"
   "fmt"
)

func main() {
    client := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })

    pong, err := client.Ping().Result()
    fmt.Println(pong, err)
    // Output: PONG <nil>
}
