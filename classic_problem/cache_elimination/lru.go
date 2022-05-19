package cache_elimination

import (
	"time"
)

type Node struct {
	key              string
	value            interface{}
	lastModifiedTime time.Time
	expiration       time.Duration
	prev, next       *Node
}

type LruCache struct {
	capacity   int
	size       int
	cache      map[string]*Node
	head, tail *Node
}

func NewLruCache(capacity int) *LruCache {
	lc := &LruCache{
		capacity: capacity,
		size:     0,
		cache:    make(map[string]*Node),
		head:     &Node{},
		tail:     &Node{},
	}
	lc.head.next = lc.tail
	lc.tail.prev = lc.head
	return lc
}

func (lc *LruCache) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (lc *LruCache) addToHead(node *Node) {
	node.prev = lc.head
	node.next = lc.head.next
	lc.head.next.prev = node
	lc.head.next = node
}

func (lc *LruCache) moveToHead(node *Node) {
	lc.removeNode(node)
	lc.addToHead(node)
}

func (lc *LruCache) Get(key string) (bool, interface{}) {
	node, ok := lc.cache[key]
	if !ok {
		return false, nil
	}
	lc.moveToHead(node)
	return true, node.value
}

func (lc *LruCache) removeTail() *Node {
	tailNode := lc.tail.prev
	lc.removeNode(tailNode)
	return tailNode
}

func (lc *LruCache) Set(key string, value interface{}, expiration time.Duration) {
	now := time.Now()

	node, ok := lc.cache[key]
	if !ok {
		if lc.size >= lc.capacity {
			tailNode := lc.removeTail()
			delete(lc.cache, tailNode.key)
			lc.size--
		}
		node = &Node{
			key:              key,
			value:            value,
			lastModifiedTime: now,
			expiration:       expiration,
		}
		lc.cache[key] = node
		lc.addToHead(node)
		lc.size++
	} else {
		node.value = value
		node.lastModifiedTime = now
		lc.moveToHead(node)
		lc.size++
	}
}
