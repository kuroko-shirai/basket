package main

import (
	"context"
	"log"

	"github.com/kuroko-shirai/basket"
)

func sum(a, b int) int {
	return a + b
}

func main() {
	newBasket := basket.New(func(args []any) int {
		a := args[0].(int)
		b := args[1].(int)

		return sum(a, b)
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
