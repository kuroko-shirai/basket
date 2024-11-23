package storage

type Task interface {
	Do() Return
}

type Return struct {
	Value any
	Err   error
}

type HandleFunc func() Return

type WithRecover struct {
	Task

	Handle HandleFunc
}

func NewTask(handle HandleFunc) Task {
	return &WithRecover{
		Handle: handle,
	}
}

func (task *WithRecover) Do() Return {
	return task.Handle()
}
