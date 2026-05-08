package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu sync.Mutex
	entries map[string]CacheEntry
}

type CacheEntry struct {
	createdAt time.Time
	val []byte
}

func NewCache(interval time.Duration) *Cache {
	newCache := &Cache{
		entries: make(map[string]CacheEntry),
	}
	go newCache.reapLoop(interval)
	return newCache
}


func (cache *Cache) Add(key string, val []byte) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	entry := CacheEntry{
		createdAt: time.Now(),
		val: val,
	}		
	cache.entries[key] = entry
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	value, ok := cache.entries[key]
	if !ok {
		return nil, false
	}
	return value.val, true
}

func (cache *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		cache.mu.Lock()		
		for k, v := range cache.entries {
			if time.Since(v.createdAt) > interval {
				delete(cache.entries, k)
			}		
		}
		cache.mu.Unlock()
	}
}