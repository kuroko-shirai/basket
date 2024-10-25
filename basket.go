package basket

import (
	"sort"

	"github.com/kuroko-shirai/basket/internal/models/item"
)

type Basket struct {
	Size  int
	TTL   int64
	Items []item.Item
}

func New(size int, ttl int64) *Basket {
	return &Basket{
		Size: size,
		TTL:  ttl,
	}
}

func (b *Basket) Add(element []int32) {
	sort.Slice(element, func(i, j int) bool {
		return element[i] < element[j]
	})

	b.Items = append(b.Items, *item.New(element))
}
