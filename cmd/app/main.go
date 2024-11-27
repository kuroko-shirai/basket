package main

import (
	"context"
	"log"

	"github.com/kuroko-shirai/basket"
)

func main() {
	newStorage := basket.New(func(args []any) int {
		a1 := args[0].(int)
		a2 := args[1].(int)

		return func(a, b int) int {
			return a + b
		}(a1, a2)
	}, func(ctx context.Context, arg any) {
		sum := arg.(int)

		log.Println("release:", sum)
	}, basket.Int, basket.Int)

	newStorage.Add(1, 1)
	newStorage.Add(1, 2)
	newStorage.Add(1, 1)

	newStorage.Do()

	newStorage.Add(1, 1)
	newStorage.Add(1, 1)

	newStorage.Do()

	newStorage.Release(context.TODO())
}
