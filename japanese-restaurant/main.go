package main

import (
	"context"
	"sync"

	"github.com/tulara/japanese-restaurant/japanese-restaurant/customer"
	"github.com/tulara/japanese-restaurant/japanese-restaurant/queue"
	"github.com/tulara/japanese-restaurant/japanese-restaurant/restaurant"
)

func main() {
	q := queue.NewRedisStreamsQueue()
	restaurant := restaurant.New(q)

	ctx := context.Background()
	done := make(chan struct{})

	var wg sync.WaitGroup

	wg.Go(func() {
		for {
			restaurant.Open(ctx, done)
		}
	})

	wg.Go(func() {
		customersPlaceOrders(ctx, q)
	})

	wg.Wait()
	close(done)
	q.Close()

	restaurant.ListInventory()
}

func customersPlaceOrders(ctx context.Context, q queue.Queue) {
	customerElinor := customer.New(q)
	customerEdward := customer.New(q)
	customerMarianne := customer.New(q)

	//pork gyoza
	customerElinor.PlaceOrder(ctx, queue.NewOrder(map[int]int{1: 1}))
	// salmon ngiri
	customerEdward.PlaceOrder(ctx, queue.NewOrder(map[int]int{6: 1}))
	// edamame and tempura eggplant
	customerMarianne.PlaceOrder(ctx, queue.NewOrder(map[int]int{2: 6, 5: 1}))
}
