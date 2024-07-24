package cache

import (
    "sync"
    "time"
    "container/list"
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
    mutex    sync.Mutex
}

func NewLRUCache(capacity int) *LRUCache {
    return &LRUCache{
        capacity: capacity,
        items:    make(map[string]*list.Element),
        order:    list.New(),
    }
}

func (c *LRUCache) Set(key string, value interface{}, expiration time.Duration) {
    c.mutex.Lock()
    defer c.mutex.Unlock()

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

func (c *LRUCache) Get(key string) (interface{}, bool) {
    c.mutex.Lock()
    defer c.mutex.Unlock()

    if elem, found := c.items[key]; found {
        if time.Now().UnixNano() > elem.Value.(*CacheItem).Expiration {
            c.order.Remove(elem)
            delete(c.items, key)
            return nil, false
        }
        c.order.MoveToFront(elem)
        return elem.Value.(*CacheItem).Value, true
    }
    return nil, false
}

func (c *LRUCache) Delete(key string) {
    c.mutex.Lock()
    defer c.mutex.Unlock()

    if elem, found := c.items[key]; found {
        c.order.Remove(elem)
        delete(c.items, key)
    }
}

func (c *LRUCache) evict() {
    elem := c.order.Back()
    if elem != nil {
        c.order.Remove(elem)
        delete(c.items, elem.Value.(*CacheItem).Key)
    }
}
