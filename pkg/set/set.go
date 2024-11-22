package set

// Set - структура, хранящая уникальные числа.
type Set[T comparable] struct {
	slice    []T
	elements map[T]struct{}
}

func New[T comparable](size ...int) Set[T] {
	if len(size) > 1 {
		panic("incorrect size of set")
	}

	if len(size) == 1 {
		return Set[T]{
			elements: make(map[T]struct{}, size[0]),
			slice:    make([]T, 0, size[0]),
		}
	}

	return Set[T]{
		elements: make(map[T]struct{}),
		slice:    make([]T, 0),
	}
}

func (s *Set[T]) Add(element T) {
	if s.elements == nil {
		s.elements = make(map[T]struct{})
	}

	if _, ok := s.elements[element]; !ok {
		s.elements[element] = struct{}{}
		s.slice = append(s.slice, element)
	}
}

func (s *Set[T]) Slice() []T {
	return s.slice
}
