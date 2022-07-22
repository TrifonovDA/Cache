package main

import (
	"Cache/storage"
	"errors"
	"sync"
)

var ErrNotFound = errors.New("The value not found.")

type SimpleCache struct {
	storage map[string]string
	mu      sync.RWMutex
}

func NewCache() storage.Cache {
	return &SimpleCache{
		storage: make(map[string]string),
	}
}
func (c *SimpleCache) Set(key, value string) error {
	c.mu.Lock()
	c.storage[key] = value
	c.mu.Unlock()
	return nil
}

func (c *SimpleCache) Get(key string) (string, error) {
	c.mu.RLock()
	value, ok := c.storage[key]
	if !ok {
		return "", ErrNotFound
	}
	c.mu.Unlock()
	return value, nil
}

func (c *SimpleCache) Delete(key string) error {
	c.mu.Lock()
	delete(c.storage, key)
	c.mu.Unlock()
	return nil
}
func main() {

}
