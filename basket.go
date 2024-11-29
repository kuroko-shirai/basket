package basket

import (
	"context"
	"log"

	t "github.com/kuroko-shirai/task"

	"github.com/kuroko-shirai/basket/internal/models"
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
	fractions models.Factions
	onfly     models.Processes
	completed models.Processes
	releaser  func(ctx context.Context, arg any)
}

func New[T comparable](
	fun func(args []any) T,
	rel func(ctx context.Context, arg any),
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

func (b *Basket[T]) Add(args ...any) {
	newQueryID := b.queries.Add(args)

	newFractionID := b.fractions.Add(args)

	b.onfly.Add(newFractionID, newQueryID)
}

func (b *Basket[T]) Do() {
	for fID, fsArgs := range b.fractions {
		ret := b.task.Do(fsArgs...)

		if queries, ok := b.onfly[fID]; ok {
			for _, qID := range queries {
				if _, ok := b.queries[qID]; ok {
					b.queries[qID] = models.Query[T]{
						Args: b.queries[qID].Args,
						Ret:  ret,
					}

					delete(b.fractions, fID)
				}
			}
		}

		b.completed[fID] = append(
			b.completed[fID],
			b.onfly[fID]...,
		)
		delete(b.onfly, fID)
	}
}

func (b *Basket[T]) Release(ctx context.Context) {
	for _, qsID := range b.completed {
		g := t.WithRecover(
			func(recovery any, args ...any) {
				log.Println("panic: %!w", recovery)
			},
		)

		for _, qID := range qsID {
			g.Do(func() func() error {
				return func() error {
					if query, ok := b.queries[qID]; ok {
						b.releaser(ctx, query.Ret)
					}

					return nil
				}
			}())
		}

		if err := g.Wait(); err != nil {
			log.Println("something wrong in first batch", err)
		}
	}

	for gID := range b.completed {
		delete(b.completed, gID)
	}

	for qID := range b.queries {
		delete(b.queries, qID)
	}
}

func (b *Basket[T]) Size() int {
	return len(b.queries)
}
