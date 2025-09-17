type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() T {
	n := len(s.items) - 1
	item := s.items[n]
	s.items = s.items[:n]
	return item
}

func NewStack[T]() *Stack[T] {
	return &Stack{
		items : make([]T, 0)
	}
}