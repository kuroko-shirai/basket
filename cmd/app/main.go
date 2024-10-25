package main

import (
	"fmt"

	"github.com/kuroko-shirai/basket"
)

func main() {
	newBasket := basket.New(10, 10)

	newBasket.Add([]int32{2, 16, 9, 0, 1})
	newBasket.Add([]int32{2, 16, 9, 0, 1})
	newBasket.Add([]int32{2, 16, 9, 3, 1})

	for _, item := range newBasket.Items {
		fmt.Println(item.Element)

	}
}
