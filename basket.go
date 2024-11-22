package basket

import (
	"fmt"
	"log"
	"time"

	"github.com/kuroko-shirai/basket/internal/models/cache"
	"github.com/kuroko-shirai/basket/internal/models/item"
	"github.com/kuroko-shirai/basket/internal/models/response"
	"github.com/kuroko-shirai/basket/pkg/set"
	"github.com/kuroko-shirai/task"
	cmap "github.com/orcaman/concurrent-map/v2"
)

type Cache interface {
	Set(key int32, value response.Response)
	Keys() []int32
	Get(key int32) (response.Response, bool)
	Has(key int32) bool
}

type Basket struct {
	Size    int
	TTL     time.Duration
	Buckets []item.Item
	Keys    cmap.ConcurrentMap[int32, set.Set]
	cache   Cache
}

func New(size int, poll time.Duration, ttl time.Duration) (*Basket, error) {
	newCache, err := cache.New(poll, ttl)
	if err != nil {
		return nil, err
	}

	return &Basket{
		Size:    size,
		TTL:     ttl,
		cache:   newCache,
		Keys:    cmap.NewWithCustomShardingFunction[int32, set.Set](func(key int32) uint32 { return uint32(key) }),
		Buckets: make([]item.Item, 0),
	}, nil
}

func (b *Basket) Add(t task.Task) {
	b.Buckets = append(b.Buckets, *item.New(t))
	b.sync(len(b.Buckets))
}

func (b *Basket) Get(id int32) {
	if resp, ok := b.cache.Get(id); ok {
		log.Println(
			"id_cache:", id,
			"message:", resp.Message,
		)

		if buckets, ok := b.Keys.Get(id); ok {
			fmt.Println(">>:", id, "[", buckets, "]")
		}

		// for index, rec := range b.Buckets {
		// 	fmt.Println(">> 1")
		// 	if rec.ID == id && rec.Index != index {
		// 		b.remove(index)
		// 		fmt.Println(">> 2")
		// 	}
		// }
	}
}

// func (b *Basket) remove(index int) {
// 	b.Buckets = slice.Remove(b.Buckets, index)
// }

func (b *Basket) Save(id int32, resp response.Response) {
	b.cache.Set(id, resp)
}

func (b *Basket) Do() {
	for id, task := range b.Buckets {
		log.Println("bucket_id:", id)
		task.Element.Do()
	}
}

func (b *Basket) sync(idBucket int) {
	for _, idCache := range b.cache.Keys() {
		buckets, _ := b.Keys.Get(idCache)
		fmt.Println(buckets)
		buckets.Add(idBucket)
		b.Keys.Set(idCache, buckets)

	}
}

/*
// Метод добавляет
func (b *Basket) Re(basketsItemIndex int, basketsItem item.Item) {
	for index, item := range b.Buckets {
		if index == basketsItemIndex {
			continue
		}

		if basketsItem.Element == item.Element {
			// Send data to request queue
			Do()

			// Remove elements
			b.remove(index)
		}
	}

	fmt.Println(b.Buckets)
	b.remove(basketsItemIndex)
	fmt.Println(b.Buckets)
}

func Do() {
	time.Sleep(3 * time.Second)
}
*/
