package stack

type ElementType = int

type IStack interface {
	IsEmpty() bool
	Clear() bool
	Len() int
	Pop() ElementType
	Push(e ElementType)
	PushAll(es ...ElementType)
	Peek() ElementType
	Traverse(visit func(e ElementType))
}
