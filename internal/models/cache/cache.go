package cache

import (
	"errors"
	"time"

	"github.com/kuroko-shirai/basket/internal/models/response"
	"github.com/kuroko-shirai/cache"
)

type Cache cache.Cache[response.Response]

func New(poll time.Duration, ttl time.Duration) (*cache.Cache[response.Response], error) {
	newCache, err := cache.New[response.Response](&cache.Config{
		Poll: 5 * time.Second,
		TTL:  3 * time.Second,
		CLS:  true,
	})
	if err != nil {
		return nil, errors.New("invalid cache")
	}

	return newCache, nil
}

func (c *Cache) Set(key int32, value response.Response) {
	c.Set(key, value)
}

func (cache *Cache) Keys() []int32 {
	return cache.Keys()
}

func (cache *Cache) Get(key int32) (response.Response, bool) {
	return cache.Get(key)
}

func (cache *Cache) Has(key int32) bool {
	return cache.Has(key)
}
