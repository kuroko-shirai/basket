package basket

import (
	"fmt"
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

func New(size int, poll time.Duration, ttl time.Duration) (*Basket, error) {
	newCache, err := cache.New(poll, ttl)
	if err != nil {
		return nil, err
	}

	return &Basket{
		Size:  size,
		TTL:   ttl,
		Cache: newCache,
	}, nil
}

func (b *Basket) Add(element int32) {
	b.Items = append(b.Items, *item.New(element))
}

func (b *Basket) remove(index int) {
	b.Items[index] = b.Items[len(b.Items)-1] // Copy last element to index i.
	b.Items[len(b.Items)-1] = item.Item{}    // Erase last element (write zero value).
	b.Items = b.Items[:len(b.Items)-1]       // Truncate slice.
}

func (b *Basket) Re(basketsItemIndex int, basketsItem item.Item) {
	for index, item := range b.Items {
		if index == basketsItemIndex {
			continue
		}

		if basketsItem.Element == item.Element {
			// Send data to request queue

			// Remove elements
			b.remove(index)
		}
	}

	fmt.Println(b.Items)
	b.remove(basketsItemIndex)
	fmt.Println(b.Items)
}
