package stack

type Stack[T any] []T

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(item T) {
	*s = append(*s, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zeroValue T
		return zeroValue, false
	}

	length := (*s).Len()
	item := (*s)[length-1]
	*s = (*s)[:length-1]

	return item, true
}

func (s *Stack[T]) Len() int {
	return len(*s)
}

func (s *Stack[T]) IsEmpty() bool {
	return (*s).Len() == 0
}

func (s *Stack[T]) Clear() {
	*s = Stack[T]{}
}

func (s *Stack[T]) Top() (T, bool) {
	if (*s).IsEmpty() {
		var zeroValue T
		return zeroValue, false
	}
	return (*s)[(*s).Len()-1], true
}
