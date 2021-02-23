package cache

import (
	"container/list"
	"errors"
)

var (
	ErrNotExists   = errors.New("Key does not exists in cache")
	ErrInvalidSize = errors.New("Cache size cannot be 0 or negative")
)

type cacheItem struct {
	key   string
	value interface{}
}

type LRUCache struct {
	size  int
	queue *list.List
	items map[string]*list.Element
}

// Get returns cached value via key.
// If key does not exists returns ErrNotExists
func (lru *LRUCache) Get(key string) (value interface{}, err error) {
	item, ok := lru.items[key]

	if !ok {
		return nil, ErrNotExists
	}

	lru.queue.MoveToFront(item)
	value = item.Value.(cacheItem).value
	return value, nil
}

// Set store key-value pair.
// If size of cache exceed max length, older key will be dropped.
func (lru *LRUCache) Set(key string, value interface{}) {
	item, ok := lru.items[key]
	if ok {
		// key exists => update value and move key to top in queue
		lru.queue.MoveToFront(item)
		item.Value = cacheItem{key, value}
		return
	}

	if lru.queue.Len() == lru.size {
		// key does not exists, we should add the new one.
		// but max size exceeded => remove less used
		itemToRemove := lru.queue.Remove(lru.queue.Back()).(cacheItem)
		delete(lru.items, itemToRemove.key)
	}

	newItem := lru.queue.PushFront(cacheItem{key, value})
	lru.items[key] = newItem
}

// Delete removes key-value pair from cache.
func (lru *LRUCache) Delete(key string) (err error) {
	item, ok := lru.items[key]
	if !ok {
		return ErrNotExists
	}

	lru.queue.Remove(item)
	delete(lru.items, key)
	return nil
}

// Keys return list of stored keys in order
// from least to most recetly used. O(N) comlexity
func (lru *LRUCache) Keys() (keys []string) {
	keys = make([]string, lru.queue.Len())
	item := lru.queue.Front()
	for i := 0; i < lru.queue.Len(); i++ {
		keys[i] = item.Value.(cacheItem).key
		item = item.Next()
	}
	return keys
}

// NewLRUCache creates new LRU instance with provided max size
func NewLRUCache(size int) (cache *LRUCache, err error) {
	if size <= 0 {
		return nil, ErrInvalidSize
	}

	values := make(map[string]*list.Element, size)
	queue := list.New()
	return &LRUCache{size, queue, values}, nil
}
