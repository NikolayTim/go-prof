package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
	GetItems() map[Key]*ListItem
}

type lruCache struct {
	capacity int
	queue    List
	mu       sync.RWMutex
	items    map[Key]*ListItem
}

func (cache *lruCache) GetItems() map[Key]*ListItem {
	return cache.items
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
		item := cache.queue.PushFront(value)

		cache.items[key] = item

		return true
	}

	if cache.queue.Len() == cache.capacity {
		last := cache.queue.Back()

		for index, value := range cache.items {
			if value == last {
				delete(cache.items, index)
				break
			}
		}

		cache.queue.Remove(last)
	}

	item := cache.queue.PushFront(value)
	cache.items[key] = item

	return false
}

func (cache *lruCache) Clear() {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	cache.queue = NewList()
	cache.items = make(map[Key]*ListItem, cache.capacity)
}
