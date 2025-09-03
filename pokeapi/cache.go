package pokeapi

import (
	"sync"
	"time"
)

type cacheEntry struct {
	value     []byte
	createdAt time.Time
}

type cache struct {
	entries map[string]cacheEntry
	mux     *sync.Mutex
}

func newCache() *cache {
	c := cache{
		entries: make(map[string]cacheEntry),
		mux:     &sync.Mutex{},
	}
	go c.reapLoop(5 * time.Minute)
	return &c
}

func newCacheEntry(value []byte) cacheEntry {
	return cacheEntry{
		value:     value,
		createdAt: time.Now(),
	}
}

func (c *cache) add(name string, value []byte) {
	entry := newCacheEntry(value)
	defer c.mux.Lock()
	c.entries[name] = entry
}

func (c *cache) delete(name string) {
	defer c.mux.Lock()
	delete(c.entries, name)
}

func (c cache) get(name string) ([]byte, bool) {
	entry, ok := c.entries[name]
	return entry.value, ok
}

func (c *cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for tick := range ticker.C {
		c.reap(interval, tick)
	}
}

func (c *cache) reap(interval time.Duration, tick time.Time) {
	for name, entry := range c.entries {
		if tick.Sub(entry.createdAt) > interval {
			c.delete(name)
		}
	}
}
