package models

import (
	"reflect"
)

type Fractions map[int][]any

func NewFractions() Fractions {
	return make(map[int][]any)
}

func (fs Fractions) Add(args []any) int {
	for fID, fArgs := range fs {
		if reflect.DeepEqual(fArgs, args) {
			return fID
		}
	}

	id := len(fs)
	fs[id] = args

	return id
}
