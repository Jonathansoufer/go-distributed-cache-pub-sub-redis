package main

import (
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	ctx := client.Context()
	for {
		if err := client.Publish(ctx, "chan1", "Hello World").Err(); err != nil {
			log.Fatal(err)
		}
		time.Sleep(1 * time.Second)
	}
}
