package item

import "time"

type Item struct {
	Birth   int64
	Element any
}

func New(element any) *Item {
	return &Item{
		Element: element,
		Birth:   time.Now().Unix(),
	}
}
