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
}

func New[T comparable](
	fun func(args []any) T,
	signatures ...any,
) *Storage[T] {
	return &Storage[T]{
		task:      newTask(fun, signatures),
		queries:   newQueries[T](),
		fractions: newFractions(),
		onfly:     newProcesses(),
		completed: newProcesses(),
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

func (s *Storage[T]) realease() {
	for key, qsID := range s.completed {
		var wg sync.WaitGroup

		wg.Add(1)
		for _, qID := range qsID {

			newTask := t.New(
				func(recovery any) {
					log.Printf("Panic in the workflow process! %!w", recovery)
				},
				func(ctx context.Context, out chan<- ChCacheGetItemsOnOff) func() {
					return func() {
						defer wg.Done()

						out <- ChCacheGetItemsOnOff{
							State: 1,
							Err:   nil,
						}
					}
				}(context.Background(), ch),
			)
			s.queries[qID].release()
		}
	}
}
