package main

import (
	"testing"

	redis_cache_mock "github.com/Jonathansoufer/go-distributed-cache-pub-sub-redis/mocks"
)

func TestNewStore(t *testing.T) {
	c := &redis_cache_mock.MockCacher{}
	s := NewStore(c)

	if s == nil {
		t.Errorf("NewStore returned nil")
	}

	// Add more assertions here
}

func TestStore_Get(t *testing.T) {
	c := &redis_cache_mock.MockCacher{}
	s := NewStore(c)

	// Add test data to the store
	s.Set(1, "one")
	s.Set(2, "two")

	// Test getting a value from the cache
	c.EXPECT().Get(1).Return("one", true)
	val, err := s.Get(1)
	if err != nil {
		t.Errorf("Get returned an error: %v", err)
	}
	if val != "one" {
		t.Errorf("Get returned the wrong value: %s", val)
	}

	// Test getting a value from the internal storage
	c.EXPECT().Get(3).Return("", false)
	val, err = s.Get(3)
	if err != nil {
		t.Errorf("Get returned an error: %v", err)
	}
	if val != "three" {
		t.Errorf("Get returned the wrong value: %s", val)
	}

}
