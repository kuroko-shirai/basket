package main

import (
	"context"
	"log"
	"time"

	"github.com/kuroko-shirai/basket"
)

func sum(a, b int) int {
	time.Sleep(time.Millisecond)

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

	{
		t0 := time.Now()

		for i := 1; i <= 10; i++ {
			newBasket.Add(1, 1)
			if newBasket.Size() == 5 {
				newBasket.Do()
				newBasket.Release(context.TODO())
			}
		}

		log.Println("process took", time.Since(t0).Milliseconds())
		log.Println("basket size", newBasket.Size())
	}

	{
		t0 := time.Now()

		for i := 1; i <= 10; i++ {
			sum(1, 1)
		}

		log.Println("process took", time.Since(t0).Milliseconds())
	}
}
