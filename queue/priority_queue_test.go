package queue

import (
	"container/heap"
	"testing"
)

var Items = []*Item{
	{
		Value:    2,
		Priority: 2,
	},

	{
		Value:    5,
		Priority: 5,
	},
	{
		Value:    1,
		Priority: 1,
	},

	{
		Value:    3,
		Priority: 3,
	},

	{
		Value:    4,
		Priority: 4,
	},
}

func TestPriorityQueue(t *testing.T)  {
	priorityQueue := NewPriorityQueue()

	for _, item := range Items {
		heap.Push(priorityQueue, item)
	}
	item := heap.Pop(priorityQueue)
	if item==nil || item.(*Item).Value!=1{
		t.Errorf("PriorityQueue: pop expect %v, but got %v",1, item.(*Item).Value)
	}
}
