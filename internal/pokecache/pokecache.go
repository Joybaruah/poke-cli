package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mux   *sync.Mutex
}

type cacheEntry struct {
	val        []byte
	created_at time.Time
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
		mux:   &sync.Mutex{},
	}
	go c.realLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.cache[key] = cacheEntry{
		val:        val,
		created_at: time.Now(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	cacheE, ok := c.cache[key]
	return cacheE.val, ok
}

func (c *Cache) realLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(interval)
	}
}

func (c *Cache) reap(interval time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	timeAgo := time.Now().UTC().Add(-interval)
	for k, v := range c.cache {
		if v.created_at.Before(timeAgo) {
			delete(c.cache, k)
		}
	}
}
