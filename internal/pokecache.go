// Package internal je za zdaj za cache, mogoče bom dal pa še PokeAPI sem
package internal

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entries  map[string]cacheEntry
	mu       sync.Mutex
	interval time.Duration
}

func NewCache(interval int) *Cache {
	duration := time.Duration(interval)
	newCache := Cache{
		entries:  map[string]cacheEntry{},
		mu:       sync.Mutex{},
		interval: duration,
	}
	go newCache.reapLoop(duration)
	return &newCache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	newCacheEntry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.entries[key] = newCacheEntry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, found := c.entries[key]
	if found {
		return entry.val, true
	}
	return []byte{}, false
}

func (c *Cache) reapLoop(duration time.Duration) {
	t := time.NewTicker(time.Millisecond)
	defer t.Stop()
	for range t.C {
		c.mu.Lock()
		for key, entry := range c.entries {
			if duration < time.Since(entry.createdAt) {
				delete(c.entries, key)
			}
		}
		c.mu.Unlock()
	}
}
