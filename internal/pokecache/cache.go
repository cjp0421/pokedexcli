package pokecache

import (
	"fmt"
	"time"
)

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		cacheEntries: make(map[string]cacheEntry),
	}
	go cache.reapLoop(interval)
	return cache
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		func() {
			fmt.Println("Cleaning cache")
			c.mu.Lock()
			defer c.mu.Unlock()
			for key, entry := range c.cacheEntries {
				now := time.Now()
				if now.After(entry.createdAt.Add(interval)) {
					delete(c.cacheEntries, key)
				}
			}
		}()
	}

}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	newEntry := cacheEntry{
		createdAt: time.Now(),
		cachedVal: val,
	}
	c.cacheEntries[key] = newEntry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if value, ok := c.cacheEntries[key]; ok {
		return value.cachedVal, true
	}
	return []byte{}, false
}
