package list

import (
	"testing"
)

func TestDoubleLinkedList(t *testing.T) {
	l := NewDoubleLinkedList()
	initData := []int{1, 2, 3, 4}
	length := len(initData)
	for _, v := range initData {
		l.AddNodeTail(v)
	}
	if l.len != len(initData) {
		t.Errorf("expect length %d bug got %d", len(initData), l.len)
	}

	iterator := l.GetIterator(false)
	i := 0
	for iterator.HasNext() {
		value := iterator.Next()
		if value != initData[i] {
			t.Errorf("iterator value expect %d, but got %d", initData[i], value)
		}
		i++
	}

	node := l.Index(0)
	if node != l.head {
		t.Error("index 0 should place head")
	}
	if node.value != initData[0] {
		t.Errorf("value of head should be %d", node.value)
	}

	node = l.Index(-1)
	if node != l.tail {
		t.Error("index -1 should place tail")
	}
	if node.value != initData[len(initData)-1] {
		t.Errorf("value of tail should be %d", initData[len(initData)-1])
	}

	l.AddNodeTail(5)
	length++
	node = l.Index(-1)
	if node.value != 5 {
		t.Errorf("value expect %d, but got %d", 5, node.value)
	}
	if node != l.tail {
		t.Errorf("tail should update new node, value: %d", 5)
	}
	if l.len != length {
		t.Errorf("length expect %d, but got %d", length, l.len)
	}

	l.AddNodeHead(6)
	length++
	node = l.Index(0)
	if node.value != 6 {
		t.Errorf("value expect %d, but got %d", 6, node.value)
	}
	if node != l.head {
		t.Errorf("head should update to new node, value %d", 6)
	}
	if l.len != length {
		t.Errorf("length expect %d, but got %d", length, l.len)
	}

	node = l.Index(-4)
	if node.value != 2 {
		t.Errorf("index %d, value expect %d, but got %d", -4, 2, node.value)
	}

	l.DeleteNode(node)
	length--
	if l.len != length {
		t.Errorf("after delete a element, length expect %d, but got %d", length, l.len)
	}
	node = l.Index(-4)
	if node.value != 1 {
		t.Errorf("after delete %d node, index %d, value expect %d, but got %d", -4, -4, 1, node.value)
	}

	node = l.tail
	l.Rotate()
	if node != l.head {
		t.Error("DoubleLinkedList rotate fail")
	}

	copy := l.Duplicate()
	if copy.len != l.len {
		t.Error("copy length not equal to DoubleLinkedList")
	}
	p1, p2 := l.head, copy.head
	for p1 != nil {
		if p1.value != p2.value {
			t.Error("copy fail")
		}
		p1 = p1.next
		p2 = p2.next
	}

}
