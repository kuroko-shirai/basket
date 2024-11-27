package models

type Query[T comparable] struct {
	Ret  T
	Args []any
}

type Queries[T comparable] map[int]Query[T]

func NewQueries[T comparable]() Queries[T] {
	return make(map[int]Query[T])
}

func (qs Queries[T]) Add(args []any) int {
	id := len(qs)
	qs[id] = Query[T]{
		Args: args,
	}

	return id
}
