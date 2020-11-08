package stack

type Stack struct {
	data []ElementType
}

func New() *Stack {
	return &Stack{}
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack) Clear() {
	s.data = s.data[:0]
}

func (s *Stack) Len() int {
	return len(s.data)
}

func (s *Stack) Push(e ElementType) {
	s.data = append(s.data, e)
}

func (s *Stack) PushAll(es ...ElementType) {
	s.data = append(s.data, es...)
}

// Pop if stack.data is empty, could cause fatal error
func (s *Stack) Pop() ElementType {
	e := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return e
}

// Peek if stack.data is empty, could cause fatal error
func (s *Stack) Peek() ElementType {
	e := s.data[len(s.data)-1]
	return e
}

func (s *Stack) Traverse(visit func(e ElementType)) {
	for i := len(s.data) - 1; i >= 0; i-- {
		visit(s.data[i])
	}
}
