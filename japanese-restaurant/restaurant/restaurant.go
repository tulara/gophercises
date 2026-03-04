package restaurant

import (
	"context"
	"fmt"
	"sync"

	"github.com/tulara/japanese-restaurant/japanese-restaurant/queue"
)

type Restaurant struct {
	q            queue.Queue
	inventoryMtx sync.RWMutex
	inventory    map[int]int // product id and number stored
}

func New(q queue.Queue) *Restaurant {
	return &Restaurant{
		q: q,
		inventory: map[int]int{
			1: 5,
			2: 5,
			3: 5,
			4: 5,
			5: 5,
			6: 5,
		},
	}
}

func (r *Restaurant) Open(ctx context.Context, done <-chan struct{}) bool {
	// if orders have finished sending, make sure they are processed before closing
	// otherwise keep listening for an order

	select {
	case <-done:
		orders, err := r.q.Drain(ctx)
		if err != nil {
			fmt.Printf("Error draining restaurant: %v", err)
			return false
		}
		for _, o := range orders {
			r.updateInventory(ctx, o)
		}
		return false
	default:
		found, err := r.ProcessOrder(ctx)
		if err != nil {
			fmt.Printf("Error processing order: %v\n", err)
			return false
		}
		return found
	}
}

func (r *Restaurant) ProcessOrder(ctx context.Context) (bool, error) {
	order, err := r.q.ReceiveOrder(ctx)
	if err != nil {
		return false, err
	}

	if order.ID == 0 {
		return false, nil
	}

	fmt.Printf("Order %d is being prepared\n", order.ID)

	r.updateInventory(ctx, order)
	return true, nil

}

func (r *Restaurant) updateInventory(ctx context.Context, order queue.Order) {
	r.inventoryMtx.Lock()
	defer r.inventoryMtx.Unlock()
	for mealID, amountOrdered := range order.Meals {
		id := mealID
		currentStock := r.inventory[id]

		mealName := Menu[id]
		if amountOrdered > currentStock {
			r.inventory[id] = 0
			numberMissing := amountOrdered - currentStock
			fmt.Printf("Apologies, %d %s will be missing from your order\n", numberMissing, mealName)
		} else {
			r.inventory[id] = currentStock - amountOrdered
		}

		fmt.Printf("Order up for %s.\n", mealName)
	}
}

func (r *Restaurant) ListInventory() {
	fmt.Println("\nThe restaurant now has...")

	r.inventoryMtx.RLock()
	defer r.inventoryMtx.RUnlock()

	for meal, stock := range r.inventory {
		mealName := Menu[meal]
		fmt.Printf("%d of the %s\n", stock, mealName)
	}
}
