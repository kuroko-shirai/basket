package storage

import (
	"errors"
	"fmt"
)

var (
	Int8   int8
	Int16  int16
	Int32  int32
	Int64  int64
	Int    int
	String string
)

// functor - executes a function with the given signature and compares arguments with signature.
func functor[T comparable](fun func(args []any) T, signatures ...any) (func(...any) T, error) {
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

type task[T comparable] struct {
	Function func(...any) T
}

func newTask[T comparable](fun func(args []any) T, signatures []any) task[T] {
	f, _ := functor(fun, signatures...)

	return task[T]{
		Function: f,
	}
}

// do - returns a function with the given signature.
func (t *task[T]) do(signatures ...any) T {
	return t.Function(signatures...)
}
