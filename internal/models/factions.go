package models

import (
	"reflect"
)

type Factions map[int][]any

func NewFractions() Factions {
	return make(map[int][]any)
}

func (fs Factions) Add(args []any) int {
	for fID, fArgs := range fs {
		if reflect.DeepEqual(fArgs, args) {
			return fID
		}
	}

	id := len(fs)
	fs[id] = args

	return id
}
