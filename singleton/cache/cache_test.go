// Package cache provides a testing suite for its singleton cache implementation.
package cache

import "testing"

// TestGetInstance tests the singleton property of the cache.
// It verifies that multiple calls to GetInstance return the same cache instance.
func TestGetInstance(t *testing.T) {
	cache1 := GetInstance[string]()
	cache2 := GetInstance[string]()

	if cache1 != cache2 {
		t.Fatal("Should return same instance.")
	}
}

// TestGetSet tests the basic functionality of setting and retrieving a value from the cache.
// It verifies that after setting a value, it can be retrieved correctly.
func TestGetSet(t *testing.T) {
	cache := GetInstance[string]()
	cache.Set("one", "one")
	value, exist := cache.Get("one")
	if !exist {
		t.Fatal("Value should exist.")
	}
	if value != "one" {
		t.Fatal("Value should be 'one'.")
	}
}

// TestGetNoValue verifies the behavior of the cache when retrieving a non-existent key.
// It ensures that the cache correctly reports that a value does not exist for a given key.
func TestGetNoValue(t *testing.T) {
	cache := GetInstance[string]()
	if _, exist := cache.Get("two"); exist {
		t.Fatal("Value should not exist.")
	}
}
