package pokecache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	mu              sync.RWMutex
	data            map[string]CacheEntry
	intervalSeconds int
}

func NewCache(intervalSeconds int) *Cache {
	c := Cache{
		data:            make(map[string]CacheEntry),
		intervalSeconds: intervalSeconds,
	}
	go c.tick()
	return &c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = CacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	entry, ok := c.data[key]
	return entry.val, ok
}

func (c *Cache) Remove(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.data, key)
}

func (c *Cache) tick() {
	duration := time.Duration(c.intervalSeconds) * time.Second
	ticker := time.NewTicker(duration)
	defer ticker.Stop()

	for range ticker.C {
		for key, entry := range c.data {
			if time.Since(entry.createdAt) > duration {
				c.Remove(key)
			}
		}
	}
}
