package item

import "time"

// Item
type Item struct {
	Timestamp int64
	Element   any
}

func New(element any) *Item {
	return &Item{
		Element:   element,
		Timestamp: time.Now().Unix(),
	}
}
