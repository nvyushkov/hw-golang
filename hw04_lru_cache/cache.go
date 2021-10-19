package hw04lrucache

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
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	_, hasContain := l.items[key]

	if hasContain {
		l.queue.MoveToFront(l.items[key])
	} else {
		l.items[key] = l.queue.PushFront(key)
	}

	if l.queue.Len() > l.capacity {
		lastItem := l.queue.Back()
		l.queue.Remove(lastItem)
		for k, v := range l.items {
			if v.Value != lastItem.Value {
				continue
			}
			delete(l.items, k)
			break
		}
	}

	l.items[key].Value = value

	return hasContain
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	_, hasContain := l.items[key]

	if !hasContain {
		return nil, hasContain
	}

	l.queue.MoveToFront(l.items[key])

	return l.items[key].Value, hasContain
}

func (l *lruCache) Clear() {
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
