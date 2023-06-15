package main

import "fmt"

type Store struct {
	data  map[int]string
	cache Cacher
}

func NewStore(c Cacher) *Store {
	data := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}

	return &Store{
		data:  data,
		cache: c,
	}
}

func (s *Store) Set(key int, value string) {
	s.data[key] = value
}

func (s *Store) Get(key int) (string, error) {
	val, ok := s.cache.Get(key)
	if ok {
		if err := s.cache.Remove(key); err != nil {
			fmt.Println(err)
		}
		return val, nil
	}

	val, ok = s.data[key]
	if !ok {
		return "", fmt.Errorf("key not found: %d", key)
	}

	if err := s.cache.Set(key, val); err != nil {
		return "", err
	}

	fmt.Println("Returning key from internal storage")

	return val, nil
}
