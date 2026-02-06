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

func New(q queue.Queue) Restaurant {
	return Restaurant{
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

func (r *Restaurant) Open(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				r.ProcessOrder(ctx)
			}
		}
	}()
}

func (r *Restaurant) ProcessOrder(ctx context.Context) error {
	order, err := r.q.RecieveOrder(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("Order %d is being prepared\n", order.ID())

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

	return nil
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
