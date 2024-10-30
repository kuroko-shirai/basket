package main

import (
	"fmt"
	"time"

	"github.com/kuroko-shirai/basket"
	"github.com/kuroko-shirai/basket/internal/models/response"
)

func main() {
	newBasket := basket.New(10, 5*time.Second, 3*time.Second)

	newBasket.Add(1) // add 1 item
	newBasket.Add(3) // add 2 item
	newBasket.Add(4) // add 3 item

	for index, item := range newBasket.Items {

		time.Sleep(1 * time.Second) // Do something

		newBasket.Cache.Set(
			item.Element, response.Response{
				Timestamp: time.Now().Unix(),
				Message:   fmt.Sprintf("message %d", index),
			})

		for
	}

	fmt.Println(newBasket)
}
