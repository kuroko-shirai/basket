package storage

type query[T comparable] struct {
	Ret  T
	Args []any
}

type queries[T comparable] map[int]query[T]

func newQueries[T comparable]() queries[T] {
	return make(map[int]query[T])
}

func (qs queries[T]) Add(args []any) int {
	id := len(qs)
	qs[id] = query[T]{
		Args: args,
	}

	return id
}
