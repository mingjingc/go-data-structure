package stack

type ElementType = int

type Stack interface {
	IsEmpty() bool
	Clear()
	Len() int
	Pop() ElementType
	Push(e ElementType)
	PushAll(es ...ElementType)
	Peek() ElementType
	Traverse(visit func(e ElementType))
}
