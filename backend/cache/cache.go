package cache

import (
	"container/list"
	"time"
)

type CacheItem struct {
	Key        string
	Value      interface{}
	Expiration int64
}

type LRUCache struct {
	capacity int
	items    map[string]*list.Element
	order    *list.List
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		items:    make(map[string]*list.Element),
		order:    list.New(),
	}
}

func (c *LRUCache) Set(key string, value interface{}, expiration time.Duration) {
	if elem, found := c.items[key]; found {
		c.order.MoveToFront(elem)
		elem.Value.(*CacheItem).Value = value
		elem.Value.(*CacheItem).Expiration = time.Now().Add(expiration).UnixNano()
		return
	}

	item := &CacheItem{
		Key:        key,
		Value:      value,
		Expiration: time.Now().Add(expiration).UnixNano(),
	}
	elem := c.order.PushFront(item)
	c.items[key] = elem

	if c.order.Len() > c.capacity {
		c.evict()
	}
}

func (c *LRUCache) Delete(key string) bool {
	if elem, found := c.items[key]; found {
		c.order.Remove(elem)
		delete(c.items, key)
		return true
	}
	return false
}

func (c *LRUCache) evict() {
	elem := c.order.Back()
	if elem != nil {
		c.order.Remove(elem)
		delete(c.items, elem.Value.(*CacheItem).Key)
	}
}

func (c *LRUCache) GetList() map[string]interface{} {
	result := make(map[string]interface{})
	for _, elem := range c.items {
		cacheItem := elem.Value.(*CacheItem)
		if time.Now().UnixNano() <= cacheItem.Expiration {
			result[cacheItem.Key] = cacheItem.Value
		}
	}
	return result
}
