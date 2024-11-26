package main

import (
	"log"

	"github.com/kuroko-shirai/basket/internal/models/storage"
)

func main() {
	newStorage := storage.New(func(args []any) int {
		a1 := args[0].(int)
		a2 := args[1].(int)

		return func(a, b int) int {
			return a + b
		}(a1, a2)
	}, storage.Int, storage.Int)

	newStorage.Add(1, 1)
	newStorage.Add(1, 2)
	newStorage.Add(1, 1)

	newStorage.Do()

	log.Println(newStorage)

	newStorage.Add(1, 1)
	newStorage.Add(1, 1)

	log.Println(newStorage)

	newStorage.Do()

	log.Println(newStorage)
}
