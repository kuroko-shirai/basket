package storage

import (
	"reflect"
)

type fractions map[int][]any

func newFractions() fractions {
	return make(map[int][]any)
}

func (fp fractions) Add(args []any) int {
	for fpID, fpArgs := range fp {
		if reflect.DeepEqual(fpArgs, args) {
			return fpID
		}
	}

	id := len(fp)
	fp[id] = args

	return id
}
