package models

import (
	"errors"
	"fmt"
)

type Task[T comparable] struct {
	Function func(...any) T
}

// functor - executes a function with the given signature
// and compares arguments with signature.
func functor[T comparable](
	fun func(args []any) T,
	signatures ...any,
) (func(...any) T, error) {
	return func(args ...any) T {
		if len(args) != len(signatures) {
			panic(errors.New("wrong number of arguments"))
		}
		for i, arg := range args {
			if fmt.Sprintf("%T", arg) != fmt.Sprintf("%T", signatures[i]) {
				panic(errors.New("wrong type of argument"))
			}
		}

		return fun(args)
	}, nil
}

func NewTask[T comparable](
	fun func(args []any) T,
	signatures []any,
) Task[T] {
	f, _ := functor(fun, signatures...)

	return Task[T]{
		Function: f,
	}
}

// Do - returns a function with the given signature.
func (t *Task[T]) Do(signatures ...any) T {
	return t.Function(signatures...)
}
