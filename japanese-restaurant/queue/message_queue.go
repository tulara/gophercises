package queue

import "context"

type Queue interface {
	SendOrder(ctx context.Context, order Order) bool
	ReceiveOrder(ctx context.Context) (Order, error)
	Close()
}
