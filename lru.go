package LRU

import (
	"container/list"
)

type LRU struct {
	capacity int
	list     *list.List
	items    map[interface{}]*list.Element
}

type Entry struct {
	key   interface{}
	value interface{}
}

func New(capacity int) *LRU {
	if capacity < 0 {
		panic("Capacity must greater than 0")
	}
	return &LRU{capacity: capacity, list: list.New(), items: make(map[interface{}]*list.Element)}
}

func (l *LRU) Set(key, value interface{}) bool {
	if elem, ok := l.items[key]; ok {
		l.list.MoveToFront(elem)
		return false
	}

	entry := &Entry{key, value}
	elem := l.list.PushBack(entry)
	l.items[key] = elem

	if l.list.Len() > l.capacity {
		l.removeOldest()
	}
	return true
}

func (l *LRU) Get(key interface{}) (interface{}, interface{}, bool) {
	if elem, ok := l.items[key]; ok {
		l.list.MoveToFront(elem)
		return key, l.items[key].Value, true
	}

	return nil, nil, false
}

//func (l *LRU) Purge(key interface{}) bool {
//
//}

func (l *LRU) removeOldest() bool {
	elem := l.list.Back()

	if elem != nil {
		elemInterface := l.list.Remove(elem)
		delete(l.items, elemInterface.(*Entry).key)
		return true
	}

	return false
}
