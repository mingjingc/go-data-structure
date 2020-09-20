package cache

import (
	"github.com/mingjingc/go-data-structure/list"
)

type entry struct {
	key   int
	value interface{}
}

type LRUCache struct {
	cache    map[int]*list.DoubleLinkedListNode
	list     *list.DoubleLinkedList
	capacity int
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		cache:    map[int]*list.DoubleLinkedListNode{},
		list:     list.NewDoubleLinkedList(),
		capacity: capacity,
	}
}

func (c *LRUCache) Get(key int) interface{} {
	node := c.get(key)
	if node == nil {
		return nil
	}
	return node.Value.(*entry).value
}
func (c *LRUCache) get(key int) *list.DoubleLinkedListNode {
	node, ok := c.cache[key]
	if !ok {
		return nil
	}
	c.list.MoveToHead(node)
	return node
}

func (c *LRUCache) Put(key int, value interface{}) {
	node := c.get(key)
	if node != nil {
		node.Value.(*entry).value = value
		return
	}
	if c.list.GetLen() == c.capacity {
		node = c.list.GetTail()
		c.list.DeleteNode(node)
		deleteKey := node.Value.(*entry).key
		delete(c.cache, deleteKey)
	}
	e := &entry{
		key:   key,
		value: value,
	}
	c.list.AddNodeHead(e)
	node = c.list.GetHead()
	c.cache[key] = node
}
