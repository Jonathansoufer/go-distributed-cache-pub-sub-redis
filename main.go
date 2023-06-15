package main

import (
	"fmt"
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
	ttl := 4 * time.Second
	s := NewStore(NewRedisCache(client, ttl))
	val, err := s.Get(1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(val)
}
