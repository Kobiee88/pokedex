package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	LocationAreas map[string]cacheEntry
	interval      time.Duration
	mutex         *sync.RWMutex
}

type cacheEntry struct {
	val       []byte
	createdAt time.Time
}

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.LocationAreas[key] = cacheEntry{
		val:       val,
		createdAt: time.Now(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	entry, exists := c.LocationAreas[key]
	if !exists {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) Reap() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		c.mutex.Lock()
		for key, entry := range c.LocationAreas {
			if time.Since(entry.createdAt) > c.interval {
				delete(c.LocationAreas, key)
			}
		}
		c.mutex.Unlock()
	}
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		LocationAreas: make(map[string]cacheEntry),
		interval:      interval,
		mutex:         &sync.RWMutex{},
	}
	go cache.Reap()
	return cache
}
