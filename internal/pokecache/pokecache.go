package pokecache

import (
	"sync"
	"time"
)

//cache struct holds map and mutex
//cache map holds createdAt and val

type Cache struct{
	cache map[string]cacheEntry
	mux *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val []byte //list of http response bodies
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache : make(map[string]cacheEntry),
		mux : &sync.Mutex{},
	}
	go c.ReapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) { // , mux * sync.RWMutex
	c.mux.Lock()
	defer c.mux.Unlock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val: val, 
	}
}

func (c *Cache) Get(key string,) ([]byte, bool) {  //mux * sync.RWMutex
	c.mux.Lock()
	defer c.mux.Unlock()
	entry, ok := c.cache[key]
	return entry.val, ok
} 

//always running, checking if time to reap
func (c *Cache) ReapLoop(interval time.Duration) { //, mux * sync.RWMutex
		ticker := time.NewTicker(interval)
		for range ticker.C {
			c.reap(time.Now().UTC(), interval)
		}
}

//deletes all old cache entries
func (c *Cache) reap(now time.Time, timeCutoff time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	for k, v := range c.cache {
		if v.createdAt.Before(now.Add(-timeCutoff)) {
			delete(c.cache, k)
		}
	}
}



