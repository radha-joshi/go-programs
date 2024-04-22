// Package main demonstrates the usage of a singleton cache within a concurrent application.
package main

import (
	"fmt"
	"singleton/cache"
	"sync"

	"github.com/rs/zerolog/log"
)

// Application represents the core structure for the application demonstrating cache usage.
type Application struct{}

// Start initializes and runs multiple goroutines that access and modify a shared cache.
// It creates 5 goroutines, each manipulating cache entries and logging results.
func (a *Application) Start() {
	waitGroup := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		waitGroup.Add(1)
		go func(i int) {
			defer waitGroup.Done()
			process(i)
		}(i)
	}
	waitGroup.Wait() // waits for all goroutines to complete their tasks
}

// process simulates a task that modifies and queries the cache.
// It takes an integer index to demonstrate separate processing contexts.
// Each process sets values in the cache and retrieves them to log the outputs.
func process(index int) {
	cache := cache.GetInstance[string]() // retrieving the singleton instance of the cache
	cache.Set("one", fmt.Sprintf("one:%d", index))
	cache.Set("two", fmt.Sprintf("two:%d", index))

	value, exist := cache.Get("one")
	if exist {
		log.Info().Msg(fmt.Sprintf("Process: %d, Value: %s", index, value))
	}

	value2, exist := cache.Get("two")
	if exist {
		log.Info().Msg(fmt.Sprintf("Process: %d, Value: %s", index, value2))
	}
}
