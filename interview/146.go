/*
*  @author: zengjinlin@didiglobal.com
*  @Date: 2021/2/2
	设计一个LRU
*/

package interview

type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*DLinkNode
	head, tail *DLinkNode
}

type DLinkNode struct {
	key   int
	value int
	prev   *DLinkNode
	next  *DLinkNode
}

func initDLinkNode(key, value int) *DLinkNode {
	return &DLinkNode{
		key:   key,
		value: value,
	}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		capacity:capacity,
		cache: map[int]*DLinkNode{},
		head:initDLinkNode(0,0),
		tail:initDLinkNode(0,0),
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok{
		return -1
	}
	node := this.cache[key]
	this.moveToHead(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cache[key]; !ok {
		node := initDLinkNode(key, value)
		this.cache[key] = node
		this.addToHead(node)
		this.size++
		if this.size > this.capacity {
			remove := this.removeToTail()
			delete(this.cache, remove.key)
			this.size --
		}
	} else {
		node := this.cache[key]
		node.value = value
		this.moveToHead(node)
	}
}

func (this *LRUCache) addToHead(node *DLinkNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *DLinkNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) moveToHead (node *DLinkNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeToTail() *DLinkNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}