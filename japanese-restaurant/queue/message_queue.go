package queue

import "context"

type Queue interface {
	SendOrder(ctx context.Context, order Order) error
	ReceiveOrder(ctx context.Context) (Order, error)
	Close()
}
