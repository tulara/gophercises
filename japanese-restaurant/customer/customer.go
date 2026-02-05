package customer

import (
	"log/slog"

	"github.com/tulara/japanese-restaurant/japanese-restaurant/queue"
)

type Customer struct {
	q queue.Queue
}

func New(q queue.Queue) Customer {
	return Customer{
		q: q,
	}
}

func (c *Customer) PlaceOrder(order queue.Order) {
	sent := c.q.SendOrder(order)
	if !sent {
		slog.Error("Failed to place order")
	}
}
