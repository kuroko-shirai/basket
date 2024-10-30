package main

import (
	"fmt"
	"time"

	"github.com/kuroko-shirai/basket"
	"github.com/kuroko-shirai/basket/internal/models/response"
)

func main() {
	newBasket := basket.New(10, 10)

	newBasket.Add([]int32{2, 16, 9, 0, 1})
	newBasket.Add([]int32{2, 16, 9, 0, 1})
	newBasket.Add([]int32{2, 16, 9, 3, 1})

	for _, item := range newBasket.Items {
		newBasket.Cache.Set(item.Element..., response.Response{
				Timestamp: time.Now().Unix(),
				Message:   "Hello",
			})
	}

	fmt.Println(newBasket)
}
