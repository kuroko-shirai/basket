package cache

import (
	"errors"
	"time"

	cmap "github.com/orcaman/concurrent-map/v2"
)

type Cache[T any] struct {
	cm   cmap.ConcurrentMap[int32, Container[T]]
	ttl  time.Duration
	poll time.Duration
}

type Container[T any] struct {
	Value    T
	LastCall time.Time
}

func New[T any](poll time.Duration, ttl time.Duration) (*Cache[T], error) {
	cm := cmap.NewWithCustomShardingFunction[int32, Container[T]](func(key int32) uint32 {
		return uint32(key)
	})

	return &Cache[T]{
		cm:   cm,
		ttl:  ttl,
		poll: poll,
	}, nil
}

func (cache *Cache[T]) Set(key int32, value Container[T]) {
	cache.cm.Set(key, value)
}

func (cache *Cache[T]) Keys() []int32 {
	return cache.cm.Keys()
}

func (cache *Cache[T]) Get(key int32) (Container[T], bool) {
	if time.Since(cache.cm.Items()[key].LastCall) > cache.ttl {
		return cache.cm.Get(key)
	}
	return Container[T]{}, false
}

func (cache *Cache[T]) CheckKeyTTL(key int32) {
	if time.Since(cache.cm.Items()[key].LastCall) > cache.ttl {
		cache.cm.Remove(key)
	}
}

func (cache *Cache[T]) Process() {
	for {
		for _, key := range cache.cm.Keys() {
			cache.CheckKeyTTL(key)
		}

		time.Sleep(cache.poll)
	}
}

func (cache *Cache[T]) AddKeys(keys []int32, values []T) error {
	if len(keys) != len(values) {
		return errors.New("invalid dimensions")
	}

	for i, key := range keys {
		cache.cm.Set(key, Container[T]{Value: values[i], LastCall: time.Now()})
	}

	return nil
}

func (cache *Cache[T]) Has(key int32) bool {
	return cache.cm.Has(key)
}