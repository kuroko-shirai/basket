package storage

import (
	"reflect"
)

type Fraction struct {
	ID   int
	Args []any
}

type FractionsPool map[int][]any

func NewFractionPool() FractionsPool {
	return make(map[int][]any, 0)
}

func (fp FractionsPool) Add(args []any) int {
	for fpID, fpArgs := range fp {
		if reflect.DeepEqual(fpArgs, args) {
			return fpID
		}
	}

	id := len(fp)
	fp[id] = args

	return id
}
