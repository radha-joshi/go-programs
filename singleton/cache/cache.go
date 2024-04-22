// Package cache provides a generic, concurrent-safe, singleton cache.
// It leverages sync package for mutexes to ensure safe concurrent access
// and implements a simple key-value storage.
package cache

import (
	"fmt"
	"sync"

	"github.com/rs/zerolog/log"
)

// Cache is an interface defining methods to interact with a cache.
// It supports generic type T which allows storing different data types.
// Cache is concurrent-safe and can be used by multiple goroutines.
type Cache[T any] interface {
	// Get retrieves the value associated with the given key.
	// If the key exists, the value and true are returned.
	// If the key does not exist, the zero value for type T and false are returned.
	Get(key string) (value T, exists bool)

	// Set stores the value associated with the given key in the cache.
	Set(key string, value T)
}

// cache implements the Cache interface using a generic type T.
// It provides a thread-safe mechanism to store and retrieve items.
type cache[T any] struct {
	mutex sync.RWMutex // protects access to the map
	items map[string]interface{}
}

var (
	instance Cache[any] // instance holds the singleton cache instance.
	once     sync.Once  // once ensures the singleton instance is created only once.
)

// createCacheInstance initializes a new cache instance.
// It is called once per application lifetime.
func createCacheInstance[T any]() {
	instance = &cache[any]{
		items: make(map[string]interface{}),
	}
	log.Info().Msg(fmt.Sprintf("Cache instance created for type %T", *new(T)))
}

// GetInstance returns the singleton instance of the Cache.
// It ensures that only one cache instance is created, regardless of the type,
// and provides access to this single instance.
func GetInstance[T any]() Cache[any] {
	once.Do(createCacheInstance[T])
	return instance
}

// Set adds or updates the value at the specified key.
// The operation locks for writing to ensure exclusive access.
func (c *cache[T]) Set(key string, value T) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.items[key] = value
}

// Get retrieves the value stored at the specified key.
// The operation locks for reading to allow concurrent access.
// It returns the value and a boolean indicating whether the key was found.
func (c *cache[T]) Get(key string) (value T, exists bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	rawValue, exists := c.items[key]
	if exists {
		value = rawValue.(T)
	}
	return value, exists
}
