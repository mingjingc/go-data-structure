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
			t.Errorf("iterator Value expect %d, but got %d", initData[i], value)
		}
		i++
	}

	node := l.Index(0)
	if node != l.head {
		t.Error("index 0 should place head")
	}
	if node.Value != initData[0] {
		t.Errorf("Value of head should be %d", node.Value)
	}

	node = l.Index(-1)
	if node != l.tail {
		t.Error("index -1 should place tail")
	}
	if node.Value != initData[len(initData)-1] {
		t.Errorf("Value of tail should be %d", initData[len(initData)-1])
	}

	l.AddNodeTail(5)
	length++
	node = l.Index(-1)
	if node.Value != 5 {
		t.Errorf("Value expect %d, but got %d", 5, node.Value)
	}
	if node != l.tail {
		t.Errorf("tail should update new node, Value: %d", 5)
	}
	if l.len != length {
		t.Errorf("length expect %d, but got %d", length, l.len)
	}

	l.AddNodeHead(6)
	length++
	node = l.Index(0)
	if node.Value != 6 {
		t.Errorf("Value expect %d, but got %d", 6, node.Value)
	}
	if node != l.head {
		t.Errorf("head should update to new node, Value %d", 6)
	}
	if l.len != length {
		t.Errorf("length expect %d, but got %d", length, l.len)
	}

	node = l.Index(-4)
	if node.Value != 2 {
		t.Errorf("index %d, Value expect %d, but got %d", -4, 2, node.Value)
	}

	l.DeleteNode(node)
	length--
	if l.len != length {
		t.Errorf("after delete a element, length expect %d, but got %d", length, l.len)
	}
	node = l.Index(-4)
	if node.Value != 1 {
		t.Errorf("after delete %d node, index %d, Value expect %d, but got %d", -4, -4, 1, node.Value)
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
		if p1.Value != p2.Value {
			t.Error("copy fail")
		}
		p1 = p1.next
		p2 = p2.next
	}

}
