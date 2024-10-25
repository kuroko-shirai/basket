package basket

import (
	"sort"

	"github.com/kuroko-shirai/basket/internal/models/cache"
	"github.com/kuroko-shirai/basket/internal/models/item"
	"github.com/kuroko-shirai/basket/internal/models/response"
)

type Basket struct {
	Size  int
	TTL   int64
	Items []item.Item
	Cache cache.Cache[response.Response]
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
