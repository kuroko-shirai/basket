package storage

type Storage[T comparable] struct {
	queries            queries[T]
	fractions          fractions
	fractionsToQueries fractionsToQueries
	task               task[T]
}

func New[T comparable](fun func(args []any) T, signatures ...any) *Storage[T] {
	return &Storage[T]{
		queries:            newQueries[T](),
		fractions:          newFractions(),
		fractionsToQueries: newFractionsToQueries(),
		task:               newTask(fun, signatures),
	}
}

func (s *Storage[T]) Add(args ...any) {
	newQueryID := s.queries.Add(args)

	newFractionID := s.fractions.Add(args)

	s.fractionsToQueries.Add(newFractionID, newQueryID)
}

func (s *Storage[T]) Do() {
	for fpID, fpArgs := range s.fractions {
		ret := s.task.do(fpArgs...)

		if queries, ok := s.fractionsToQueries[fpID]; ok {
			for _, id := range queries {
				if _, ok := s.queries[id]; ok {
					s.queries[id] = query[T]{
						Args: s.queries[id].Args,
						Ret:  ret,
					}

					delete(s.fractions, fpID)
				}
			}
		}
	}
}
