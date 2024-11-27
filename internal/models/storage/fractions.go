package storage

import (
	"reflect"
)

type fractions map[int][]any

func newFractions() fractions {
	return make(map[int][]any)
}

func (fs fractions) add(args []any) int {
	for fID, fArgs := range fs {
		if reflect.DeepEqual(fArgs, args) {
			return fID
		}
	}

	id := len(fs)
	fs[id] = args

	return id
}
