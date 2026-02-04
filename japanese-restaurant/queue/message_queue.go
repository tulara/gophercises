package queue

import "context"

type Queue interface {
	SendOrder(order Order) bool
	RecieveOrder(ctx context.Context) (Order, error)
}
