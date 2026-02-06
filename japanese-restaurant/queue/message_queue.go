package queue

import "context"

type Queue interface {
	SendOrder(ctx context.Context, order Order) bool
	RecieveOrder(ctx context.Context) (Order, error)
}
