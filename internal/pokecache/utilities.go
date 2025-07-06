package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	cachedVal []byte
}

type Cache struct {
	cacheEntries map[string]cacheEntry
	mu           sync.Mutex
}
