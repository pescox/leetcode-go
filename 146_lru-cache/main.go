package main

import (
	"container/list"
)

func main() {}

type item struct {
	key   int
	value int
}

type LRUCache struct {
	capacity int
	list     *list.List
	m        map[int]*list.Element
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		list:     list.New(),
		m:        make(map[int]*list.Element),
	}
}

func (this *LRUCache) Get(key int) int {
	if ele, ok := this.m[key]; ok {
		this.list.MoveToFront(ele)
		return ele.Value.(item).value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if ele, ok := this.m[key]; ok {
		this.list.MoveToFront(ele)
		return
	}

	if len(this.m) == this.capacity {
		delete(this.m, this.list.Back().Value.(item).key)
		this.list.Remove(this.list.Back())
	}

	ele := this.list.PushFront(item{key: key, value: value})
	this.m[key] = ele
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
