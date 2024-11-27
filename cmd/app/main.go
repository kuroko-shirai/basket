package main

import (
	"context"
	"log"

	"github.com/kuroko-shirai/basket"
)

func main() {
	newBasket := basket.New(func(args []any) int {
		a1 := args[0].(int)
		a2 := args[1].(int)

		return func(a, b int) int {
			return a + b
		}(a1, a2)
	}, func(ctx context.Context, arg any) {
		sum := arg.(int)

		log.Println("release:", sum)
	}, basket.Int, basket.Int)

	newBasket.Add(1, 1)
	newBasket.Add(1, 2)
	newBasket.Add(1, 1)

	newBasket.Do()

	newBasket.Add(1, 1)
	newBasket.Add(1, 1)

	newBasket.Do()

	newBasket.Release(context.TODO())
}
