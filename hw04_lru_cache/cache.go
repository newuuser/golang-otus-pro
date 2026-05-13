package hw04lrucache

import "sync"

/*type List interface {
	Len() int                           // длина списка
	Front() *ListItem                   // первый элемент списка
	Back() *ListItem                    // последний элемент списка
	PushFront(v interface{}) *ListItem  // добавить значение в начало
	PushBack(v interface{}) *ListItem   // добавить значение в конец
	Remove(i *ListItem)                 // удалить элемент
	MoveToFront(i *ListItem)
}*/

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
	mutex    sync.Mutex
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	v, ok := c.items[key]
	if ok {
		v.Value = value
		c.queue.MoveToFront(v)
		return true
	}
	if c.queue.Len() == c.capacity {
		c.queue.Remove(c.queue.Back())
	}
	c.items[key] = c.queue.PushFront(value)
	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	v, ok := c.items[key]
	if !ok {
		return nil, false
	}
	c.queue.MoveToFront(v)
	return v.Value, true
}

func (c *lruCache) Clear() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.queue = NewList()
	clear(c.items)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
		mutex:    sync.Mutex{},
	}
}
