package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	CreatedAt time.Time
	Val       []byte
}

type Cache struct {
	Entries map[string]cacheEntry
	Mutex   sync.Mutex
	TTL     time.Duration
}

func NewCache(ttl time.Duration) *Cache {
	cache := &Cache{
		Entries: make(map[string]cacheEntry),
		TTL:     ttl,
	}
	go cache.reapLoop()
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	c.Entries[key] = cacheEntry{
		CreatedAt: time.Now(),
		Val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	entry, exists := c.Entries[key]
	if !exists || time.Since(entry.CreatedAt) > c.TTL {
		return nil, false
	}
	return entry.Val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.TTL)
	defer ticker.Stop()

	for range ticker.C {
		c.Mutex.Lock()
		for key, entry := range c.Entries {
			if time.Since(entry.CreatedAt) > c.TTL {
				delete(c.Entries, key)
			}
		}
		c.Mutex.Unlock()
	}
}
