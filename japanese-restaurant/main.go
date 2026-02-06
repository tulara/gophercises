package main

import (
	"context"

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
	ctx := context.Background()

	done := restaurant.Open(ctx)

	// customers can only place one order for now.
	customerElinor.PlaceOrder(ctx, queue.NewOrder(map[int]int{1: 1}))
	customerEdward.PlaceOrder(ctx, queue.NewOrder(map[int]int{6: 1}))
	customerMarianne.PlaceOrder(ctx, queue.NewOrder(map[int]int{2: 6, 5: 1}))

	q.Close()
	<-done // wait for restaurant to finish processing orders.

	restaurant.ListInventory()
}
