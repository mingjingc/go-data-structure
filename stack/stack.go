package stack

type SliceStack struct {
	data []ElementType
}

func New() Stack {
	return &SliceStack{}
}

func (s *SliceStack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *SliceStack) Clear() {
	s.data = s.data[:0]
}

func (s *SliceStack) Len() int {
	return len(s.data)
}

func (s *SliceStack) Push(e ElementType) {
	s.data = append(s.data, e)
}

func (s *SliceStack) PushAll(es ...ElementType) {
	s.data = append(s.data, es...)
}

// Pop if stack.data is empty, could cause fatal error
func (s *SliceStack) Pop() ElementType {
	e := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return e
}

// Peek if stack.data is empty, could cause fatal error
func (s *SliceStack) Peek() ElementType {
	e := s.data[len(s.data)-1]
	return e
}

func (s *SliceStack) Traverse(visit func(e ElementType)) {
	for i := len(s.data) - 1; i >= 0; i-- {
		visit(s.data[i])
	}
}
