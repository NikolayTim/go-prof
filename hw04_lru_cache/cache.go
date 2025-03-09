package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	Cache // Remove me after realization.

	capacity int
	queue    List
	items    map[Key]*ListItem
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (cache lruCache) Get(key Key) (interface{}, bool) {
	if listItem, ok := cache.items[key]; ok {
		cache.queue.MoveToFront(listItem)

		return listItem.Value, true
	}

	return nil, false
}

func (cache lruCache) Set(key Key, value interface{}) bool {
	if listItem, ok := cache.items[key]; ok {
		listItem.Value = value

		return true
	}

	return false
}

func (cache lruCache) Clear() {
	cache.queue = NewList()
	cache.items = make(map[Key]*ListItem, cache.capacity)
}
