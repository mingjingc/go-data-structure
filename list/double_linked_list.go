package list

type DoubleLinkedListNode struct {
	pre   *DoubleLinkedListNode
	next  *DoubleLinkedListNode
	value interface{}
}

type DoubleLinkedList struct {
	head *DoubleLinkedListNode
	tail *DoubleLinkedListNode
	len  int
}

func NewDoubleLinkedList() *DoubleLinkedList {
	return &DoubleLinkedList{
		head: nil,
		tail: nil,
		len:  0,
	}
}

// 删除所有元素
func (l *DoubleLinkedList) listEmpty() {
	l.tail = nil
	l.head = nil
	l.len = 0
}

func (l *DoubleLinkedList) AddNodeHead(value interface{}) {
	node := &DoubleLinkedListNode{
		pre:   nil,
		next:  nil,
		value: value,
	}
	if l.head == nil {
		l.head = node
		l.tail = node
		l.len++
		return
	}

	l.head.pre = node
	node.next = l.head
	l.head = node
	l.len++
}
func (l *DoubleLinkedList) AddNodeTail(value interface{}) {
	node := &DoubleLinkedListNode{
		pre:   nil,
		next:  nil,
		value: value,
	}
	if l.head == nil {
		l.head = node
		l.tail = node
		l.len++
		return
	}

	l.tail.next = node
	node.pre = l.tail
	l.tail = node
	l.len++
}
func (l *DoubleLinkedList) InsertNode(oldNode *DoubleLinkedListNode, value interface{}, after bool) {
	node := &DoubleLinkedListNode{
		pre:   nil,
		next:  nil,
		value: value,
	}
	if after {
		node.next = oldNode.next
		node.pre = oldNode
		if oldNode.next != nil {
			oldNode.next.pre = node
		}
		oldNode.next = node
		if l.tail == oldNode {
			l.tail = node
		}
	} else {
		node.next = oldNode
		node.pre = oldNode.pre
		if oldNode.pre != nil {
			oldNode.pre.next = node
		}
		oldNode.pre = node
		if l.head == oldNode {
			l.head = node
		}
	}
	l.len++
}

func (l *DoubleLinkedList) DeleteNode(node *DoubleLinkedListNode) {
	if node == l.head && l.len == 1 {
		l.head = nil
		l.tail = nil
	} else if node == l.head {
		node.next.pre = nil
		l.head = node.next
		node.next = nil
	} else if node == l.tail {
		node.pre.next = nil
		l.tail = node.pre
		node.pre = nil
	} else {
		node.pre.next = node.next
		node.next.pre = node.pre
		node.pre = nil
		node.next = nil
	}
	l.len--
}

// index为非负数，0表示头
// index为负数，-1表示尾
func (l *DoubleLinkedList) Index(index int) *DoubleLinkedListNode {
	var node *DoubleLinkedListNode
	if index >= 0 {
		p := l.head
		for p != nil {
			if index == 0 {
				return p
			}
			index--
			p = p.next
		}
	} else {
		p := l.tail
		index = (-index) - 1
		for p != nil {
			if index == 0 {
				return p
			}
			index--
			p = p.pre
		}
	}

	return node
}

// 旋转一个单位
func (l *DoubleLinkedList) Rotate() {
	if l.len < 2 {
		return
	}
	node := l.tail
	l.tail = l.tail.pre
	l.tail.next = nil

	l.head.pre = node
	node.next = l.head
	node.pre = nil
	l.head = node
}

func (l *DoubleLinkedList) Join(list *DoubleLinkedList) {
	if l.len == 0 {
		l.tail = list.tail
		l.head = list.head
		l.len = list.len
		return
	}
	l.tail.next = list.head
	list.head.pre = l.tail
	l.tail = list.tail
}

func (l *DoubleLinkedList) Duplicate() *DoubleLinkedList {
	copy := &DoubleLinkedList{
		head: nil,
		tail: nil,
		len:  0,
	}

	p := l.head
	for p != nil {
		copy.AddNodeTail(p.value)
		p = p.next
	}
	return copy
}