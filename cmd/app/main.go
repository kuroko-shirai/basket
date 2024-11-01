package main

import (
	"fmt"
	"time"

	"github.com/kuroko-shirai/basket"
	"github.com/kuroko-shirai/basket/internal/models/response"
)

func main() {
	newBasket, err := basket.New(10, 5*time.Second, 3*time.Second)
	if err != nil {
		panic(err)
	}

	newBasket.Add(3) // Add 1 item
	newBasket.Add(1) // Add 2 item
	newBasket.Add(4) // Add 3 item
	newBasket.Add(3) // Add 4 item
	newBasket.Add(1) // Add 5 item

	fmt.Println(">:", newBasket)

	for basketsItemIndex, basketsItem := range newBasket.Items {

		time.Sleep(1 * time.Second) // Do something and got a response
		newResponse := fmt.Sprintf("message %d", basketsItemIndex)

		newBasket.Cache.Set(
			basketsItem.Element, response.Response{
				Timestamp: time.Now().Unix(),
				Message:   newResponse,
			},
		)

		// Обратное движение с проверкой элементов
		newBasket.Re(basketsItemIndex, basketsItem)
		fmt.Println(">:", newBasket)
	}

	fmt.Println("<:", newBasket)
}
