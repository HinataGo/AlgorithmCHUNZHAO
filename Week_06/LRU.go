package Week_06

import "runtime/debug"

func init() {
	debug.FreeOSMemory()
	debug.SetGCPercent(-1)
}

type LRUCache struct {
	keyValue [3001]*valueNode
	list     *valueList
	cap      int
	len      int
}

type valueNode struct {
	key   int
	value int
	prev  *valueNode
	next  *valueNode
}

type valueList struct {
	handle *valueNode
}

func newValueList() *valueList {
	list := &valueList{
		handle: &valueNode{},
	}
	list.handle.next = list.handle
	list.handle.prev = list.handle
	return list
}

func (list *valueList) push(node *valueNode) {
	node.next = list.handle
	node.prev = list.handle.prev
	list.handle.prev.next = node
	list.handle.prev = node
}

func (list *valueList) pop(node *valueNode) *valueNode {
	if node == nil {
		node = list.handle.next
	}
	if node == list.handle {
		return nil
	}
	node.prev.next = node.next
	node.next.prev = node.prev
	node.next = nil
	node.prev = nil
	return node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		keyValue: [3001]*valueNode{},
		list:     newValueList(),
		cap:      capacity,
		len:      0,
	}
}

func (c *LRUCache) Get(key int) int {
	if c.len == 0 {
		return -1
	}
	if node := c.keyValue[key]; node != nil {
		c.update(node)
		return node.value
	}
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	if c.cap == 0 {
		return
	}
	if node := c.keyValue[key]; node != nil {
		node.value = value
		c.update(node)
		return
	}
	for c.len >= c.cap { // 用 for 不用 if 可支持扩容缩容
		c.pop()
	}
	node := &valueNode{key: key, value: value}
	c.insert(node)
}

func (c *LRUCache) update(node *valueNode) {
	c.list.pop(node)
	c.list.push(node)
}

func (c *LRUCache) insert(node *valueNode) {
	if c.len >= c.cap {
		return
	}
	c.list.push(node)
	c.keyValue[node.key] = node
	c.len++
}

func (c *LRUCache) pop() {
	if c.len == 0 {
		return
	}
	node := c.list.pop(nil)
	c.keyValue[node.key] = nil
	c.len--
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
