package queue

import "testing"

func TestCircularQueue(t *testing.T) {
	queue := NewCircularQueue(3)

	queue.EnQueue(1)
	queue.EnQueue(2)
	if queue.Len()!=2{
		t.Errorf("Queue length expect %d, but got %d",2,queue.Len())
	}
	value:=queue.DeQueue()
	if value==nil || value!=1{
		t.Errorf("DeQueue expect %v, but got %v",1,value)
	}

	queue.EnQueue(3)
	queue.EnQueue(4)
	if queue.EnQueue(5) {
		t.Errorf("Queue is full, Enqueue %v should be failed!",5)
	}

	value = queue.DeQueue()
	if value==nil || value!=2{
		t.Errorf("Dequeue expect %v, but got %v",2,value)
	}
	value = queue.DeQueue()
	if value==nil || value!=3{
		t.Errorf("Dequeue expect %v, but got %v",3,value)
	}
	value = queue.DeQueue()
	if value==nil || value!=4{
		t.Errorf("Dequeue expect %v, but got %v",4,value)
	}

	if !queue.isEmpty(){
		t.Errorf("Queue should be empty!")
	}

	if !queue.EnQueue(6){
		t.Errorf("Enqueue %v failed",6)
	}
	if !queue.EnQueue(7){
		t.Errorf("Enqueue %v failed",7)
	}
	if !queue.EnQueue(8){
		t.Errorf("Enqueue %v failed",8)
	}
	value = queue.DeQueue()
	if value == nil || value!=6{
		t.Errorf("Dequeue expect %v, but got %v",6,value)
	}

}
