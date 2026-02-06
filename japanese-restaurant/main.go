package main

import (
	"context"
	"time"

	"github.com/tulara/japanese-restaurant/japanese-restaurant/customer"
	"github.com/tulara/japanese-restaurant/japanese-restaurant/queue"
	"github.com/tulara/japanese-restaurant/japanese-restaurant/restaurant"
)

func main() {
	q := queue.NewChannelQueue()
	restaurant := restaurant.New(q)
	customerElinor := customer.New(q)
	customerEdward := customer.New(q)
	customerMarianne := customer.New(q)
	ctx, cancel := context.WithCancel(context.Background())

	restaurant.Open(ctx)

	// customers can only place one order for now.
	// TODO: this concurrently
	// TODO: customers place orders in English
	customerElinor.PlaceOrder(ctx, queue.NewOrder(map[int]int{1: 1}))
	customerEdward.PlaceOrder(ctx, queue.NewOrder(map[int]int{6: 1}))
	customerMarianne.PlaceOrder(ctx, queue.NewOrder(map[int]int{2: 1, 5: 1}))

	time.Sleep(2 * time.Second)
	cancel()

	restaurant.ListInventory()
}
