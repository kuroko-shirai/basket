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

type Task struct {
	Function func(...any)
}

func NewTask(signatures []any) Task {
	f, _ := Function(signatures...)

	return Task{
		Function: f,
	}
}

// Function возвращает функцию с заданной сигнатурой.
func Function(signatures ...any) (func(...any), error) {
	return func(args ...any) {
		if len(args) != len(signatures) {
			panic(errors.New("wrong number of arguments"))
		}
		for i, arg := range args {
			if fmt.Sprintf("%T", arg) != fmt.Sprintf("%T", signatures[i]) {
				panic(errors.New("wrong type of argument"))
			}
		}

		fmt.Println(args...) // Здесь заменить на task.Task.Do()
	}, nil
}

// Do возвращает функцию с заданной сигнатурой.
func (t *Task) Do(signatures ...any) {
	t.Function(signatures...)
}

func TypeOf(arg any) string {
	return fmt.Sprintf("%T", arg)
}
