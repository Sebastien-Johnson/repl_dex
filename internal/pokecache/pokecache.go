package pokecache

import (
	"sync"
	"time"
)

//cache struct holds map and mutex
//cache map holds createdAt and val

type Cache struct{
	cache map[string]cacheEntry
	mu sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	val []byte //list of http response bodies
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache : make(map[string]cacheEntry),
	}
	//clears cache on interval
	go c.ReapLoop(interval)
	return c
	// mux := &sync.RWMutex{}

	// go c.ReapLoop(interval, mux)
	// return c

}

func (c *Cache) Add(key string, val []byte) { // , mux * sync.RWMutex
	// mux.Lock()
	c.cache[key] = cacheEntry{
		val: val, 
		createdAt: time.Now().UTC(),
	}
	// mux.Unlock()
}

func (c *Cache) Get(key string,) ([]byte, bool) {  //mux * sync.RWMutex
	// mux.RLock()
	entry, ok := c.cache[key]
	// mux.RUnlock()
	return entry.val, ok
} 

func (c *Cache) ReapLoop(interval time.Duration) { //, mux * sync.RWMutex
		ticker := time.NewTicker((interval))
		for range ticker.C {
			c.reap(interval)
		}
}

//deletes all old cache entries
func (c *Cache) reap(interval time.Duration) {
	timeCutoff := time.Now().UTC().Add(-interval)
	//mux.Lock()
	for k, v := range c.cache {
		if v.createdAt.Before(timeCutoff) {
			delete(c.cache, k)
		}
	}
	//mux.Unlock()
}



