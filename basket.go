package basket

import (
	"context"
	"log"
	"sync"

	"github.com/kuroko-shirai/basket/internal/models"
	t "github.com/kuroko-shirai/task"
)

var (
	Int8   int8
	Int16  int16
	Int32  int32
	Int64  int64
	Int    int
	String string
)

type Basket[T comparable] struct {
	task      models.Task[T]
	queries   models.Queries[T]
	fractions models.Fractions
	onfly     models.Processes
	completed models.Processes
	releaser  func(arg any)
}

func New[T comparable](
	fun func(args []any) T,
	rel func(arg any),
	signatures ...any,
) *Basket[T] {
	return &Basket[T]{
		task:      models.NewTask(fun, signatures),
		queries:   models.NewQueries[T](),
		fractions: models.NewFractions(),
		onfly:     models.NewProcesses(),
		completed: models.NewProcesses(),
		releaser:  rel,
	}
}

func (s *Basket[T]) Add(args ...any) {
	newQueryID := s.queries.Add(args)

	newFractionID := s.fractions.Add(args)

	s.onfly.Add(newFractionID, newQueryID)
}

func (s *Basket[T]) Do() {
	for fID, fsArgs := range s.fractions {
		ret := s.task.Do(fsArgs...)

		if queries, ok := s.onfly[fID]; ok {
			for _, qID := range queries {
				if _, ok := s.queries[qID]; ok {
					s.queries[qID] = models.Query[T]{
						Args: s.queries[qID].Args,
						Ret:  ret,
					}

					delete(s.fractions, fID)
				}
			}
		}

		s.completed[fID] = append(
			s.completed[fID],
			s.onfly[fID]...,
		)
		delete(s.onfly, fID)
	}
}

func (s *Basket[T]) Release(ctx context.Context) {
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
