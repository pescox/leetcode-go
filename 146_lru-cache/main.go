package main

import (
	"log"
)

// 运用你所掌握的数据结构，设计和实现一个 LRU 缓存机制 。
// LRUCache(int capacity) 以正整数作为容量 capacity 初始化 LRU 缓存
// int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
// void put(int key, int value) 如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字-值」。
// 当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。

// 要求：在 O(1) 时间复杂度内完成这两种操作

// 解题思路： 利用map和双链表
func main() {
	cache := Constructor(2)
	cache.Put(1, 1)
	log.Printf("cache1: %#+v\n", cache)
	cache.Put(2, 2)
	log.Printf("cache2: %#+v\n", cache)
	cache.Get(1)
	log.Printf("cache3: %#+v\n", cache)
	cache.Put(3, 3)
	log.Printf("cache4: %#+v\n", cache)
}

type node struct {
	key   int
	value int
	prev  *node
	next  *node
}

func newNode(key, value int) *node {
	return &node{
		key:   key,
		value: value,
	}
}

type doubleList struct {
	head *node
	tail *node
	len  int
}

func newDoubleList() *doubleList {
	return &doubleList{}
}

func (l *doubleList) add(n *node) {
	if l.tail == nil {
		l.head = n
		l.tail = n
		l.len++
		return
	}
	l.tail.next = n
	n.prev = l.tail
	l.tail = n
	l.len++
}

func (l *doubleList) deleteFirst() {
	if l.head == nil {
		return
	}
	if l.head.next == nil {
		l.head = nil
		l.tail = nil
		l.len--
		return
	}

	l.head = l.head.next
	l.head.prev = nil
	l.len--
}

func (l *doubleList) delete(n *node) {
	if n.prev == nil {
		l.deleteFirst()
		return
	}
	if n.next == nil {
		l.tail = n.prev
	} else {
		n.next.prev = n.prev
	}

	n.prev.next = n.next
	n.prev = nil
	n.next = nil
	l.len--
}

type LRUCache struct {
	capacity int
	m        map[int]*node
	l        *doubleList
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		m:        make(map[int]*node, capacity),
		l:        newDoubleList(),
	}
}

func (this *LRUCache) Get(key int) int {
	n, ok := this.m[key]
	if !ok {
		return -1
	}

	this.l.delete(n)
	this.l.add(n)
	return n.value
}

func (this *LRUCache) Put(key int, value int) {
	new := newNode(key, value)
	if old, ok := this.m[key]; ok {
		this.l.delete(old)
		this.l.add(new)
		this.m[key] = new
		return
	}
	if this.capacity == len(this.m) {
		delete(this.m, this.l.head.key)
		this.l.deleteFirst()
	}
	this.m[key] = new
	this.l.add(new)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
