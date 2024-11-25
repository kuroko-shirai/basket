package storage

type Query[T comparable] struct {
	Ret  T
	Args []any
}

type QueryPool[T comparable] map[int]Query[T]

func NewQueryPool[T comparable]() QueryPool[T] {
	return make(map[int]Query[T], 0)
}

func (qp QueryPool[T]) Add(args []any) int {
	id := len(qp)
	qp[id] = Query[T]{
		Args: args,
	}

	return id
}
