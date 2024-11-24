package storage

import (
	"reflect"
)

type Fraction struct {
	ID   int
	Args []any
}

type FractionsPool struct {
	Fractions []Fraction
}

func NewFractionPool() FractionsPool {
	return FractionsPool{
		Fractions: make([]Fraction, 0),
	}
}

func (fp *FractionsPool) Add(args []any) int {
	for _, fraction := range fp.Fractions {
		if reflect.DeepEqual(fraction.Args, args) {
			return fraction.ID
		}
	}

	newFraction := Fraction{
		ID:   len(fp.Fractions),
		Args: args,
	}

	fp.Fractions = append(fp.Fractions, newFraction)

	return newFraction.ID
}
