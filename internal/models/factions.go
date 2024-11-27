package models

import (
	"reflect"

	"github.com/google/uuid"
)

type Factions map[uuid.UUID][]any

func NewFractions() Factions {
	return make(map[uuid.UUID][]any)
}

func (fs Factions) Add(args []any) uuid.UUID {
	for fID, fArgs := range fs {
		if reflect.DeepEqual(fArgs, args) {
			return fID
		}
	}

	id := uuid.New()

	fs[id] = args

	return id
}
