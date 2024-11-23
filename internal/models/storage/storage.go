package storage

type Storage struct {
	QueryPool            QueryPool
	FractionsPool        FractionsPool
	FractionsQueriesList FractionsQueriesList
}

func New() *Storage {
	return &Storage{
		QueryPool:            NewQueryPool(),
		FractionsPool:        NewFractionPool(),
		FractionsQueriesList: NewFractionsQueriesList(),
	}
}

func (s *Storage) Add(args ...any) {
	newQueryID := s.QueryPool.Add(args)

	newFractionID := s.FractionsPool.Add(args)

	s.FractionsQueriesList.Add(newFractionID, newQueryID)
}

func (s *Storage) Do() {

}
