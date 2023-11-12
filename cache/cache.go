package cache

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type Item struct {
	Value      interface{}
	Expiration int64
}

type Cache struct {
	items map[string]Item
	mu    sync.RWMutex
}

func New() *Cache {
	cache := Cache{
		items: make(map[string]Item),
		mu:    sync.RWMutex{},
	}

	go cache.Cleaner()

	return &cache
}

func (c *Cache) Get(key string) (interface{}, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	value, ok := c.items[key]
	if !ok {
		return nil, errors.New(fmt.Sprintf("key %s is not found", key))
	}
	if time.Now().UnixNano() > value.Expiration {
		return nil, errors.New(fmt.Sprintf("KEY %s is not found", key))
	}

	return value.Value, nil
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	if _, ok := c.items[key]; ok {
		delete(c.items, key)
	}
	c.mu.Unlock()
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	c.mu.Lock()
	c.items[key] = Item{
		Value:      value,
		Expiration: time.Now().Add(ttl).UnixNano(),
	}
	c.mu.Unlock()

	// go c.deleteAfterTime(key, ttl)
}

func (c *Cache) Cleaner() {
	for {
		c.mu.RLock()
		for key, ttl := range c.items {
			c.mu.RUnlock()
			// fmt.Println(time.Duration(time.Now().Unix()), ttl.Expiration)
			if time.Now().UnixNano() > ttl.Expiration {
				c.Delete(key)
			}
			c.mu.RLock()
		}
		c.mu.RUnlock()
	}
}
