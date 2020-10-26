package queue

type Item struct {
	Value interface{}
	Priority int
}

type PriorityQueue []*Item

func NewPriorityQueue() *PriorityQueue  {
	return &PriorityQueue{}
}

func (pq PriorityQueue) Len() int {
	return len(pq)
}
func (pq PriorityQueue) Less(i,j int) bool {
	return pq[i].Priority < pq[j].Priority
}
func (pq PriorityQueue) Swap(i,j int)  {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Pop() interface{} {
	if pq.Len() == 0 {
		return nil
	}
	item := (*pq)[pq.Len()-1]
	*pq = (*pq)[:pq.Len()-1]
	return item
}

func (pq *PriorityQueue) Push(item interface{})   {
	*pq = append(*pq, item.(*Item))
}