package hw04lrucache

import (
	"sync"
)

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	mu       sync.Mutex
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	key   Key
	value interface{}
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	ci := cacheItem{
		value: value,
		key:   key,
	}

	if l.items[key] != nil {
		l.items[key].Value = ci
		l.queue.MoveToFront(l.items[key])
		return true
	}

	if l.queue.Len() == l.capacity {
		delete(l.items, l.queue.Back().Value.(cacheItem).key)
		l.queue.Remove(l.queue.Back())
	}
	l.queue.PushFront(ci)
	l.items[key] = l.queue.Front()
	return false
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.items[key] != nil {
		l.queue.MoveToFront(l.items[key])
		return l.items[key].Value.(cacheItem).value, true
	}
	return nil, false
}

func (l *lruCache) Clear() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.queue = NewList()
	l.items = make(map[Key]*ListItem, l.capacity)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
