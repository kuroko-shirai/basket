package storage

type Storage[T comparable] struct {
	QueryPool            QueryPool[T]
	FractionsPool        FractionsPool
	FractionsQueriesList FractionsQueriesList
	Task                 Task[T]
}

func New[T comparable](fun func(args []any) T, signatures ...any) *Storage[T] {
	return &Storage[T]{
		QueryPool:            NewQueryPool[T](),
		FractionsPool:        NewFractionPool(),
		FractionsQueriesList: NewFractionsQueriesList(),
		Task:                 NewTask(fun, signatures),
	}
}

func (s *Storage[T]) Add(args ...any) {
	newQueryID := s.QueryPool.Add(args)

	newFractionID := s.FractionsPool.Add(args)

	s.FractionsQueriesList.Add(newFractionID, newQueryID)
}

func (s *Storage[T]) Do() {
	for _, fraction := range s.FractionsPool.Fractions {
		ret := s.Task.Do(fraction.Args...)

		if queries, ok := s.FractionsQueriesList[fraction.ID]; ok {
			for _, id := range queries {
				if _, ok := s.QueryPool[id]; ok {
					s.QueryPool[id] = Query[T]{
						Args: s.QueryPool[id].Args,
						Ret:  ret,
					}
				}
			}
		}
	}
}
