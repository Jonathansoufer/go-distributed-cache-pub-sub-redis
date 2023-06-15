package main

import (
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

type Store struct {
	data map[int]string
}

func NewStore() *Store {
	return &Store{
		data: make(map[int]string),
	}
}

func (s *Store) Set(key int, value string) {
	s.data[key] = value
}

func (s *Store) Get(key int) (string, error) {
	val, ok := s.data[key]
	if !ok {
		return "", fmt.Errorf("key not found")
	}

	return val, nil
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	s := NewStore()
	val, err := s.Get(1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(val)
	fmt.Printf("Hello, world.%v\n", client)
}