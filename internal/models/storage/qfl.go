package storage

type fractionsToQueries map[int][]int

func newFractionsToQueries() fractionsToQueries {
	return make(map[int][]int)
}

func (qfl fractionsToQueries) Add(fractionID, queryID int) {
	queriesList, ok := qfl[fractionID]
	if !ok {
		queriesList = make([]int, 0, 1)
	}
	queriesList = append(queriesList, queryID)
	qfl[fractionID] = queriesList
}

func (qfl fractionsToQueries) Get(fractionID int) ([]int, bool) {
	queries, ok := qfl[fractionID]

	return queries, ok
}
