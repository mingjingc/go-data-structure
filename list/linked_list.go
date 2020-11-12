package list

type Node struct {
	pre   *Node
	next  *Node
	Value interface{}
}

type LinkedList struct {
	head *Node
	tail *Node
	len  int
}

type LinkedListIterator struct {
	next          *Node
	startFromTail bool
}

func New() *LinkedList {
	return &LinkedList{
		head: nil,
		tail: nil,
		len:  0,
	}
}

func (l *LinkedList) GetIterator(startFromTail bool) Iterator {
	iterator := &LinkedListIterator{
		next:          nil,
		startFromTail: startFromTail,
	}
	if startFromTail {
		iterator.next = l.tail
		return iterator
	}
	iterator.next = l.head
	return iterator
}

func (i *LinkedListIterator) HasNext() bool {
	return i.next != nil
}

func (i *LinkedListIterator) Next() interface{} {
	value := i.next.Value
	if i.startFromTail {
		i.next = i.next.pre
		return value
	}
	i.next = i.next.next
	return value
}

// delete all
func (l *LinkedList) listEmpty() {
	l.tail = nil
	l.head = nil
	l.len = 0
}

func (l *LinkedList) AddNodeHead(value interface{}) {
	node := &Node{
		pre:   nil,
		next:  nil,
		Value: value,
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
func (l *LinkedList) AddNodeTail(value interface{}) {
	node := &Node{
		pre:   nil,
		next:  nil,
		Value: value,
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
func (l *LinkedList) InsertNode(oldNode *Node, value interface{}, after bool) {
	node := &Node{
		pre:   nil,
		next:  nil,
		Value: value,
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

func (l *LinkedList) MoveToHead(node *Node) {
	if node == l.head {
		return
	}
	node.pre.next = node.next
	if node.next != nil {
		node.next.pre = node.pre
	} else {
		l.tail = node.pre
	}

	node.pre = nil
	node.next = l.head
	l.head.pre = node
	l.head = node
}

func (l *LinkedList) DeleteNode(node *Node) {
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
func (l *LinkedList) Index(index int) *Node {
	var node *Node
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

func (l *LinkedList) GetHead() *Node {
	return l.Index(0)
}

func (l *LinkedList) GetTail() *Node {
	return l.Index(-1)
}

func (l *LinkedList) GetLen() int {
	return l.len
}

// Rotate tail to head
func (l *LinkedList) Rotate() {
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

func (l *LinkedList) Join(list *LinkedList) {
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

func (l *LinkedList) Duplicate() *LinkedList {
	copy := New()

	p := l.head
	for p != nil {
		copy.AddNodeTail(p.Value)
		p = p.next
	}
	return copy
}
