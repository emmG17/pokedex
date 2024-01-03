package pokecache

import (
	"sync"
	"time"
  "errors"
)

type cacheEntry struct {
  val []byte
  createdAt time.Time
}

type Cache struct {
  cache map[string]cacheEntry
  mutex *sync.RWMutex
}

func (c *Cache) Add(key string, val []byte) error {
  c.mutex.Lock()
  c.cache[key] = cacheEntry{val: val, createdAt: time.Now()}
  c.mutex.Unlock()
  return nil
}

func (c *Cache) Get(key string) ([]byte, error) {
  c.mutex.RLock()
  entry, ok := c.cache[key]
  c.mutex.RUnlock()
  if !ok {
    return nil, errors.New("Key not found in cache")
  }
  return entry.val, nil
}

func (c *Cache) delete(key string) error {
  c.mutex.Lock()
  delete(c.cache, key)
  c.mutex.Unlock()
  return nil
}

func (c *Cache) reapLoop(interval time.Duration) {
  ticker := time.NewTicker(interval)
  defer ticker.Stop()
  for {
    select {
    case t := <- ticker.C:
      for _, key := range c.expiredKeys(t, interval) {
        c.mutex.Lock()
        delete(c.cache, key)
        c.mutex.Unlock()
      }
    }
  }
}

func (c *Cache) expiredKeys (currentTime time.Time, interval time.Duration) []string {
  keys := make([]string, 0)
  c.mutex.RLock()
  for key, entry := range c.cache {
    if currentTime.Sub(entry.createdAt) > interval {
      keys = append(keys, key)
    }
  }
  c.mutex.RUnlock()
  return keys
}

func NewCache(interval int) Cache {
  c := Cache{
    cache: make(map[string]cacheEntry),
    mutex: &sync.RWMutex{},
  }
  go c.reapLoop(time.Second * time.Duration(interval))
  return c
}
