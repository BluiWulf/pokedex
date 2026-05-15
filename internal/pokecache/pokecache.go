package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt 	time.Time
	val			[]byte
}

type Cache struct {
	cache 		map[string]cacheEntry
	mux			*sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	pokecache := Cache{
		cache: map[string]cacheEntry{},
		mux: &sync.Mutex{},
	}

	go pokecache.reapLoop(interval)

	return &pokecache
}

func (cache *Cache) Add(key string, val []byte) {
	cache.mux.Lock()
	cache.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
	cache.mux.Unlock()
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mux.Lock()
	entry, exists := cache.cache[key]
	cache.mux.Unlock()
	if exists {
		return entry.val, true
	}
	return nil, false
}

func (cache *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		for key, value := range cache.cache {
			elapsed := time.Since(value.createdAt)
			if elapsed > interval {
				delete(cache.cache, key)
			}
		}
	}
}

