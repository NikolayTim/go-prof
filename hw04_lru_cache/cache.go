package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	mu       sync.RWMutex
	items    map[Key]*ListItem
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (cache *lruCache) Get(key Key) (interface{}, bool) {
	cache.mu.RLock()
	defer cache.mu.RUnlock()

	if item, ok := cache.items[key]; ok {
		cache.queue.MoveToFront(item)

		return item.Value, true
	}

	return nil, false
}

func (cache *lruCache) Set(key Key, value interface{}) bool {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	if item, ok := cache.items[key]; ok {
		cache.queue.Remove(item)
		item := cache.queue.PushFront(value, key)

		cache.items[key] = item

		return true
	}

	if cache.queue.Len() == cache.capacity {
		last := cache.queue.Back()

		delete(cache.items, last.Key)
		cache.queue.Remove(last)
	}

	item := cache.queue.PushFront(value, key)
	cache.items[key] = item

	return false
}

func (cache *lruCache) Clear() {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	cache.queue = NewList()
	cache.items = make(map[Key]*ListItem, cache.capacity)
}
