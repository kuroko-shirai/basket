package storage

type Query struct {
	ID   int
	Args []any
}

type QueryPool struct {
	Queries []Query
}

func NewQueryPool() QueryPool {
	return QueryPool{
		Queries: make([]Query, 0),
	}
}

func (qp *QueryPool) Add(args ...any) int {
	newQuery := Query{
		ID:   len(qp.Queries),
		Args: args,
	}

	qp.Queries = append(qp.Queries, newQuery)

	return newQuery.ID
}
