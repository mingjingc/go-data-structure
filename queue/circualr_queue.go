package queue

type CircularQueue struct {
	queue []interface{}
	// number of element, can be easily check empty or full without calculate rear and front
	length int
	// capacity of element
	capacity int

	front int
	rear int
}

func NewCircularQueue(capacity int) *CircularQueue {
	return &CircularQueue{
		queue:    make([]interface{},capacity),
		length:   0,
		capacity: capacity,
		front:    0,
		rear:     0,
	}
}

func (cq *CircularQueue) EnQueue(value interface{}) bool  {
	if cq.length == cq.capacity {
		return false
	}

	cq.queue[cq.rear] = value
	cq.rear = (cq.rear+1)%cq.capacity
	cq.length++
	return true
}
func (cq *CircularQueue) DeQueue() interface{}  {
	if cq.length==0{
		return nil
	}
	value:=cq.queue[cq.front]
	cq.front = (cq.front+1)%cq.capacity
	cq.length--
	return value
}
func (cq *CircularQueue) isEmpty() bool  {
	return cq.length==0
}

func (cq *CircularQueue) Len() int {
	return cq.length
}
func (cq *CircularQueue) Cap() int {
	return cq.capacity
}