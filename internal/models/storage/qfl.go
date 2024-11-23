package storage

type FractionsQueriesList map[int][]int

func NewFractionsQueriesList() FractionsQueriesList {
	return make(map[int][]int, 0)
}

func (qfl FractionsQueriesList) Add(fractionID, queryID int) {
	queriesList, ok := qfl[fractionID]
	if !ok {
		queriesList = make([]int, 0, 1)
	}
	queriesList = append(queriesList, queryID)
	qfl[fractionID] = queriesList
}

func (qfl FractionsQueriesList) Get(fractionID int) ([]int, bool) {
	queries, ok := qfl[fractionID]
	return queries, ok
}
