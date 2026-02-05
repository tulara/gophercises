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

func (r *Restaurant) ProcessOrder(ctx context.Context) error {
	order, err := r.q.RecieveOrder(ctx)
	if err != nil {
		return err
	}

	id := order.ID()

	r.inventoryMtx.Lock()
	defer r.inventoryMtx.Unlock()

	currentStock := r.inventory[id]
	r.inventory[id] = currentStock - 1

	mealName := Menu[id]
	fmt.Printf("Order for %s recieved.\n", mealName)

	return nil
}

func (r *Restaurant) ListInventory() {
	fmt.Println("We have...")

	r.inventoryMtx.RLock()
	defer r.inventoryMtx.RUnlock()

	for meal, stock := range r.inventory {
		mealName := Menu[meal]
		fmt.Printf("%d of the %s\n", stock, mealName)
	}
}
