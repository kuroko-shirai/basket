package item

import (
	"time"

	"github.com/kuroko-shirai/task"
)

type Item struct {
	Timestamp int64
	Element   task.Task
}

func New(element task.Task) *Item {
	return &Item{
		Element:   element,
		Timestamp: time.Now().Unix(),
	}
}
