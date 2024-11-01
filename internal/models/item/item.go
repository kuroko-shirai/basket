package item

import "time"

type Item struct {
	Timestamp int64
	Element   int32
}

func New(element int32) *Item {
	return &Item{
		Element:   element,
		Timestamp: time.Now().Unix(),
	}
}
