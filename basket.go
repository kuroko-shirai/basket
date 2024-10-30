package basket

import (
	"time"

	"github.com/kuroko-shirai/basket/internal/models/cache"
	"github.com/kuroko-shirai/basket/internal/models/item"
	"github.com/kuroko-shirai/basket/internal/models/response"
)

type Cache interface {
	Set(key int32, value response.Response)
	Keys() []int32
	Get(key int32) (response.Response, bool)
	Has(key int32) bool
}

type Basket struct {
	Size  int
	TTL   time.Duration
	Items []item.Item
	Cache Cache
}

func New(size int, poll time.Duration, ttl time.Duration) *Basket {
	newCache, err := cache.New(poll, ttl)
	if err != nil {
		return nil
	}

	return &Basket{
		Size:  size,
		TTL:   ttl,
		Cache: newCache,
	}
}

func (b *Basket) Add(element int32) {
	b.Items = append(b.Items, *item.New(element))
}
