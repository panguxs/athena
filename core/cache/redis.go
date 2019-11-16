package cache

import (
	"fmt"

	"github.com/go-redis/redis/v7"
)

//ExampleNewClient example redis client for test
func ExampleNewClient() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
}
