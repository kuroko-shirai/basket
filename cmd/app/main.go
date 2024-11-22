package main

import (
	"fmt"

	"github.com/kuroko-shirai/basket/pkg/set"
)

func main() {
	newSet := set.New[int]()
	newSet.Add(1)
	newSet.Add(2)
	newSet.Add(2)
	newSet.Add(3)

	fmt.Println(newSet)
	// newBasket, err := basket.New(10, 5*time.Second, 3*time.Second)
	// if err != nil {
	// 	panic(err)
	// }

	// // Пусть service-cache хранит пару ответов от некоторого
	// // процесса.
	// newBasket.Save(2, response.Response{
	// 	Timestamp: time.Now().Unix(),
	// 	Message:   "message-2",
	// })

	// newBasket.Save(5, response.Response{
	// 	Timestamp: time.Now().Unix(),
	// 	Message:   "message-5",
	// })

	// // Заполняем корзину некоторыми элементами, которые на
	// // текущий момент имеют тип int32. В дальнейшем, это
	// // будут запросы, либо некоторые задачи (func), которые
	// // обращаются к service-cache.
	// newBasket.Add(task.New(
	// 	func(recovery any) {
	// 		log.Printf("Panic! %!w", recovery)
	// 	},
	// 	func(ctx context.Context, id int32) func() {
	// 		return func() {
	// 			newBasket.Get(id)
	// 		}
	// 	}(context.Background(), int32(5)),
	// ))

	// newBasket.Add(task.New(
	// 	func(recovery any) {
	// 		log.Printf("Panic! %!w", recovery)
	// 	},
	// 	func(ctx context.Context, id int32) func() {
	// 		return func() {
	// 			newBasket.Get(id)
	// 		}
	// 	}(context.Background(), int32(5)),
	// ))

	// newBasket.Add(task.New(
	// 	func(recovery any) {
	// 		log.Printf("Panic! %!w", recovery)
	// 	},
	// 	func(ctx context.Context, id int32) func() {
	// 		return func() {
	// 			newBasket.Get(id)
	// 		}
	// 	}(context.Background(), int32(5)),
	// ))

	// newBasket.Add(task.New(
	// 	func(recovery any) {
	// 		log.Printf("Panic! %!w", recovery)
	// 	},
	// 	func(ctx context.Context, id int32) func() {
	// 		return func() {
	// 			newBasket.Get(id)
	// 		}
	// 	}(context.Background(), int32(2)),
	// ))

	// keys := newBasket.Keys.Items()
	// for k, v := range keys {
	// 	log.Println("k:", k, "v:", v)
	// }

	// //newBasket.Do()

	// time.Sleep(3 * time.Second)
}
