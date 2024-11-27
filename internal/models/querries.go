package models

import (
	"github.com/google/uuid"
)

type Query[T comparable] struct {
	Ret  T
	Args []any
}

type Queries[T comparable] map[uuid.UUID]Query[T]

func NewQueries[T comparable]() Queries[T] {
	return make(map[uuid.UUID]Query[T])
}

func (qs Queries[T]) Add(args []any) uuid.UUID {
	id := uuid.New()

	qs[id] = Query[T]{
		Args: args,
	}

	return id
}
