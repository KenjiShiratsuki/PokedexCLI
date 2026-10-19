package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheFrame map[string]cacheEntry
	mu         sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	nc := &Cache{
		cacheFrame: make(map[string]cacheEntry),
	}
	go nc.reapLoop(interval)
	return nc
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cacheFrame[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if entry, ok := c.cacheFrame[key]; ok {
		return entry.val, ok
	}
	return nil, false
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.mu.Lock()
		for key, entry := range c.cacheFrame {
			if time.Since(entry.createdAt) > interval {
				delete(c.cacheFrame, key)
			}
		}
		c.mu.Unlock()
	}
}
