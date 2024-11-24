package storage

type Storage struct {
	QueryPool            QueryPool
	FractionsPool        FractionsPool
	FractionsQueriesList FractionsQueriesList
	Task                 Task
}

func New(signatures ...any) *Storage {
	return &Storage{
		QueryPool:            NewQueryPool(),
		FractionsPool:        NewFractionPool(),
		FractionsQueriesList: NewFractionsQueriesList(),
		Task:                 NewTask(signatures),
	}
}

func (s *Storage) Add(args ...any) {
	newQueryID := s.QueryPool.Add(args)

	newFractionID := s.FractionsPool.Add(args)

	s.FractionsQueriesList.Add(newFractionID, newQueryID)
}

func (s *Storage) Do() {
	for _, fraction := range s.FractionsPool.Fractions {
		s.Task.Do(fraction.Args...)
	}
}
