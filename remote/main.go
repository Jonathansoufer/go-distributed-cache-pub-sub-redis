package main

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	ctx := context.Background()
	sub := client.Subscribe(ctx, "chan1")

	for {
		msg, err := sub.ReceiveMessage(ctx)
		if err != nil {
			log.Fatal(err)
		}

		log.Println(msg.Payload)
	}

}
