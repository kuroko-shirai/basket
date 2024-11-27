package storage

import (
	"context"
	"log"
	"sync"

	t "github.com/kuroko-shirai/task"
)

type Storage[T comparable] struct {
	task      task[T]
	queries   queries[T]
	fractions fractions
	onfly     processes
	completed processes
	releaser  func(arg any)
}

func New[T comparable](
	fun func(args []any) T,
	rel func(arg any),
	signatures ...any,
) *Storage[T] {
	return &Storage[T]{
		task:      newTask(fun, signatures),
		queries:   newQueries[T](),
		fractions: newFractions(),
		onfly:     newProcesses(),
		completed: newProcesses(),
		releaser:  rel,
	}
}

func (s *Storage[T]) Add(args ...any) {
	newQueryID := s.queries.Add(args)

	newFractionID := s.fractions.Add(args)

	s.onfly.Add(newFractionID, newQueryID)
}

func (s *Storage[T]) Do() {
	for fID, fsArgs := range s.fractions {
		ret := s.task.do(fsArgs...)

		if queries, ok := s.onfly[fID]; ok {
			for _, qID := range queries {
				if _, ok := s.queries[qID]; ok {
					s.queries[qID] = query[T]{
						Args: s.queries[qID].Args,
						Ret:  ret,
					}

					delete(s.fractions, fID)
				}
			}
		}

		s.completed[fID] = append(s.completed[fID], s.onfly[fID]...)
		delete(s.onfly, fID)
	}
}

func (s *Storage[T]) Release(ctx context.Context) {
	for _, qsID := range s.completed {
		var wg sync.WaitGroup

		wg.Add(len(qsID))

		for _, qID := range qsID {
			newTask := t.New(
				func(recovery any) {
					log.Println("panic: %!w", recovery)
				},
				func(ctx context.Context) func() {
					return func() {
						defer wg.Done()

						if query, ok := s.queries[qID]; ok {
							s.releaser(query.Ret)
						}
					}
				}(context.Background()),
			)

			newTask.Do()
		}

		wg.Wait()
	}

	for gID := range s.completed {
		delete(s.completed, gID)
	}

	for qID := range s.queries {
		delete(s.queries, qID)
	}
}
