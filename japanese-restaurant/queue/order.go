package queue

import "sync/atomic"

var nextOrderID atomic.Int64

type Order struct {
	id    int
	Meals map[int]int //meal id and number ordered
}

func NewOrder(meals map[int]int) Order {
	id := nextOrderID.Add(1)
	return Order{
		id:    int(id),
		Meals: meals,
	}
}

func (o *Order) ID() int {
	return o.id
}
