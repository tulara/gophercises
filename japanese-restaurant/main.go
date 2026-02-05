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

	// customers can only place one order for now.
	// TODO: this concurrently
	// TODO: customers place orders in English
	customerElinor.PlaceOrder(queue.NewOrder(map[int]int{1: 1}))
	restaurant.ProcessOrder(context.Background())

	customerEdward.PlaceOrder(queue.NewOrder(map[int]int{6: 1}))
	restaurant.ProcessOrder(context.Background())

	customerMarianne.PlaceOrder(queue.NewOrder(map[int]int{2: 1, 5: 1}))
	restaurant.ProcessOrder(context.Background())

	restaurant.ListInventory()
}
